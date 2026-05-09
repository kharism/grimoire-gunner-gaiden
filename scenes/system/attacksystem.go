package system

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kharism/GrimoireGunner2/scenes/assets"
	"github.com/kharism/GrimoireGunner2/scenes/component"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/filter"
)

var delay, _ = time.ParseDuration("200ms")

var lastShoot = time.Now()

func basicProjectile(ecs *ecs.ECS, pos component.PositionComponentData) {
	component.NewProjectile(ecs.World, component.ProjectileParam{
		Vx:     20,
		Vy:     0,
		Pos:    pos,
		Damage: 2,
		Sprite: assets.Bullet,
		OnHit:  SingleHitProjectile,
	})
}

// use this as single hit projectile. Once a projectile hit,
// apply damage then disappear. Basically the default behaviour of any projectile based attack
func SingleHitProjectile(ecs *ecs.ECS, projectile, receiver *donburi.Entry) {
	damage := component.Damage.Get(projectile).Damage
	component.Health.Get(receiver).HP -= damage
	ecs.World.Remove(projectile.Entity())
}

var Revertbackspritetime time.Time

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

			basicProjectile(e, component.PositionComponentData{X: playerPos.X + float64(GridLength), Y: playerPos.Y, Z: playerPos.Z})
		}
	}
	if time.Now().After(Revertbackspritetime) {
		playerE, _ := playerQuery.FirstEntity(e.World)
		component.Sprite.Get(playerE).Image = assets.SvenSprite1
	}
}
