package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/joelschutz/stagehand"
	scene "github.com/kharism/GrimoireGunner2/scenes"
	"github.com/kharism/GrimoireGunner2/scenes/system"
	"github.com/kharism/GrimoireGunner2/scenes/system/weapons"
	"github.com/kharism/hanashi/core"
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
	return 10 + 124, 360 - 150 + 10
}

// get the starting position of the text
func (g *Game) GetTextPosition() (x, y int) {
	return 10 + 124, 360 - 120 + 10
}
func (g *Game) GetTextBGPosition() (x, y int) {
	return 0, 360 - 150
}
func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("GrimoireGunnerGaiden")
	state := scene.NewSceneData()
	// TODO: read config off permanent storage
	state.BGMVolume = 0
	state.SfxVolume = 64

	state.Weapons = []system.RenderableCaster{
		weapons.NewGrenade(),
		weapons.NewWdSword(),
	}

	openingScene := scene.NewHanashiScene(scene.Scene1(&Game{}))
	openingScene.EscapeTrigger = scene.TriggerToMain

	combatScene := &scene.CombatScene{}

	ruleSet := map[stagehand.Scene[*scene.SceneData]][]stagehand.Directive[*scene.SceneData]{
		scene.MainMenuInstance: {
			stagehand.Directive[*scene.SceneData]{Dest: scene.OptionSceneInstance, Trigger: scene.TriggerToOption},
		},
		scene.OptionSceneInstance: {
			stagehand.Directive[*scene.SceneData]{Dest: scene.MainMenuInstance, Trigger: scene.TriggerToMain},
		},
		openingScene: {
			stagehand.Directive[*scene.SceneData]{Dest: scene.MainMenuInstance, Trigger: scene.TriggerToMain},
		},
	}
	core.DetectKeyboardNext = func() bool {
		return inpututil.IsKeyJustReleased(ebiten.KeyQ)
	}
	manager := stagehand.NewSceneDirector[*scene.SceneData](combatScene, state, ruleSet)
	if err := ebiten.RunGame(manager); err != nil {
		log.Fatal(err)
	}
}
