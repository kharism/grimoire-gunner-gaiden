package enemies

import (
	"math"

	"github.com/kharism/GrimoireGunner2/scenes/assets"
	"github.com/kharism/GrimoireGunner2/scenes/component"
	"github.com/kharism/GrimoireGunner2/scenes/system/weapons"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/filter"
)

func LoadTank(e *ecs.ECS, row, col int) {
	tankEntity := e.World.Create(
		component.Health,
		component.Position,
		component.Sprite,
		component.Ticker,
		component.Velocity,
		component.SingleGridMovementTag,
	)
	tankEntry := e.World.Entry(tankEntity)
	component.Health.Set(tankEntry, &component.HealthData{
		HP:    100,
		MaxHP: 100,
		Name:  "Tank",
	})
	component.Sprite.Set(tankEntry, &component.SpriteData{
		Image: assets.TankSprite1,
	})
	gridStartX := col
	gridStartY := row
	component.Position.Set(tankEntry, &component.PositionComponentData{
		X: StartX + float64(gridStartX)*GridLength,
		Z: StartY + float64(gridStartY)*GridWidth,
		Y: StartY + float64(gridStartY)*GridWidth,
	})
	component.Velocity.Set(tankEntry, &component.VelocityComponentData{
		X: 0,
		Y: 0,
		Z: 0,
	})
	tankTicker := &TankTicker{tankEntry: tankEntry, CurrentTick: 0, CompleteTick: 60, State: "WAITING", ecs: e}
	component.Ticker.SetValue(tankEntry, component.DummyTicker{tankTicker})

}

type TankTicker struct {
	tankEntry    *donburi.Entry
	CurrentTick  int
	CompleteTick int
	State        string
	ecs          *ecs.ECS
}

// the behaviour of tank in simple state machine
// waiting->move->warmup->attack->wait
// move will be forward backward
func (c *TankTicker) Tick() {
	var tankVX = 14.0
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
			// make the target negative so the state change only when it
			// reached destination
			c.CompleteTick = -10
			targetX := playerPos.X + float64(4*component.GridLength)
			targetX = math.Min(float64(component.GridStartPointX)+7*float64(component.GridLength), targetX)
			targetX = math.Max(float64(component.GridStartPointX)+4*float64(component.GridLength), targetX)
			direction := targetX - pos.X
			//vTank := component.Velocity.Get(c.tankEntry)
			if direction > 0 {
				vTank.X = tankVX
			} else if direction < 0 {
				vTank.X = -tankVX
			} else {
				vTank.X = 0
				c.State = "WARMUP"
				c.CompleteTick = 90
				component.Sprite.Get(c.tankEntry).Image = assets.TankSprite2
			}

		case "MOVE":
			targetX := playerPos.X + float64(4*component.GridLength)
			targetX = math.Min(float64(component.GridStartPointX)+7*float64(component.GridLength), targetX)
			targetX = math.Max(float64(component.GridStartPointX)+4*float64(component.GridLength), targetX)
			direction := targetX - pos.X
			//vTank := component.Velocity.Get(c.tankEntry)
			if direction > 0 {
				vTank.X = tankVX
			} else if direction < 0 {
				vTank.X = -tankVX
			} else {
				c.State = "WARMUP"
				c.CompleteTick = 20
				component.Sprite.Get(c.tankEntry).Image = assets.TankSprite2
			}
		case "WARMUP":
			c.State = "ATTACK"
			c.CompleteTick = 27
			//component.Sprite.Get(c.tankEntry).Image = assets.TankSprite1
			projectile := weapons.ArcProjectile(c.ecs, component.PositionComponentData{
				X: pos.X - float64(component.GridLength),
				Y: pos.Y - 90,
				Z: pos.Z,
			})
			component.Velocity.Get(projectile).X *= -1
		case "ATTACK":
			c.State = "WAITING"
			c.CompleteTick = 60
			component.Sprite.Get(c.tankEntry).Image = assets.TankSprite2
		}
		c.CurrentTick = 0
		//c.CompleteTick = 100
	}
	if c.State == "MOVE" {
		if vTank.X == 0 {
			targetX := playerPos.X + float64(4*component.GridLength)
			direction := targetX - pos.X
			//vTank := component.Velocity.Get(c.tankEntry)
			if direction > 0 {
				vTank.X = 14.0
			} else if direction < 0 {
				vTank.X = -14.0
			} else {
				// trigger state change
				c.CurrentTick = c.CompleteTick - 3
			}

		}

	}
	if c.State == "ATTACK" {
		curFrame := assets.TankAnimFrames[c.CurrentTick/9]
		component.Sprite.Get(c.tankEntry).Image = curFrame
	}
}
