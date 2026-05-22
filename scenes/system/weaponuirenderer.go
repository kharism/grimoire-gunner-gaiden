package system

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/kharism/GrimoireGunner2/scenes/assets"
	"github.com/yohamta/donburi/ecs"
)

type RenderableWeaponUi interface {
	// get the icon, must be 40x40 pixel
	GetIcon() *ebiten.Image
	// get the current progress must be between 0-1.0
	GetCooldownProgress() float64
	GetDamage() int
}

func RenderWeapon(e *ecs.ECS, screen *ebiten.Image) {
	blackRect := ebiten.NewImage(screen.Bounds().Dx(), 60)
	blackRect.Fill(color.RGBA{R: 125, G: 80, B: 80, A: 255})
	geom := ebiten.GeoM{}
	Ystart := 320.0
	geom.Translate(0, Ystart)
	screen.DrawImage(blackRect, &ebiten.DrawImageOptions{
		GeoM: geom,
	})
	offset := 10.0
	for curWeaponIdx, weapon := range WeaponSlot {
		if curWeaponIdx == SelectedSlot {
			highlightRect := ebiten.NewImage(110, 60)
			highlightRect.Fill(color.RGBA{R: 20, G: 180, B: 80, A: 255})
			geom2 := ebiten.GeoM{}
			geom2.Translate(offset, Ystart)
			screen.DrawImage(highlightRect, &ebiten.DrawImageOptions{
				GeoM: geom2,
			})
		}
		geom := ebiten.GeoM{}
		geom.Translate(offset, Ystart)
		progress := weapon.GetCooldownProgress()
		icon := weapon.GetIcon()
		damage := weapon.GetDamage()
		opts := &ebiten.DrawRectShaderOptions{
			GeoM: geom,
		}
		iconBound := icon.Bounds()
		opts.Images[0] = icon
		opts.Uniforms = make(map[string]interface{})
		opts.Uniforms["Iter"] = float32(progress)
		screen.DrawRectShader(iconBound.Dx(), iconBound.Dy(), assets.CooldownShader, opts)
		geom.Reset()
		geom.Scale(1.4, 1.4)
		geom.Translate(offset+float64(iconBound.Dx()), Ystart)
		textDrawOpt := text.DrawOptions{
			LayoutOptions: text.LayoutOptions{
				PrimaryAlign: text.AlignStart,
			},
			DrawImageOptions: ebiten.DrawImageOptions{
				GeoM: geom,
			},
		}
		text.Draw(screen, fmt.Sprintf("%.d", damage), assets.PixelOperatorFace, &textDrawOpt)
		offset += 100
	}

}
