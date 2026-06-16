package assets

import (
	"bytes"
	"image"
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

//go:embed "img/sprites/tank1.png"
var tank1 []byte

//go:embed "img/sprites/tank2.png"
var tank2 []byte

//go:embed "img/sprites/tank_fire_anim.png"
var tank_anim []byte

//go:embed "img/sprites/cyberhound1.png"
var cyberhound1 []byte

//go:embed "img/sprites/cyberhound2.png"
var cyberhound2 []byte

//go:embed "img/sprites/cyberhound3.png"
var cyberhound3 []byte

var SvenSprite1 *ebiten.Image
var SvenSprite2 *ebiten.Image
var SvenSprite3 *ebiten.Image
var CubeSprite *ebiten.Image
var Bullet *ebiten.Image
var Bomb *ebiten.Image
var TankSprite1 *ebiten.Image
var TankSprite2 *ebiten.Image //warmup
var TankSprite3 *ebiten.Image
var TankAnimFrames []*ebiten.Image

var CyberHound1 *ebiten.Image
var CyberHound2 *ebiten.Image //warmup
var CyberHound3 *ebiten.Image

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

		imgReader = bytes.NewReader(tank1)
		TankSprite1, _, _ = ebitenutil.NewImageFromReader(imgReader)
		imgReader = bytes.NewReader(tank2)
		TankSprite2, _, _ = ebitenutil.NewImageFromReader(imgReader)
		imgReader = bytes.NewReader(tank_anim)
		TankSprite3, _, _ = ebitenutil.NewImageFromReader(imgReader)
		TankAnimFrames = []*ebiten.Image{}
		for i := 0; i < TankSprite3.Bounds().Dx(); i += 80 {
			TankAnimFrames = append(TankAnimFrames, TankSprite3.SubImage(image.Rect(i, 0, i+80, 120)).(*ebiten.Image))
		}
		imgReader = bytes.NewReader(cyberhound1)
		CyberHound1, _, _ = ebitenutil.NewImageFromReader(imgReader)
		imgReader = bytes.NewReader(cyberhound2)
		CyberHound2, _, _ = ebitenutil.NewImageFromReader(imgReader)
		imgReader = bytes.NewReader(cyberhound3)
		CyberHound3, _, _ = ebitenutil.NewImageFromReader(imgReader)

		BombIcon = ebiten.NewImage(40, 40)
		BombIcon.Fill(color.Black)
		BombIcon.DrawImage(Bomb, &ebiten.DrawImageOptions{})
	}
}
