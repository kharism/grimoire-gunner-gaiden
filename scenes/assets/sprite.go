package assets

import (
	"bytes"
	"image/color"

	_ "embed"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

//go:embed "img/sprites/sven_1.png"
var sven_sprite1 []byte

//go:embed "img/sprites/sven_2.png"
var sven_sprite2 []byte

//go:embed "img/sprites/sven_3.png"
var sven_sprite3 []byte

//go:embed "img/sprites/cube.png"
var cube_sprite1 []byte

//go:embed "img/sprites/bomb.png"
var bomb_sprite []byte

//go:embed "img/sprites/bullet.png"
var bullet []byte

var SvenSprite1 *ebiten.Image
var SvenSprite2 *ebiten.Image
var SvenSprite3 *ebiten.Image
var CubeSprite *ebiten.Image
var Bullet *ebiten.Image
var Bomb *ebiten.Image

func init() {
	if SvenSprite1 == nil {
		imgReader := bytes.NewReader(sven_sprite1)
		SvenSprite1, _, _ = ebitenutil.NewImageFromReader(imgReader)

		imgReader = bytes.NewReader(sven_sprite2)
		SvenSprite2, _, _ = ebitenutil.NewImageFromReader(imgReader)

		imgReader = bytes.NewReader(sven_sprite3)
		SvenSprite3, _, _ = ebitenutil.NewImageFromReader(imgReader)

		imgReader = bytes.NewReader(cube_sprite1)
		CubeSprite, _, _ = ebitenutil.NewImageFromReader(imgReader)

		imgReader = bytes.NewReader(bullet)
		Bullet, _, _ = ebitenutil.NewImageFromReader(imgReader)

		imgReader = bytes.NewReader(bomb_sprite)
		Bomb, _, _ = ebitenutil.NewImageFromReader(imgReader)

		BombIcon = ebiten.NewImage(40, 40)
		BombIcon.Fill(color.Black)
		BombIcon.DrawImage(Bomb, &ebiten.DrawImageOptions{})
	}
}
