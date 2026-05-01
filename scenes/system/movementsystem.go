package system

import (
	"math"

	"github.com/kharism/GrimoireGunner2/scenes/component"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/filter"
)

var GridWidth int  //Y Axis
var GridLength int //X Axis

var GridStartPointX int
var GridStartPointY int

func PlayerMovementHandler(e *ecs.ECS) {
	query := donburi.NewQuery(
		filter.Contains(
			component.Position,
			component.Velocity,
			component.PlayerTag,
		),
	)
	for entry := range query.Iter(e.World) {
		pos := component.Position.Get(entry)
		vel := component.Velocity.Get(entry)

		//fmt.Println(curCol)
		var curCol int
		var nextCol int
		if vel.X > 0 {
			curCol = int(float64(pos.X-float64(GridStartPointX)) / float64(GridLength))
			nextCol = int(float64(pos.X+vel.X-float64(GridStartPointX)) / float64(GridLength))
		} else if vel.X < 0 {
			curCol = int(math.Ceil(float64(pos.X-float64(GridStartPointX)) / float64(GridLength)))
			nextCol = int(math.Ceil(float64(pos.X+vel.X-float64(GridStartPointX)) / float64(GridLength)))
		}

		//pos.X += vel.X
		if curCol != nextCol && vel.X != 0 {
			pos.X = float64(GridStartPointX + (nextCol * GridLength))
			vel.X = 0
		} else {
			pos.X += vel.X
		}

		pos.Y += vel.Y
		pos.Z += vel.Z
		if entry.HasComponent(component.Acceleration) {
			acc := component.Acceleration.Get(entry)
			vel.X += acc.DX
			vel.Y += acc.DY
			vel.Z += acc.DZ
		}
	}
}
