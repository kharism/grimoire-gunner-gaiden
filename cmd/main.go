package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/joelschutz/stagehand"
	scene "github.com/kharism/GrimoireGunner2/scenes"
)

const (
	screenWidth  = 640
	screenHeight = 360
)

type Game struct {
	count int
}

func (g *Game) Update() error {
	return nil
}
func (g *Game) Draw(screen *ebiten.Image) {

}
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 360
}

// return width and height of the scene
func (g *Game) GetLayout() (width, height int) {
	return 640, 360
}

// return the starting text position where the box containing name of the character appear on the scene
// return negative number if no such box needed
func (g *Game) GetNamePosition() (x, y int) {
	return 128, 600 - 150
}

// get the starting position of the text
func (g *Game) GetTextPosition() (x, y int) {
	return 128, 600 - 120
}
func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("GrimoireGunner2")
	state := scene.NewSceneData()
	// TODO: read config off permanent storage
	state.BGMVolume = 128
	state.SfxVolume = 64
	ruleSet := map[stagehand.Scene[*scene.SceneData]][]stagehand.Directive[*scene.SceneData]{}
	manager := stagehand.NewSceneDirector[*scene.SceneData](scene.MainMenuInstance, state, ruleSet)
	if err := ebiten.RunGame(manager); err != nil {
		log.Fatal(err)
	}
}
