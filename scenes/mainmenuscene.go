package scene

import (
	"fmt"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/joelschutz/stagehand"
	"github.com/kharism/GrimoireGunner2/scenes/assets"
)

type MainMenuScene struct {
	sm           *stagehand.SceneDirector[*SceneData]
	data         *SceneData
	selectedMenu int
	musicPlayer  *assets.AudioPlayer
	loopMusic    bool
}

var menus = []string{
	"New Game",
	"Story",
	"Options",
	"Exit",
}
var menusFunc = []func(){
	StartGame,
	Story,
	Options,
	Exit,
}

func StartGame() {

}
func Story() {

}
func Options() {
	MainMenuInstance.sm.ProcessTrigger(TriggerToOption)
}
func Exit() {
	os.Exit(0)
}

var MainMenuInstance = &MainMenuScene{}

func (r *MainMenuScene) Update() error {
	if r.loopMusic && !r.musicPlayer.AudioPlayer().IsPlaying() {
		r.musicPlayer.AudioPlayer().Rewind()
		r.musicPlayer.AudioPlayer().Play()
	}
	if r.musicPlayer != nil {
		r.musicPlayer.Update()
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		r.selectedMenu += 1
		if r.selectedMenu == len(menus) {
			r.selectedMenu -= 1
		}
		r.musicPlayer.QueueSFX(assets.MenuMove, assets.TypeOgg)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		r.selectedMenu -= 1
		if r.selectedMenu == -1 {
			r.selectedMenu += 1
		}
		r.musicPlayer.QueueSFX(assets.MenuMove, assets.TypeOgg)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
		menusFunc[r.selectedMenu]()
	}
	return nil
}
func (r *MainMenuScene) Draw(screen *ebiten.Image) {
	textColor := ebiten.ColorScale{}
	textColor.Scale(0, 0, 0, 1)
	for idx, i := range menus {
		pos := ebiten.GeoM{}
		if idx == r.selectedMenu {
			pos.Scale(1.5, 1)
		}
		pos.Translate(0, float64(150+40*idx))

		screen.DrawImage(assets.MenuButtonBg, &ebiten.DrawImageOptions{GeoM: pos})
		pos.Reset()
		pos.Scale(1.6, 1.6)
		pos.Translate(20, float64(150+40*idx))

		text.Draw(screen, i, assets.UnispaceFace, &text.DrawOptions{

			DrawImageOptions: ebiten.DrawImageOptions{GeoM: pos, ColorScale: textColor},
		})
	}
}
func (r *MainMenuScene) Load(state *SceneData, manager stagehand.SceneController[*SceneData]) {
	r.sm = manager.(*stagehand.SceneDirector[*SceneData]) // This type assertion is important
	r.data = state
	if r.musicPlayer == nil {
		var err error
		r.musicPlayer, err = assets.NewAudioPlayer(assets.Menumusic, assets.TypeOgg, state.BGMVolume, state.SfxVolume)
		if err != nil {
			fmt.Println(err.Error())
		}
		r.musicPlayer.AudioPlayer().Play()
	} else {
		// s.musicPlayer.audioPlayer.Rewind()
		r.musicPlayer.AudioPlayer().Play()
	}
}
func (s *MainMenuScene) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 360
}
func (s *MainMenuScene) Unload() *SceneData {
	s.loopMusic = false
	s.musicPlayer.AudioPlayer().Rewind()
	s.musicPlayer.AudioPlayer().Pause()
	return s.data
}
