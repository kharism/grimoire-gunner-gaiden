package scene

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kharism/GrimoireGunner2/scenes/system"
	"github.com/yohamta/donburi"
)

type SceneData struct {
	PlayerHP       int
	PlayerMaxHP    int
	PlayerCurrEn   int
	PlayerMaxEn    int
	PlayerEnRegen  int
	PlayerRow      int
	PlayerCol      int
	World          donburi.World
	HanashiChoices map[string]any

	Weapons []system.RenderableCaster

	MusicSeek time.Duration

	SfxVolume int //make sure the volume less than 128
	BGMVolume int

	Bg *ebiten.Image
}

func NewSceneData() *SceneData {
	return &SceneData{}
}
