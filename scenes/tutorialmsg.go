package scene

import (
	"github.com/kharism/GrimoireGunner2/scenes/assets"
	"github.com/kharism/hanashi/core"
)

func SceneTutorial(layouter core.GetLayouter) *core.Scene {
	scene := core.NewScene()
	scene.SetLayouter(layouter)

	scene.Characters = []*core.Character{
		core.NewCharacterImage("Sven", assets.Sven),
	}
	scene.FontFace = assets.PixelOperatorFace
	portraitMoveParam := core.MoveParam{Sx: 10, Sy: 223, Tx: 10, Ty: 223}
	portraitScaleParam := &core.ScaleParam{Sx: 2, Sy: 2}
	scene.Events = []core.Event{
		&core.ComplexEvent{Events: []core.Event{
			core.NewCharacterAddEvent("Sven", portraitMoveParam, portraitScaleParam),
			core.NewDialogueEvent("Sven", "(Quick tutorial)", scene.FontFace),
		}},
		core.NewDialogueEvent("Sven", "(Move with arrow keys)", scene.FontFace),
		core.NewDialogueEvent("Sven", "(E for default attack)", scene.FontFace),
		core.NewDialogueEvent("Sven", "(R for sub-weapon)", scene.FontFace),
		core.NewDialogueEvent("Sven", "(1,2,3 for switch sub-weapon)", scene.FontFace),
	}
	scene.TxtBg = assets.DialogBox
	// pp, err := core.NewDefaultAudioInterfacer()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	os.Exit(-1)
	// }
	// scene.AudioInterface = pp
	scene.Events[0].Execute(scene)
	return scene
}
