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

//go:embed "font/Unispace Bd.otf"
var UnispaceBdTTF []byte

//go:embed "bgm/opening.ogg"
var Menumusic []byte

//go:embed "sfx/menu.ogg"
var MenuMove []byte

var MenuButtonBg *ebiten.Image

var OptionBg *ebiten.Image

var ArrowBg *ebiten.Image

var UnispaceFont *text.GoTextFaceSource

var UnispaceFace *text.GoTextFace

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
	s3, err := text.NewGoTextFaceSource(bytes.NewReader(UnispaceBdTTF))
	if err != nil {
		log.Fatal(err)
	}
	UnispaceFont = s3
	UnispaceFace = &text.GoTextFace{
		Source: UnispaceFont,
		Size:   15,
	}
}
