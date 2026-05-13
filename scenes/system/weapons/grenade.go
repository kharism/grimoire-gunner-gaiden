package weapons

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kharism/GrimoireGunner2/scenes/assets"
	"github.com/kharism/GrimoireGunner2/scenes/component"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/filter"
)

type Grenade struct {
	CurrentTick  int
	CompleteTick int
}

func NewGrenade() *Grenade {
	return &Grenade{CurrentTick: 100, CompleteTick: 100}
}
func (c *Grenade) Cast(e *ecs.ECS) {

	playerQuery := donburi.NewQuery(filter.Contains(
		component.PlayerTag,
	))
	c.CurrentTick = 0
	playerE, _ := playerQuery.FirstEntity(e.World)
	playerPos := component.Position.GetValue(playerE)
	ArcProjectile(e, component.PositionComponentData{X: playerPos.X + float64(component.GridLength), Y: playerPos.Y - 90, Z: playerPos.Z})
}
func (c *Grenade) GetIcon() *ebiten.Image {
	return assets.BombIcon
}
func (c *Grenade) GetDamage() int {
	return 100
}
func (c *Grenade) Tick() {
	c.CurrentTick += 3
	if c.CurrentTick > c.CompleteTick {
		c.CurrentTick = c.CompleteTick
	}
}
func (c *Grenade) GetCooldownProgress() float64 {
	return (float64(c.CurrentTick) / float64(c.CompleteTick))
}
