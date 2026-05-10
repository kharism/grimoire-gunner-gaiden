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

var lastShoot = time.Now()

var Revertbackspritetime time.Time

var grenade = weapons.NewGrenade()

func PlayerAttackHandler(e *ecs.ECS) {
	playerQuery := donburi.NewQuery(filter.Contains(
		component.PlayerTag,
	))
	if ebiten.IsKeyPressed(ebiten.KeyE) {
		if time.Since(lastShoot) >= delay {
			lastShoot = time.Now()
			//shoot
			playerE, _ := playerQuery.FirstEntity(e.World)
			playerPos := component.Position.GetValue(playerE)
			component.Sprite.Get(playerE).Image = assets.SvenSprite2
			Revertbackspritetime = time.Now().Add(300 * time.Millisecond)

			weapons.BasicProjectile(e, component.PositionComponentData{X: playerPos.X + float64(component.GridLength), Y: playerPos.Y, Z: playerPos.Z})
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		playerE, _ := playerQuery.FirstEntity(e.World)
		component.Sprite.Get(playerE).Image = assets.SvenSprite2
		Revertbackspritetime = time.Now().Add(300 * time.Millisecond)
		grenade.Cast(e)
	}
	if time.Now().After(Revertbackspritetime) {
		playerE, _ := playerQuery.FirstEntity(e.World)
		component.Sprite.Get(playerE).Image = assets.SvenSprite1
	}
}
