package assets

import (
	"bytes"
	"log"

	_ "embed"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

//go:embed img/menu_btn_bg.png
var menuBg []byte

//go:embed "img/option_bg.png"
var optionBg []byte

//go:embed img/Arrow.png
var arrowBg []byte

//go:embed "img/dialog_box.png"
var dialogbox []byte

//go:embed "img/portrait/sven.png"
var sven []byte

//go:embed "img/portrait/rose.png"
var rose []byte

//go:embed "img/portrait/sudri.png"
var sudri []byte

//go:embed "img/portrait/clara1.png"
var clara1 []byte

//go:embed "img/portrait/wang1.png"
var wang1 []byte

//go:embed "img/bg/BG1.png"
var bg1 []byte

//go:embed "img/grid_blue.png"
var grid_blue []byte

//go:embed "img/grid_red.png"
var grid_red []byte

//go:embed "img/grid_dmg.png"
var grid_dmg []byte

//go:embed "img/danger.png"
var danger []byte

//go:embed "font/Unispace Bd.otf"
var UnispaceBdTTF []byte

//go:embed "font/pixelOperator8-Bold.ttf"
var PixelOperator8Bd []byte

//go:embed "bgm/opening.ogg"
var Menumusic []byte

//go:embed "sfx/menu.ogg"
var MenuMove []byte

//go:embed "bgm/DnB1.ogg"
var BGM1 []byte

var MenuButtonBg *ebiten.Image

var OptionBg *ebiten.Image

var ArrowBg *ebiten.Image

var DialogBox *ebiten.Image

var Sven *ebiten.Image
var Rose *ebiten.Image
var Sudri *ebiten.Image
var Clara1 *ebiten.Image
var Wang1 *ebiten.Image

var GridBlue *ebiten.Image
var GridRed *ebiten.Image
var GridDmg *ebiten.Image
var GridDanger *ebiten.Image

var BG1 *ebiten.Image

var UnispaceFont *text.GoTextFaceSource

var UnispaceFace *text.GoTextFace

var PixelOperatorFont *text.GoTextFaceSource
var PixelOperatorFace *text.GoTextFace

const TileLength = 80 //x-axis
const TileWidth = 40  //y-axis

func init() {
	if MenuButtonBg == nil {
		imgReader := bytes.NewReader(menuBg)
		MenuButtonBg, _, _ = ebitenutil.NewImageFromReader(imgReader)
	}
	if OptionBg == nil {
		imgReader := bytes.NewReader(optionBg)
		OptionBg, _, _ = ebitenutil.NewImageFromReader(imgReader)
	}
	if ArrowBg == nil {
		imgReader := bytes.NewReader(arrowBg)
		ArrowBg, _, _ = ebitenutil.NewImageFromReader(imgReader)
	}
	if DialogBox == nil {
		imgReader := bytes.NewReader(dialogbox)
		DialogBox, _, _ = ebitenutil.NewImageFromReader(imgReader)
	}
	if Sven == nil {
		imgReader := bytes.NewReader(sven)
		Sven, _, _ = ebitenutil.NewImageFromReader(imgReader)

		imgReader = bytes.NewReader(rose)
		Rose, _, _ = ebitenutil.NewImageFromReader(imgReader)

		imgReader = bytes.NewReader(sudri)
		Sudri, _, _ = ebitenutil.NewImageFromReader(imgReader)

		imgReader = bytes.NewReader(clara1)
		Clara1, _, _ = ebitenutil.NewImageFromReader(imgReader)

		imgReader = bytes.NewReader(wang1)
		Wang1, _, _ = ebitenutil.NewImageFromReader(imgReader)

		imgReader = bytes.NewReader(bg1)
		BG1, _, _ = ebitenutil.NewImageFromReader(imgReader)

		imgReader = bytes.NewReader(grid_blue)
		GridBlue, _, _ = ebitenutil.NewImageFromReader(imgReader)

		imgReader = bytes.NewReader(grid_red)
		GridRed, _, _ = ebitenutil.NewImageFromReader(imgReader)

		imgReader = bytes.NewReader(grid_dmg)
		GridDmg, _, _ = ebitenutil.NewImageFromReader(imgReader)

		imgReader = bytes.NewReader(danger)
		GridDanger, _, _ = ebitenutil.NewImageFromReader(imgReader)
	}
	s3, err := text.NewGoTextFaceSource(bytes.NewReader(UnispaceBdTTF))
	if err != nil {
		log.Fatal(err)
	}
	UnispaceFont = s3
	UnispaceFace = &text.GoTextFace{
		Source: UnispaceFont,
		Size:   15,
	}
	s4, err := text.NewGoTextFaceSource(bytes.NewReader(PixelOperator8Bd))
	if err != nil {
		log.Fatal(err)
	}
	PixelOperatorFont = s4
	PixelOperatorFace = &text.GoTextFace{
		Source: PixelOperatorFont,
		Size:   12,
	}
}
