package assets

import (
	"bytes"
	_ "embed"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/kharism/GrimoireGunner2/scenes/component"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

//go:embed "img/vfx/slash2.png"
var sword_slash []byte
var SwordSlash *ebiten.Image
var SwordSlashFrames []*ebiten.Image

//go:embed "img/vfx/explosion.png"
var explosion []byte
var explosion_all *ebiten.Image
var ExplosionFrames []*ebiten.Image

func init() {
	if SwordSlash == nil {
		imgReader := bytes.NewReader(sword_slash)
		SwordSlash, _, _ = ebitenutil.NewImageFromReader(imgReader)
		swordSlashSize := SwordSlash.Bounds()
		for i := 0; i < swordSlashSize.Dx(); i += 80 {
			SwordSlashFrames = append(SwordSlashFrames, SwordSlash.SubImage(image.Rect(i, 0, i+80, 120)).(*ebiten.Image))
		}

		imgReader = bytes.NewReader(explosion)
		explosion_all, _, _ = ebitenutil.NewImageFromReader(imgReader)
		explosionSize := explosion_all.Bounds()
		for i := 0; i < explosionSize.Dx(); i += 75 {
			ExplosionFrames = append(ExplosionFrames, explosion_all.SubImage(image.Rect(i, 0, i+75, 75)).(*ebiten.Image))
		}
	}
}

type ExplosionTicker struct {
	CurrentTick    int
	explosionEntry *donburi.Entry
	world          donburi.World
	index          int
}

func (c *ExplosionTicker) Tick() {
	c.CurrentTick += 3
	if c.CurrentTick%3 == 0 {
		c.index += 1
		if c.index == 11 {
			//c.explosionEntry = nil
			c.world.Remove(c.explosionEntry.Entity())

		} else {
			if c.index < 11 {
				component.Sprite.Get(c.explosionEntry).Image = ExplosionFrames[c.index]
			}
		}
	}

}

func CreateExplosion(ecs *ecs.ECS, position component.PositionComponentData) {
	entityExplosion := ecs.World.Create(
		component.Position,
		component.Sprite,
		component.Ticker,
	)
	explosionEntry := ecs.World.Entry(entityExplosion)
	component.Position.Set(explosionEntry, &position)
	component.Sprite.Set(explosionEntry, &component.SpriteData{Image: ExplosionFrames[0]})
	component.Ticker.Set(explosionEntry, &component.DummyTicker{
		&ExplosionTicker{CurrentTick: 0, explosionEntry: explosionEntry, world: ecs.World},
	})

}
