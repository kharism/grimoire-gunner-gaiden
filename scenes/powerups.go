package scene

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/kharism/GrimoireGunner2/scenes/assets"
	"github.com/kharism/GrimoireGunner2/scenes/system"
)

type powerupmenu interface {
	GetMenus() []string
	GetCurrentPick() int
	SetCurrentPick(int)
	SelectMenu(int, *powerupState) powerupmenu
}
type powerupState struct {
	menuPosX    int
	hide        bool
	peek        bool
	menuStack   []powerupmenu
	currentMenu powerupmenu
	topMenu     powerupmenu
	combatScene *CombatScene
}

func (s *powerupState) Update() {
	powerupmovespeed := 10
	if s.menuPosX < 0 && !s.hide {
		s.menuPosX += powerupmovespeed
	}
	if s.hide && s.menuPosX > -320 {
		s.menuPosX -= powerupmovespeed
	}
	if s.menuPosX == -320 {
		//trigger state change
		if !s.peek {
			s.combatScene.currentCombatSubState = s.combatScene.defaultSubState
		}

	}
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		s.Close()
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
		// if the menu is shown
		if s.menuPosX == 0 {
			s.peek = true
			s.hide = true
		} else if s.menuPosX == -320 {
			s.peek = true
			s.hide = false
			s.hide = false
		}

	}
	if inpututil.IsKeyJustPressed(ebiten.KeyE) {
		res := s.currentMenu.SelectMenu(s.currentMenu.GetCurrentPick(), s)
		if res != nil {

		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		if s.currentMenu != s.topMenu {
			s.currentMenu = s.topMenu
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		menus := s.currentMenu.GetMenus()
		j := (s.currentMenu.GetCurrentPick() + 1) % len(menus)
		s.currentMenu.SetCurrentPick(j)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		menus := s.currentMenu.GetMenus()
		j := (s.currentMenu.GetCurrentPick() - 1) % len(menus)
		s.currentMenu.SetCurrentPick(j)
	}
}
func (s *powerupState) Close() {
	s.hide = true
	s.peek = false
}
func (s *powerupState) Open() {
	s.combatScene.currentCombatSubState = s
	s.currentMenu = &poweruptopmenu{}
	// reset stuff
	system.SetDamage(system.DefaultDamage)
	system.SetDelay(system.DefaultDelay)
	s.hide = false
	s.peek = true
}
func (s *powerupState) Draw(screen *ebiten.Image) {
	translate1 := ebiten.GeoM{}
	translate1.Translate(float64(s.menuPosX), 0)
	screen.DrawImage(assets.PowerupMenu, &ebiten.DrawImageOptions{
		GeoM: translate1,
	})
	translate1.Reset()
	translate1.Scale(1.4, 1.4)
	translate1.Translate(float64(s.menuPosX+40), 44)
	op := &text.DrawOptions{
		LayoutOptions: text.LayoutOptions{
			PrimaryAlign: text.AlignStart,
			LineSpacing:  14,
		},
		DrawImageOptions: ebiten.DrawImageOptions{
			GeoM: translate1,
		},
	}
	text.Draw(screen, "Upgrade", assets.PixelOperatorFace, op)
	translate1.Reset()
	translate1.Scale(1.4, 1.4)
	translate1.Translate(float64(s.menuPosX+40), 278)
	op.DrawImageOptions.GeoM = translate1
	text.Draw(screen, "E to Select\nSpace to Cancel\nQ to peek\nR to return", assets.PixelOperatorFace, op)
	menus := s.currentMenu.GetMenus()
	startPosY := 88
	for i, j := range menus {
		translate1.Reset()
		translate1.Scale(1.4, 1.4)
		translate1.Translate(float64(s.menuPosX+35), float64(startPosY+i*28))
		if i == s.currentMenu.GetCurrentPick() {
			// draw triangle
			translate2 := ebiten.GeoM{}
			translate2.Scale(0.2, 0.2)
			translate2.Translate(float64(s.menuPosX+15), float64(startPosY+i*28))
			screen.DrawImage(assets.ArrowBg, &ebiten.DrawImageOptions{
				GeoM: translate2,
			})
		}
		text.Draw(screen, j, assets.PixelOperatorFace, &text.DrawOptions{
			LayoutOptions: text.LayoutOptions{
				PrimaryAlign: text.AlignStart,
			},
			DrawImageOptions: ebiten.DrawImageOptions{
				GeoM: translate1,
			},
		})
	}
}

type poweruptopmenu struct {
	currentPick int
	menus       []int
}

func (s *poweruptopmenu) GetMenus() []string {
	return []string{
		"GrimoireGun",
		"Suit",
		//"GigaAtk",
	}
}
func (s *poweruptopmenu) GetCurrentPick() int {
	return s.currentPick
}
func (s *poweruptopmenu) SetCurrentPick(i int) {
	s.currentPick = i
}
func (s *poweruptopmenu) SelectMenu(i int, substate *powerupState) powerupmenu {
	switch i {
	case 0:
		substate.menuStack = append(substate.menuStack, s)
		substate.currentMenu = &grimoiregunMenu{}
	case 1:

	}
	return nil
}

type grimoiregunMenu struct {
	selectedMenu int
}

func (g *grimoiregunMenu) GetMenus() []string {
	return []string{
		"DmgUp",
		"RapidUp",
	}
}
func (g *grimoiregunMenu) GetCurrentPick() int {
	return g.selectedMenu
}
func (g *grimoiregunMenu) SetCurrentPick(i int) {
	g.selectedMenu = i
}
func (g *grimoiregunMenu) SelectMenu(i int, s *powerupState) powerupmenu {
	switch i {
	case 0:
		system.SetDamage(8)
		s.Close()
		return nil
	case 1:
		pp, _ := time.ParseDuration("100ms")
		system.SetDelay(pp)
		s.Close()
		return nil
	}

	return nil
}
