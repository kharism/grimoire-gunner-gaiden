package system

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/kharism/GrimoireGunner2/scenes/assets"
	"github.com/kharism/GrimoireGunner2/scenes/component"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/filter"
)

var hpquery = donburi.NewQuery(
	filter.Contains(component.Health),
)

func DrawHP(ecs *ecs.ECS, screen *ebiten.Image) {
	for e := range hpquery.Iter(ecs.World) {
		pos := component.Position.GetValue(e)
		translate := ebiten.GeoM{}
		translate.Translate(pos.X, pos.Y)
		op := &text.DrawOptions{
			LayoutOptions: text.LayoutOptions{
				PrimaryAlign: text.AlignCenter,
			},
			DrawImageOptions: ebiten.DrawImageOptions{
				GeoM: translate,
			},
		}
		hp := component.Health.Get(e).HP
		text.Draw(screen, fmt.Sprintf("%d", hp), assets.PixelOperatorFace, op)
	}
}
