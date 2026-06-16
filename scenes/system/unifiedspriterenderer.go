package system

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kharism/GrimoireGunner2/scenes/assets"
	"github.com/kharism/GrimoireGunner2/scenes/component"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/filter"
)

func UnifiedRenderer(ecs *ecs.ECS, screen *ebiten.Image) {
	drawGrid(ecs, screen)
	drawCharacters(ecs, screen)
}

var queryTile = donburi.NewQuery(
	filter.Contains(
		component.TileTag,
		component.Sprite,
	),
)

var damageTile = donburi.NewQuery(
	filter.Contains(
		component.Damage,
		component.Position,
	),
)

var dangerTile = donburi.NewQuery(
	filter.Contains(
		component.DangerTag,
		component.Position,
	),
)

func drawGrid(ecs *ecs.ECS, screen *ebiten.Image) {
	queryTile.Each(ecs.World, func(e *donburi.Entry) {
		//jj := AnimationSourceFromHP(e)
		//drawables = append(drawables, jj)
		image := component.Sprite.Get(e).Image
		bound := image.Bounds()
		pos := component.Position.Get(e)
		translate := ebiten.GeoM{}
		translate.Translate(-float64(bound.Dx())/2, -float64(bound.Dy()))
		translate.Translate(pos.X, pos.Y)
		drawOption := &ebiten.DrawImageOptions{
			GeoM: translate,
		}
		screen.DrawImage(image, drawOption)

	})
	damageTile.Each(ecs.World, func(e *donburi.Entry) {
		image := assets.GridDmg
		bound := image.Bounds()
		pos := component.Position.Get(e)
		translate := ebiten.GeoM{}
		translate.Translate(-float64(bound.Dx())/2, -float64(bound.Dy()))
		translate.Translate(pos.X, pos.Z)
		drawOption := &ebiten.DrawImageOptions{
			GeoM: translate,
		}
		screen.DrawImage(image, drawOption)
	})

	dangerTile.Each(ecs.World, func(e *donburi.Entry) {
		image := assets.GridDanger
		bound := image.Bounds()
		pos := component.Position.Get(e)
		translate := ebiten.GeoM{}
		translate.Translate(-float64(bound.Dx())/2, -float64(bound.Dy()))
		translate.Translate(pos.X, pos.Z)
		drawOption := &ebiten.DrawImageOptions{
			GeoM: translate,
		}
		screen.DrawImage(image, drawOption)
	})

}
func drawCharacters(e *ecs.ECS, screen *ebiten.Image) {
	querySprite := donburi.NewOrderedQuery[component.PositionComponentData](
		filter.And(
			filter.Contains(component.Position, component.Sprite),
			filter.Not(filter.Contains(component.TileTag)),
		),
	)
	querySprite.EachOrdered(e.World, component.Position, func(pp *donburi.Entry) {
		image := component.Sprite.Get(pp).Image
		bound := image.Bounds()
		pos := component.Position.Get(pp)
		translate := ebiten.GeoM{}
		translate.Translate(-float64(bound.Dx())/2, -float64(bound.Dy()))
		translate.Translate(pos.X, pos.Y)
		drawOption := &ebiten.DrawImageOptions{
			GeoM: translate,
		}
		screen.DrawImage(image, drawOption)
	})

}
