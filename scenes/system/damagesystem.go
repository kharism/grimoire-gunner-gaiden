package system

import (
	"math"

	"github.com/kharism/GrimoireGunner2/scenes/assets"
	"github.com/kharism/GrimoireGunner2/scenes/component"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/filter"
)

func CheckIsHit(targetPos, hazardPos component.PositionComponentData) bool {
	return math.Abs(targetPos.X-hazardPos.X) < float64(component.GridLength) &&
		math.Abs(targetPos.Z-hazardPos.Z) < float64(component.GridWidth) &&
		math.Abs(targetPos.Y-hazardPos.Y) < 40
}
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
			if CheckIsHit(targetPos, hazardPos) {
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
					assets.CreateExplosion(ecs, pos)
					ecs.World.Remove(target.Entity())
				}
			}
		}
	}
}
