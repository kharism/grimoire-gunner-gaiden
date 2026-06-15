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

// just simple ticker that will remove self after 1 tick
// Entity must have position component
type DamageTicker struct {
	Ecs     *ecs.ECS
	Entity  donburi.Entity
	curTick int
}

func (c *DamageTicker) Tick() {
	if c.curTick >= 1 {
		pos := component.Position.GetValue(c.Ecs.World.Entry(c.Entity))
		assets.CreateExplosion(c.Ecs, pos)
		c.Ecs.World.Remove(c.Entity)
	} else {
		c.curTick += 1
	}

}

// use this as column hit projectile. Once a projectile hit,
// apply splash damage up and down then disappear.
func ColumnHitProjectile(ecs *ecs.ECS, projectile, receiver *donburi.Entry) {
	damage := component.Damage.GetValue(projectile)
	component.Health.Get(receiver).HP -= damage.Damage
	receiverPos := component.Position.Get(receiver)
	CreateSplashDamage(ecs, damage, component.PositionComponentData{
		X: receiverPos.X,
		Y: receiverPos.Y - 40,
		Z: receiverPos.Z - 40,
	})
	CreateSplashDamage(ecs, damage, component.PositionComponentData{
		X: receiverPos.X,
		Y: receiverPos.Y + 40,
		Z: receiverPos.Z + 40,
	})
	ecs.World.Remove(projectile.Entity())
}

func CreateSplashDamage(ecs *ecs.ECS, damage component.DamageData, position component.PositionComponentData) {
	splashDamages := ecs.World.CreateMany(1, component.Damage, component.Position, component.Ticker, component.OnHit)
	component.Position.Set(ecs.World.Entry(splashDamages[0]), &component.PositionComponentData{
		X: position.X,
		Z: position.Z,
		Y: position.Y,
	})
	component.OnHit.SetValue(ecs.World.Entry(splashDamages[0]), SingleHitProjectile)
	component.Damage.Set(ecs.World.Entry(splashDamages[0]), &component.DamageData{Damage: damage.Damage})
	damageTicker := &DamageTicker{Ecs: ecs, Entity: splashDamages[0]}
	component.Ticker.Set(ecs.World.Entry(splashDamages[0]), &component.DummyTicker{damageTicker})
	//damage := component.Damage.Get(ecs.World.Entry(projectile))

}
