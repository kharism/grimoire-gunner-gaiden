package system

import (
	"github.com/kharism/GrimoireGunner2/scenes/component"
	"github.com/yohamta/donburi"

	//"github.com/yohamta/donburi/component"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/filter"
)

func PositionCheckerSystem(e *ecs.ECS) {
	query := donburi.NewQuery(filter.Contains(
		component.PositionChecker,
	))

	for i := range query.Iter(e.World) {
		component.PositionChecker.GetValue(i)(e)
	}

}
