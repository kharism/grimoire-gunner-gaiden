package system

import (
	"time"

	"github.com/kharism/GrimoireGunner2/scenes/component"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/filter"
)

var LastTick time.Time

var DelayTick, _ = time.ParseDuration("80ms")

func Tick(e *ecs.ECS) {
	if time.Since(LastTick) > DelayTick {
		LastTick = time.Now()
		query := donburi.NewQuery(
			filter.Contains(component.Ticker),
		)
		for i := range query.Iter(e.World) {
			component.Ticker.GetValue(i).H.Tick()
		}
	}

}
