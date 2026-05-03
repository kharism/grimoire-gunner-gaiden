package system

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/kharism/GrimoireGunner2/scenes/component"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/filter"
)

func DrawDebug(ecs *ecs.ECS, screen *ebiten.Image) {
	query := donburi.NewQuery(
		filter.Contains(component.PlayerTag),
		// filter.Contains(component.Fx),
	)
	for entry := range query.Iter(ecs.World) {
		pos := component.Position.GetValue(entry)
		ebitenutil.DebugPrint(screen, fmt.Sprintf("Player Pos %s", pos.String()))
	}
}
