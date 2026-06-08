package weapons

import (
	"math"

	"github.com/kharism/GrimoireGunner2/scenes/assets"
	"github.com/kharism/GrimoireGunner2/scenes/component"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func BasicProjectile(ecs *ecs.ECS, pos component.PositionComponentData) {
	component.NewProjectile(ecs.World, component.ProjectileParam{
		Vx:     20,
		Vy:     0,
		Pos:    pos,
		Damage: 2,
		Sprite: assets.Bullet,
		OnHit:  SingleHitProjectile,
	})
}

// projectile that moves in an arc
func ArcProjectile(e *ecs.ECS, pos component.PositionComponentData) *donburi.Entry {
	bombEntity := component.NewProjectile(e.World, component.ProjectileParam{
		Vx:     5.5,
		Vy:     -4.20,
		Pos:    pos,
		Damage: 100,
		Sprite: assets.Bomb,
		OnHit:  SingleHitProjectile,
	})
	//fmt.Println(pos.String())
	entry := e.World.Entry(*bombEntity)
	entry.AddComponent(component.Acceleration)
	component.Acceleration.Set(entry, &component.AccellerationComponentData{DY: 0.25})
	entry.AddComponent(component.PositionChecker)
	component.PositionChecker.SetValue(entry, func(e *ecs.ECS) bool {
		pos := component.Position.Get(entry)
		if math.Abs(pos.Y-pos.Z) <= 13 {
			e.World.Remove(entry.Entity())
			return true
		}
		return false
	})
	return entry
}

// use this as single hit projectile. Once a projectile hit,
// apply damage then disappear. Basically the default behaviour of any projectile based attack
func SingleHitProjectile(ecs *ecs.ECS, projectile, receiver *donburi.Entry) {
	damage := component.Damage.Get(projectile).Damage
	component.Health.Get(receiver).HP -= damage
	ecs.World.Remove(projectile.Entity())
}
