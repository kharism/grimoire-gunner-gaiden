package enemies

import (
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
		component.OnHit,
		component.PositionChecker,
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
	component.OnHit.SetValue(tankTicker.tankEntry, tankTicker.OnHit)
	component.PositionChecker.SetValue(tankTicker.tankEntry, tankTicker.PosChecker)
	component.Ticker.SetValue(houndEntry, component.DummyTicker{tankTicker})
}

type HoundTicker struct {
	tankEntry    *donburi.Entry
	CurrentTick  int
	CompleteTick int
	State        string
	ecs          *ecs.ECS
	OrigX        float64
	OrigY        float64
	OrigZ        float64
}

// move up and down to player's row then attack
func (c *HoundTicker) Tick() {
	var tankVY = 8.0
	var tankVX = 8.0
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
			targetY := playerPos.Y
			//targetY = math.Min(float64(component.GridStartPointY)+7*float64(component.GridWidth), targetY)
			//targetY = math.Max(float64(component.GridStartPointY)+4*float64(component.GridWidth), targetY)
			direction := targetY - pos.Y
			if direction < 0 {
				vTank.Y = -tankVY
			} else if direction > 0 {
				vTank.Y = tankVY
			} else {
				vTank.Y = 0
				c.State = "WARMUP"
				c.CompleteTick = 60
				component.Sprite.Get(c.tankEntry).Image = assets.CyberHound2
			}
		case "MOVE":
			//c.CompleteTick = -10
			targetY := playerPos.Y
			//targetY = math.Min(float64(component.GridStartPointY)+7*float64(component.GridWidth), targetY)
			//targetY = math.Max(float64(component.GridStartPointY)+4*float64(component.GridWidth), targetY)
			direction := targetY - pos.Y
			if direction < 0 {
				vTank.Y = -tankVY
			} else if direction > 0 {
				vTank.Y = tankVY
			} else {
				vTank.Y = 0
				c.State = "WARMUP"
				c.CompleteTick = 60
				component.Sprite.Get(c.tankEntry).Image = assets.CyberHound2
			}
		case "WARMUP":
			c.State = "ATTACK"
			c.CompleteTick = -10 // attack as long as we needed
			c.OrigX = pos.X
			c.OrigY = pos.Y
			c.OrigZ = pos.Z
			component.Sprite.Get(c.tankEntry).Image = assets.CyberHound3
			c.tankEntry.AddComponent(component.Damage)
			c.tankEntry.RemoveComponent(component.SingleGridMovementTag)

			component.Damage.Set(c.tankEntry, &component.DamageData{
				Damage: 10,
			})
			vTank.X = -tankVX

		}
		c.CurrentTick = 0
	}
}
func (c *HoundTicker) PosChecker(ecs *ecs.ECS) bool {
	pos := component.Position.Get(c.tankEntry)
	if pos.X < float64(component.GridStartPointX) {
		component.Velocity.Get(c.tankEntry).X = 0
		pos := component.Position.Get(c.tankEntry)
		pos.X = c.OrigX
		pos.Y = c.OrigY
		pos.Z = c.OrigZ
		c.State = "WAITING"
		component.Sprite.Get(c.tankEntry).Image = assets.CyberHound1
		c.CurrentTick = 0
		c.CompleteTick = 30
		c.tankEntry.AddComponent(component.SingleGridMovementTag)
	}
	return true
}

func (c *HoundTicker) OnHit(ecs *ecs.ECS, projectile, receiver *donburi.Entry) {
	damage := component.Damage.Get(projectile).Damage
	component.Health.Get(receiver).HP -= damage
	c.tankEntry.RemoveComponent(component.Damage)
	pos := component.Position.Get(c.tankEntry)
	pos.X = c.OrigX
	pos.Y = c.OrigY
	pos.Z = c.OrigZ
	c.State = "WAITING"
	component.Sprite.Get(c.tankEntry).Image = assets.CyberHound1
	c.tankEntry.AddComponent(component.SingleGridMovementTag)
	c.CurrentTick = 0
	c.CompleteTick = 30
	component.Velocity.Get(c.tankEntry).X = 0
}
