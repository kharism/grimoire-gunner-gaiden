package system

import (
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/kharism/GrimoireGunner2/scenes/assets"
	"github.com/yohamta/donburi/ecs"
)

func Meter(ecs *ecs.ECS) {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) && MeterTickerInstance.curTick == max_bar {
		MeterTickerInstance.OpenCloser.Open()
	}
}

var max_bar = 120.0

func DrawMeter(ecs *ecs.ECS, screen *ebiten.Image) {
	geom := ebiten.GeoM{}
	geom.Translate(130, 20)
	drawImageOption := ebiten.DrawImageOptions{
		GeoM: geom,
	}
	screen.DrawImage(assets.Meter, &drawImageOption)
	rectColor := color.RGBA{R: 0xff, G: 0xD5, B: 0x41, A: 255}
	if MeterTickerInstance.curTick >= max_bar {
		rectColor = color.RGBA{R: 0x14, G: 0xa0, B: 0x2e, A: 255}
	}
	filledBar := ebiten.NewImage(1, 11)
	filledBar.Fill(rectColor)
	scale := MeterTickerInstance.curTick * (519 - 138) / max_bar
	geom.Reset()
	geom.Scale(scale, 1)
	geom.Translate(138, 23)
	drawImageOption.GeoM = geom
	screen.DrawImage(filledBar, &drawImageOption)
	if MeterTickerInstance.curTick >= max_bar {
		translate1 := ebiten.GeoM{}
		translate1.Translate(208, 23)
		op := &text.DrawOptions{
			LayoutOptions: text.LayoutOptions{
				PrimaryAlign: text.AlignStart,
				LineSpacing:  14,
			},
			DrawImageOptions: ebiten.DrawImageOptions{
				GeoM: translate1,
			},
		}
		text.Draw(screen, "Press Space for Powerup", assets.PixelOperatorFace, op)
	}
}

type OpenCloser interface {
	Open()
	Close()
}

var MeterTickerInstance = &MeterTicker{skipTick: 10}

type MeterTicker struct {
	OpenCloser OpenCloser
	curTick    float64
	LastTick   time.Time
	skipTick   int
	curTick2   int
}

func (t *MeterTicker) Tick() {
	t.curTick += 0.8
	if t.curTick > max_bar {
		t.curTick = max_bar
	}

}
