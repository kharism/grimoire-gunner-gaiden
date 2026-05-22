package assets

import (
	"bytes"
	_ "embed"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

//go:embed "img/icon/sword_icon.png"
var sword_icon []byte

var BombIcon *ebiten.Image
var SwordIcon *ebiten.Image

func init() {
	if SwordIcon == nil {
		imgReader := bytes.NewReader(sword_icon)
		SwordIcon, _, _ = ebitenutil.NewImageFromReader(imgReader)
	}
}
