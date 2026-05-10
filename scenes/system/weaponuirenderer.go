package system

import "github.com/hajimehoshi/ebiten/v2"

type RenderableWeaponUi interface {
	// get the icon, must be 40x40 pixel
	GetIcon() *ebiten.Image
	// get the current progress must be between 0-1.0
	GetCooldownProgress() float64
	GetDamage() int
}
