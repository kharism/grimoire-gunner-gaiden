package system

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/kharism/GrimoireGunner2/scenes/component"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/filter"
)

var GridWidth int  //Y Axis
var GridLength int //X Axis

var GridStartPointX int
var GridStartPointY int

func isLegalMove(pos component.PositionComponentData) bool {
	if pos.X < 0 || pos.X > float64(GridStartPointX+(7*GridLength)) {
		return false
	}
	if pos.Y > float64(GridStartPointY+3*GridWidth) || pos.Y < float64(GridStartPointY) {
		return false
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
			posData.X += float64(GridLength)
			if isLegalMove(posData) {
				//playerEntry := world.Entry(playerEntry)
				component.Velocity.Get(playerEntry).X = xVelo
			}

		} else if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
			posData.X -= float64(GridLength)
			if isLegalMove(posData) {
				//playerEntry := c.world.Entry(c.player)
				component.Velocity.Get(playerEntry).X = -xVelo
			}

		} else if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
			posData.Y += float64(GridWidth)
			if isLegalMove(posData) {
				component.Velocity.Get(playerEntry).Y = yVelo
				component.Velocity.Get(playerEntry).Z = yVelo
			}

		} else if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
			posData.Y -= float64(GridWidth)
			if isLegalMove(posData) {
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
	for entry := range query.Iter(e.World) {
		pos := component.Position.Get(entry)
		vel := component.Velocity.Get(entry)
		pos.X += vel.X
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
			curCol = int(float64(pos.X-float64(GridStartPointX)) / float64(GridLength))
			nextCol = int(float64(pos.X+vel.X-float64(GridStartPointX)) / float64(GridLength))
		} else if vel.X < 0 {
			curCol = int(math.Ceil(float64(pos.X-float64(GridStartPointX)) / float64(GridLength)))
			nextCol = int(math.Ceil(float64(pos.X+vel.X-float64(GridStartPointX)) / float64(GridLength)))
		}

		var curRow int
		var nextRow int

		if vel.Y > 0 {
			curRow = int(float64(pos.Y-float64(GridStartPointY)) / float64(GridWidth))
			nextRow = int(float64(pos.Y+vel.Y-float64(GridStartPointY)) / float64(GridWidth))
		} else if vel.Y < 0 {
			curRow = int(math.Ceil(float64(pos.Y-float64(GridStartPointY)) / float64(GridWidth)))
			nextRow = int(math.Ceil(float64(pos.Y+vel.Y-float64(GridStartPointY)) / float64(GridWidth)))
		}

		//pos.X += vel.X
		if curCol != nextCol && vel.X != 0 {
			pos.X = float64(GridStartPointX + (nextCol * GridLength))
			vel.X = 0
		} else {
			pos.X += vel.X
		}

		if curRow != nextRow && vel.Y != 0 {
			pos.Y = float64(GridStartPointY + (nextRow * GridWidth))
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
