package assets

import (
	"bytes"

	_ "embed"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

//go:embed "img/sprites/sven_1.png"
var sven_sprite1 []byte

//go:embed "img/sprites/cube.png"
var cube_sprite1 []byte

var SvenSprite1 *ebiten.Image
var CubeSprite *ebiten.Image

func init() {
	if SvenSprite1 == nil {
		imgReader := bytes.NewReader(sven_sprite1)
		SvenSprite1, _, _ = ebitenutil.NewImageFromReader(imgReader)

		imgReader = bytes.NewReader(cube_sprite1)
		CubeSprite, _, _ = ebitenutil.NewImageFromReader(imgReader)
	}
}
