package system

import (
	"fmt"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/kharism/GrimoireGunner2/scenes/component"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/filter"
)

var MAX_COLUMN = 3

// check whether the zone player going to is available
func isLegalMove(pos component.PositionComponentData, world donburi.World) bool {
	if pos.X < 0 || pos.X > float64(component.GridStartPointX+(MAX_COLUMN*component.GridLength)) {
		return false
	}
	if pos.Y > float64(component.GridStartPointY+3*component.GridWidth) || pos.Y < float64(component.GridStartPointY) {
		return false
	}
	query := donburi.NewQuery(filter.And(
		filter.Contains(
			component.Position,
			component.Health,
		), filter.Not(filter.Contains(component.PlayerTag)),
	))

	for kk := range query.Iter(world) {

		//fmt.Println(component.Health.Get(kk).Name)
		cc := component.Position.Get(kk)
		if cc.X == pos.X && cc.Z == pos.Z {
			fmt.Println(component.Health.Get(kk).Name)
			return false
		}
	}
	return true
}
func detectMovementKey(world donburi.World) {
	xVelo := 14.0
	yVelo := xVelo / 2
	query := donburi.NewQuery(filter.Contains(component.PlayerTag))

	playerEntry, _ := query.First(world) //world.Entry(c.player)
	posData := component.Position.GetValue(playerEntry)
	vData := component.Velocity.GetValue(playerEntry)

	if !vData.IsMoving() {
		if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
			posData.X += float64(component.GridLength)
			if isLegalMove(posData, world) {
				//playerEntry := world.Entry(playerEntry)
				component.Velocity.Get(playerEntry).X = xVelo
			}

		} else if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
			posData.X -= float64(component.GridLength)
			if isLegalMove(posData, world) {
				//playerEntry := c.world.Entry(c.player)
				component.Velocity.Get(playerEntry).X = -xVelo
			}

		} else if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
			posData.Z += float64(component.GridWidth)
			if isLegalMove(posData, world) {
				component.Velocity.Get(playerEntry).Y = yVelo
				component.Velocity.Get(playerEntry).Z = yVelo
			}

		} else if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
			posData.Z -= float64(component.GridWidth)
			if isLegalMove(posData, world) {
				//playerEntry := c.world.Entry(c.player)
				component.Velocity.Get(playerEntry).Y = -yVelo
				component.Velocity.Get(playerEntry).Z = -yVelo
			}

		}
	}
}
func NonPlayerMovementHandler(e *ecs.ECS) {
	query := donburi.NewQuery(
		filter.And(
			filter.Contains(
				component.Position,
				component.Velocity,
			),
			filter.Not(filter.Contains(
				component.PlayerTag,
			)),
		),
	)
	listToRemove := []donburi.Entity{}
	for entry := range query.Iter(e.World) {
		pos := component.Position.Get(entry)
		vel := component.Velocity.Get(entry)
		if entry.HasComponent(component.SingleGridMovementTag) {
			var curCol int
			var nextCol int
			if vel.X > 0 {
				curCol = int(float64(pos.X-float64(component.GridStartPointX)) / float64(component.GridLength))
				nextCol = int(float64(pos.X+vel.X-float64(component.GridStartPointX)) / float64(component.GridLength))
			} else if vel.X < 0 {
				curCol = int(math.Ceil(float64(pos.X-float64(component.GridStartPointX)) / float64(component.GridLength)))
				nextCol = int(math.Ceil(float64(pos.X+vel.X-float64(component.GridStartPointX)) / float64(component.GridLength)))
			}

			var curRow int
			var nextRow int

			if vel.Y > 0 {
				curRow = int(float64(pos.Y-float64(component.GridStartPointY)) / float64(component.GridWidth))
				nextRow = int(float64(pos.Y+vel.Y-float64(component.GridStartPointY)) / float64(component.GridWidth))
			} else if vel.Y < 0 {
				curRow = int(math.Ceil(float64(pos.Y-float64(component.GridStartPointY)) / float64(component.GridWidth)))
				nextRow = int(math.Ceil(float64(pos.Y+vel.Y-float64(component.GridStartPointY)) / float64(component.GridWidth)))
			}

			//pos.X += vel.X
			if curCol != nextCol && vel.X != 0 {
				pos.X = float64(component.GridStartPointX + (nextCol * component.GridLength))
				vel.X = 0
			} else {
				pos.X += vel.X
			}

			if curRow != nextRow && vel.Y != 0 {
				pos.Y = float64(component.GridStartPointY + (nextRow * component.GridWidth))
				pos.Z = float64(component.GridStartPointY + (nextRow * component.GridWidth))
				vel.Y = 0
				vel.Z = 0
			} else {
				pos.Y += vel.Y
				pos.Z += vel.Z
			}

			//pos.Y += vel.Y

			if entry.HasComponent(component.Acceleration) {
				acc := component.Acceleration.Get(entry)
				vel.X += acc.DX
				vel.Y += acc.DY
				vel.Z += acc.DZ
			}
		} else {
			pos.X += vel.X
			pos.Y += vel.Y
			pos.Z += vel.Z
			if entry.HasComponent(component.Acceleration) {
				acc := component.Acceleration.Get(entry)
				vel.X += acc.DX
				vel.Y += acc.DY
				vel.Z += acc.DZ
			}
			if pos.X <= 0 || pos.X >= 640 || pos.Y > 360 || pos.Y < -40 {
				listToRemove = append(listToRemove, entry.Entity())
			}
		}

	}
	for _, i := range listToRemove {
		e.World.Remove(i)
	}
}
func PlayerMovementHandler(e *ecs.ECS) {
	detectMovementKey(e.World)
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
			curCol = int(float64(pos.X-float64(component.GridStartPointX)) / float64(component.GridLength))
			nextCol = int(float64(pos.X+vel.X-float64(component.GridStartPointX)) / float64(component.GridLength))
		} else if vel.X < 0 {
			curCol = int(math.Ceil(float64(pos.X-float64(component.GridStartPointX)) / float64(component.GridLength)))
			nextCol = int(math.Ceil(float64(pos.X+vel.X-float64(component.GridStartPointX)) / float64(component.GridLength)))
		}

		var curRow int
		var nextRow int

		if vel.Y > 0 {
			curRow = int(float64(pos.Y-float64(component.GridStartPointY)) / float64(component.GridWidth))
			nextRow = int(float64(pos.Y+vel.Y-float64(component.GridStartPointY)) / float64(component.GridWidth))
		} else if vel.Y < 0 {
			curRow = int(math.Ceil(float64(pos.Y-float64(component.GridStartPointY)) / float64(component.GridWidth)))
			nextRow = int(math.Ceil(float64(pos.Y+vel.Y-float64(component.GridStartPointY)) / float64(component.GridWidth)))
		}

		//pos.X += vel.X
		if curCol != nextCol && vel.X != 0 {
			pos.X = float64(component.GridStartPointX + (nextCol * component.GridLength))
			vel.X = 0
		} else {
			pos.X += vel.X
		}

		if curRow != nextRow && vel.Y != 0 {
			pos.Y = float64(component.GridStartPointY + (nextRow * component.GridWidth))
			pos.Z = float64(component.GridStartPointY + (nextRow * component.GridWidth))
			vel.Y = 0
			vel.Z = 0
		} else {
			pos.Y += vel.Y
			pos.Z += vel.Z
		}

		//pos.Y += vel.Y

		if entry.HasComponent(component.Acceleration) {
			acc := component.Acceleration.Get(entry)
			vel.X += acc.DX
			vel.Y += acc.DY
			vel.Z += acc.DZ
		}
	}
}
