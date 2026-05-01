package scene

import (
	"fmt"
	"os"

	"github.com/kharism/GrimoireGunner2/scenes/assets"
	"github.com/kharism/hanashi/core"
)

func Scene1(layouter core.GetLayouter) *core.Scene {
	scene := core.NewScene()
	scene.SetLayouter(layouter)

	scene.Characters = []*core.Character{
		core.NewCharacterImage("Sven", assets.Sven),
		core.NewCharacterImage("Rose", assets.Rose),
		core.NewCharacterImage("Sudri", assets.Sudri),
		core.NewCharacterImage("Clara", assets.Clara1),
		core.NewCharacterImage("Wang", assets.Wang1),
	}

	scene.FontFace = assets.PixelOperatorFace

	fmt.Println(scene.FontFace)
	portraitMoveParam := core.MoveParam{Sx: 10, Sy: 223, Tx: 10, Ty: 223}
	portraitScaleParam := &core.ScaleParam{Sx: 2, Sy: 2}
	//sceneWidth, sceneHeight := layouter.GetLayout()

	scene.Events = []core.Event{
		&core.ComplexEvent{Events: []core.Event{
			&core.PlayBgmEvent{Audio: &assets.BGM1, Type: core.TypeOgg},
			core.NewBgChangeEvent(assets.BG1, core.MoveParam{Sx: 0, Sy: 0, Tx: 0, Ty: 0, Speed: 0}, nil),
		}},
		&core.ComplexEvent{Events: []core.Event{
			//ore.NewCharacterAddEvent("Sven", portraitMoveParam, portraitScaleParam),
			core.NewDialogueEvent("", "the Mesh diver's guild lobby was not\nas full at the evening as it was in the\nmorning. The divers left were finishing\ntheir paperwork after clearing their missions.", scene.FontFace),
		}},
		&core.ComplexEvent{Events: []core.Event{
			//ore.NewCharacterAddEvent("Sven", portraitMoveParam, portraitScaleParam),
			core.NewDialogueEvent("", "Sven entered the lobby while bringing a\ndog with a horn. The dog seemed energetic\nand attempted to drag him along.\nHe tried his best to keep it properly leashed\n", scene.FontFace),
		}},
		&core.ComplexEvent{Events: []core.Event{
			core.NewCharacterAddEvent("Sven", portraitMoveParam, portraitScaleParam),
			core.NewDialogueEvent("Sven", "Hold on, boy. I know you missed\nyour owner, but try not to make any trouble for the\nreceptionists and the janitors, okay?.", scene.FontFace),
		}},
	}

	scene.TxtBg = assets.DialogBox
	pp, err := core.NewDefaultAudioInterfacer()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	scene.AudioInterface = pp

	return scene
}
