package component

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kharism/hanashi/core"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

type SpriteData struct {
	Image *ebiten.Image
	Scale *core.ScaleParam
}

var Sprite = donburi.NewComponentType[SpriteData]()

type SpriteTicker struct {
	World    donburi.World
	Entry    *donburi.Entry
	Sprites  []*ebiten.Image
	Active   bool
	TimeTick int
	frame    int
	curTick  int
}

func (t *SpriteTicker) Tick() {
	if t.Active {
		playerQuery := donburi.NewQuery(filter.Contains(
			PlayerTag,
		))
		playerE, ok := playerQuery.FirstEntity(t.World)
		if ok {
			state := State.GetValue(playerE)
			if state == STATE_STANDING {
				t.curTick += 1
				if t.curTick == t.TimeTick {
					t.frame += 1
					Sprite.Get(t.Entry).Image = t.Sprites[t.frame%len(t.Sprites)]
					t.curTick = 0
				}
			}
		}

	}

}
