package scene

import (
	"fmt"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/joelschutz/stagehand"
	"github.com/kharism/GrimoireGunner2/scenes/assets"
)

type OptionScene struct {
	sm                  *stagehand.SceneDirector[*SceneData]
	data                *SceneData
	selectedMenu        int
	musicPlayer         *assets.AudioPlayer
	currentOptionSetter optionsetter
	bgmSetter           optionsetter
	sfxSetter           optionsetter
	loopMusic           bool
}

func (r *OptionScene) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyW) {
		r.sm.ProcessTrigger(TriggerToMain)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		if r.selectedMenu < 1 {
			r.selectedMenu += 1
			r.currentOptionSetter = r.sfxSetter
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		if r.selectedMenu > 0 {
			r.selectedMenu -= 1
			r.currentOptionSetter = r.bgmSetter
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		r.currentOptionSetter.SetOptionLeft()
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		r.currentOptionSetter.SetOptionRight()
	}
	return nil
}

const startArrowX = 115
const startArrowy = 115
const distArrowY = 50

func (r *OptionScene) Draw(screen *ebiten.Image) {
	pos := ebiten.GeoM{}
	screen.DrawImage(assets.OptionBg, &ebiten.DrawImageOptions{GeoM: pos})
	pos.Reset()
	pos.Scale(0.3, 0.3)
	pos.Translate(startArrowX, startArrowy+float64(r.selectedMenu*distArrowY))
	screen.DrawImage(assets.ArrowBg, &ebiten.DrawImageOptions{GeoM: pos})

	// text stuff
	textColor := ebiten.ColorScale{}
	textColor.Scale(0, 0, 0, 1)

	bgmVol := r.musicPlayer.GetBGMVolume()
	sfxVol := r.musicPlayer.GetSfxVolume()

	pos.Reset()
	pos.Scale(1.6, 1.6)
	pos.Translate(startArrowX+20, startArrowy)
	text.Draw(screen, "BGM Volume", assets.UnispaceFace, &text.DrawOptions{
		DrawImageOptions: ebiten.DrawImageOptions{GeoM: pos, ColorScale: textColor},
	})
	pos.Translate(260, 0)
	text.Draw(screen, strconv.Itoa(bgmVol), assets.UnispaceFace, &text.DrawOptions{
		LayoutOptions:    text.LayoutOptions{PrimaryAlign: text.AlignEnd},
		DrawImageOptions: ebiten.DrawImageOptions{GeoM: pos, ColorScale: textColor},
	})
	pos.Reset()
	pos.Scale(1.6, 1.6)
	pos.Translate(startArrowX+20, startArrowy+float64(distArrowY))
	text.Draw(screen, "SFX Volume", assets.UnispaceFace, &text.DrawOptions{
		DrawImageOptions: ebiten.DrawImageOptions{GeoM: pos, ColorScale: textColor},
	})
	pos.Translate(260, 0)
	text.Draw(screen, strconv.Itoa(sfxVol), assets.UnispaceFace, &text.DrawOptions{
		LayoutOptions:    text.LayoutOptions{PrimaryAlign: text.AlignEnd},
		DrawImageOptions: ebiten.DrawImageOptions{GeoM: pos, ColorScale: textColor},
	})

}

var OptionSceneInstance = &OptionScene{}

type optionsetter interface {
	SetOptionLeft()
	SetOptionRight()
}

type bgmSetter struct {
	*assets.AudioPlayer
	sceneData *SceneData
}

func (b *bgmSetter) SetOptionLeft() {
	vol := b.AudioPlayer.GetBGMVolume()
	if vol > 0 {
		vol -= 1
		b.sceneData.BGMVolume = vol
		b.AudioPlayer.SetBgmVolume(vol)
	}
}

func (b *bgmSetter) SetOptionRight() {
	vol := b.AudioPlayer.GetBGMVolume()
	if vol < 128 {
		vol += 1
		b.sceneData.BGMVolume = vol
		b.AudioPlayer.SetBgmVolume(vol)
	}
}

type sfxSetter struct {
	*assets.AudioPlayer
	sceneData *SceneData
}

func (b *sfxSetter) SetOptionLeft() {
	vol := b.AudioPlayer.GetSfxVolume()
	if vol > 0 {
		vol -= 1
		b.sceneData.SfxVolume = vol
		b.AudioPlayer.SetSfxVolume(vol)
	}
}

func (b *sfxSetter) SetOptionRight() {
	vol := b.AudioPlayer.GetSfxVolume()
	if vol < 128 {
		vol += 1
		b.sceneData.SfxVolume = vol
		b.AudioPlayer.SetSfxVolume(vol)
	}
}
func (r *OptionScene) Load(state *SceneData, manager stagehand.SceneController[*SceneData]) {
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
	r.bgmSetter = &bgmSetter{r.musicPlayer, r.data}
	r.sfxSetter = &sfxSetter{r.musicPlayer, r.data}
	r.currentOptionSetter = r.bgmSetter
}
func (s *OptionScene) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 360
}
func (s *OptionScene) Unload() *SceneData {
	s.loopMusic = false
	s.musicPlayer.AudioPlayer().Rewind()
	s.musicPlayer.AudioPlayer().Pause()
	return s.data
}
