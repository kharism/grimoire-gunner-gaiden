package weapons

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kharism/GrimoireGunner2/scenes/assets"
	"github.com/kharism/GrimoireGunner2/scenes/component"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/filter"
)

type WideSword struct {
	CurrentTick  int
	CompleteTick int
}

func NewWdSword() *WideSword {
	return &WideSword{CurrentTick: 100, CompleteTick: 100}
}
func (c *WideSword) Cast(e *ecs.ECS) {
	playerQuery := donburi.NewQuery(filter.Contains(
		component.PlayerTag,
	))
	c.CurrentTick = 0
	playerE, _ := playerQuery.FirstEntity(e.World)
	playerPos := component.Position.GetValue(playerE)
	component.Sprite.Get(playerE).Image = assets.SvenSprite3

	slashAnimEntity := e.World.Create(component.Sprite, component.Position)
	slashAnimEntry := e.World.Entry(slashAnimEntity)

	component.Sprite.Set(slashAnimEntry, &component.SpriteData{Image: assets.SwordSlashFrames[0]})
	component.Position.Set(slashAnimEntry, &component.PositionComponentData{
		Z: playerPos.Z,
		X: playerPos.X + 80,
		Y: playerPos.Y + 30,
	})
	hazardEntries := []*donburi.Entry{}
	for i := -1; i < 2; i++ {
		// check out of bound
		if playerPos.Z+float64(40*i) < float64(component.GridStartPointY) || playerPos.Z+float64(40*i) > float64(component.GridStartPointY+40*4) {
			continue
		}
		hazardEntity := e.World.Create(
			component.Damage,
			component.OnHit,
			component.Position,
		)
		hazardEntry := e.World.Entry(hazardEntity)
		component.Damage.Set(hazardEntry, &component.DamageData{Damage: c.GetDamage()})
		component.Position.Set(hazardEntry, &component.PositionComponentData{
			X: playerPos.X + 80,
			Y: playerPos.Y + float64(40*i),
			Z: playerPos.Z + float64(40*i),
		})
		component.OnHit.SetValue(hazardEntry, SingleHitProjectile)
		hazardEntries = append(hazardEntries, hazardEntry)
	}

	ent := e.World.Create(component.Ticker)
	ent2 := e.World.Entry(ent)
	wdSlashAnim := &WdSlashAnim{TickEntry: ent2, AnimeEntry: slashAnimEntry, world: e.World, hazardEntries: hazardEntries}
	component.Ticker.SetValue(ent2, component.DummyTicker{wdSlashAnim})

}
func (c *WideSword) GetIcon() *ebiten.Image {
	return assets.SwordIcon
}
func (c *WideSword) GetDamage() int {
	return 90
}
func (c *WideSword) Tick() {
	c.CurrentTick += 7
	if c.CurrentTick > c.CompleteTick {
		c.CurrentTick = c.CompleteTick
	}
}
func (c *WideSword) GetCooldownProgress() float64 {
	return (float64(c.CurrentTick) / float64(c.CompleteTick))
}

type WdSlashAnim struct {
	AnimeEntry    *donburi.Entry
	TickEntry     *donburi.Entry
	world         donburi.World
	hazardEntries []*donburi.Entry
	CurrentTick   int
	CompleteTick  int
}

func (c *WdSlashAnim) Tick() {
	c.CurrentTick += 3
	if c.CurrentTick >= 15 {
		component.Sprite.Get(c.AnimeEntry).Image = assets.SwordSlashFrames[4]
	}
	if c.CurrentTick >= 12 {
		component.Sprite.Get(c.AnimeEntry).Image = assets.SwordSlashFrames[4]
	} else if c.CurrentTick >= 9 {
		component.Sprite.Get(c.AnimeEntry).Image = assets.SwordSlashFrames[3]
	} else if c.CurrentTick >= 6 {
		component.Sprite.Get(c.AnimeEntry).Image = assets.SwordSlashFrames[3]
	} else if c.CurrentTick >= 3 {
		component.Sprite.Get(c.AnimeEntry).Image = assets.SwordSlashFrames[2]
	}

	if c.CurrentTick >= 18 {
		c.world.Remove(c.AnimeEntry.Entity())
		c.world.Remove(c.TickEntry.Entity())
		for _, i := range c.hazardEntries {
			if c.world.Valid(i.Entity()) {
				c.world.Remove(i.Entity())
			}
		}
	}
}
