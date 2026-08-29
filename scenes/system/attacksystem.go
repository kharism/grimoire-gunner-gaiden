package system

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/kharism/GrimoireGunner2/scenes/assets"
	"github.com/kharism/GrimoireGunner2/scenes/component"
	"github.com/kharism/GrimoireGunner2/scenes/system/weapons"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/filter"
)

var delay, _ = time.ParseDuration("200ms")

// delay between shoot with e key
var DefaultDelay, _ = time.ParseDuration("200ms")

var damage = 2
var DefaultDamage = 2

var lastShoot = time.Now()

var Revertbackspritetime time.Time

type RenderableCaster interface {
	Caster
	RenderableWeaponUi
	Tick()
}

var SelectedSlot = 0
var WeaponSlot = []RenderableCaster{}

func SetDamage(dmg int) {
	damage = dmg
}

// set delay between shoot with e key
func SetDelay(dur time.Duration) {
	delay = dur
}

func PlayerAttackHandler(e *ecs.ECS) {
	playerQuery := donburi.NewQuery(filter.Contains(
		component.PlayerTag,
	))
	if ebiten.IsKeyPressed(ebiten.KeyE) {
		playerE, _ := playerQuery.FirstEntity(e.World)
		component.State.SetValue(playerE, component.STATE_ATTACK)
		if time.Since(lastShoot) >= delay {

			lastShoot = time.Now()
			//shoot

			playerPos := component.Position.GetValue(playerE)
			component.Sprite.Get(playerE).Image = assets.SvenSprite2
			Revertbackspritetime = time.Now().Add(300 * time.Millisecond)

			projectileEntity := weapons.BasicProjectile(e, component.PositionComponentData{X: playerPos.X + float64(component.GridLength), Y: playerPos.Y, Z: playerPos.Z})
			projectileEntry := e.World.Entry(*projectileEntity)
			component.Damage.Get(projectileEntry).Damage = damage
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.Key1) {
		SelectedSlot = 0
	}
	if inpututil.IsKeyJustPressed(ebiten.Key2) {
		SelectedSlot = 1
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		playerE, _ := playerQuery.FirstEntity(e.World)
		component.Sprite.Get(playerE).Image = assets.SvenSprite2
		Revertbackspritetime = time.Now().Add(300 * time.Millisecond)
		component.State.SetValue(playerE, component.STATE_ATTACK)
		if WeaponSlot[SelectedSlot].GetCooldownProgress() == 1.0 {
			WeaponSlot[SelectedSlot].Cast(e)
		}

	}
	if time.Now().After(Revertbackspritetime) {
		//component.State.SetValue(playerE, component.STATE_ATTACK)
		playerE, ok := playerQuery.FirstEntity(e.World)
		if ok {
			component.State.SetValue(playerE, component.STATE_STANDING)
		}

	}
}
