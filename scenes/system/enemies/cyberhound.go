package enemies

import (
	"math"

	"github.com/kharism/GrimoireGunner2/scenes/assets"
	"github.com/kharism/GrimoireGunner2/scenes/component"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/filter"
)

func LoadHound(e *ecs.ECS, row, col int) {
	houndEntity := e.World.Create(
		component.Health,
		component.Position,
		component.Sprite,
		component.Ticker,
		component.Velocity,
		component.SingleGridMovementTag,
	)
	houndEntry := e.World.Entry(houndEntity)
	component.Health.Set(houndEntry, &component.HealthData{
		HP:    100,
		MaxHP: 100,
		Name:  "Hound",
	})
	component.Sprite.Set(houndEntry, &component.SpriteData{
		Image: assets.CyberHound1,
	})
	gridStartX := col
	gridStartY := row
	component.Position.Set(houndEntry, &component.PositionComponentData{
		X: StartX + float64(gridStartX)*GridLength,
		Z: StartY + float64(gridStartY)*GridWidth,
		Y: StartY + float64(gridStartY)*GridWidth,
	})
	component.Velocity.Set(houndEntry, &component.VelocityComponentData{
		X: 0,
		Y: 0,
		Z: 0,
	})
	tankTicker := &HoundTicker{tankEntry: houndEntry, CurrentTick: 0, CompleteTick: 60, State: "WAITING", ecs: e}
	component.Ticker.SetValue(houndEntry, component.DummyTicker{tankTicker})
}

type HoundTicker struct {
	tankEntry    *donburi.Entry
	CurrentTick  int
	CompleteTick int
	State        string
	ecs          *ecs.ECS
}

// move up and down to player's row then attack
func (c *HoundTicker) Tick() {
	var tankVX = 17.0
	c.CurrentTick += 3
	if c.CurrentTick > c.CompleteTick {
		c.CurrentTick = c.CompleteTick
	}
	playerQ := donburi.NewQuery(
		filter.Contains(component.PlayerTag),
	)
	playerEntry, ok := playerQ.First(c.ecs.World)
	if !ok {
		return
	}
	playerPos := component.Position.Get(playerEntry)
	pos := component.Position.Get(c.tankEntry)
	vTank := component.Velocity.Get(c.tankEntry)
	if c.CurrentTick == c.CompleteTick {
		switch c.State {
		case "WAITING":
			c.State = "MOVE"
			c.CompleteTick = -10
			targetY := playerPos.Y + float64(4*component.GridWidth)
			targetY = math.Min(float64(component.GridStartPointY)+7*float64(component.GridLength), targetY)
			targetY = math.Max(float64(component.GridStartPointY)+4*float64(component.GridLength), targetY)
			direction := targetY - pos.Y
			if direction > 0 {
				vTank.Y = tankVX
			} else if direction < 0 {
				vTank.Y = -tankVX
			} else {
				vTank.Y = 0
				c.State = "WARMUP"
				c.CompleteTick = 60
				component.Sprite.Get(c.tankEntry).Image = assets.CyberHound2
			}
		case "MOVE":
			//c.CompleteTick = -10
			targetY := playerPos.Y + float64(4*component.GridWidth)
			targetY = math.Min(float64(component.GridStartPointY)+7*float64(component.GridLength), targetY)
			targetY = math.Max(float64(component.GridStartPointY)+4*float64(component.GridLength), targetY)
			direction := targetY - pos.Y
			if direction > 0 {
				vTank.Y = tankVX
			} else if direction < 0 {
				vTank.Y = -tankVX
			} else {
				vTank.Y = 0
				c.State = "WARMUP"
				c.CompleteTick = 60
				component.Sprite.Get(c.tankEntry).Image = assets.CyberHound2
			}
		}
	}
}
