package assets

import (
	"bytes"
	_ "embed"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

//go:embed "img/vfx/slash2.png"
var sword_slash []byte
var SwordSlash *ebiten.Image
var SwordSlashFrames []*ebiten.Image

func init() {
	if SwordSlash == nil {
		imgReader := bytes.NewReader(sword_slash)
		SwordSlash, _, _ = ebitenutil.NewImageFromReader(imgReader)
		swordSlashSize := SwordSlash.Bounds()
		for i := 0; i < swordSlashSize.Dx(); i += 80 {
			SwordSlashFrames = append(SwordSlashFrames, SwordSlash.SubImage(image.Rect(i, 0, i+80, 120)).(*ebiten.Image))
		}

	}
}
