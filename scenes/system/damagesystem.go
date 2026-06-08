package system

import (
	"math"

	"github.com/kharism/GrimoireGunner2/scenes/assets"
	"github.com/kharism/GrimoireGunner2/scenes/component"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/filter"
)

func DamageSystemHandler(ecs *ecs.ECS) {
	damageQuery := donburi.NewQuery(
		filter.Contains(
			component.Damage,
			component.OnHit,
			component.Position,
		),
	)
	healthQuery := donburi.NewQuery(
		filter.Contains(
			component.Position,
			component.Health,
		),
	)
	// for target := range healthQuery.Iter(ecs.World) {
	// 	health := component.Health.Get(target)
	// 	fmt.Println(health.Name)
	// }
	for hazard := range damageQuery.Iter(ecs.World) {
		validTargets := []*donburi.Entry{}
		for target := range healthQuery.Iter(ecs.World) {
			targetPos := component.Position.GetValue(target)
			hazardPos := component.Position.GetValue(hazard)
			if math.Abs(targetPos.X-hazardPos.X) < float64(component.GridLength) &&
				math.Abs(targetPos.Z-hazardPos.Z) < float64(component.GridWidth) &&
				math.Abs(targetPos.Y-hazardPos.Y) < 40 {
				// health := component.Health.Get(target)
				validTargets = append(validTargets, target)
				// fmt.Println(health.Name)
				// fmt.Println(targetPos.String())
				// fmt.Println(hazardPos.String())

			}
		}
		for _, target := range validTargets {
			if !ecs.World.Valid(hazard.Entity()) {
				break
			}
			if ecs.World.Valid(target.Entity()) {
				component.OnHit.GetValue(hazard)(ecs, hazard, target)
				if component.Health.Get(target).HP <= 0 {
					pos := component.Position.GetValue(target)
					createExplosion(ecs, pos)
					ecs.World.Remove(target.Entity())
				}
			}
		}
	}
}

type ExplosionTicker struct {
	CurrentTick    int
	explosionEntry *donburi.Entry
	world          donburi.World
	index          int
}

func (c *ExplosionTicker) Tick() {
	c.CurrentTick += 3
	if c.CurrentTick%3 == 0 {
		c.index += 1
		if c.index == 11 {
			//c.explosionEntry = nil
			c.world.Remove(c.explosionEntry.Entity())

		} else {
			if c.index < 11 {
				component.Sprite.Get(c.explosionEntry).Image = assets.ExplosionFrames[c.index]
			}
		}
	}

}

func createExplosion(ecs *ecs.ECS, position component.PositionComponentData) {
	entityExplosion := ecs.World.Create(
		component.Position,
		component.Sprite,
		component.Ticker,
	)
	explosionEntry := ecs.World.Entry(entityExplosion)
	component.Position.Set(explosionEntry, &position)
	component.Sprite.Set(explosionEntry, &component.SpriteData{Image: assets.ExplosionFrames[0]})
	component.Ticker.Set(explosionEntry, &component.DummyTicker{
		&ExplosionTicker{CurrentTick: 0, explosionEntry: explosionEntry, world: ecs.World},
	})

}
