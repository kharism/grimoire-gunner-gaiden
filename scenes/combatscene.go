package scene

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/joelschutz/stagehand"
	"github.com/kharism/GrimoireGunner2/scenes/assets"
	"github.com/kharism/GrimoireGunner2/scenes/component"
	"github.com/kharism/GrimoireGunner2/scenes/system"
	"github.com/kharism/GrimoireGunner2/scenes/system/enemies"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

type CombatScene struct {
	data   *SceneData
	sm     *stagehand.SceneDirector[*SceneData]
	world  donburi.World
	ecs    *ecs.ECS
	player donburi.Entity

	entitygrid            [4][8]int64
	musicPlayer           *assets.AudioPlayer
	loopMusic             bool
	currentCombatSubState combatSceneSubstate
	defaultSubState       combatSceneSubstate
	powerupSubState       combatSceneSubstate
}
type combatSceneSubstate interface {
	Update()
	Draw(screen *ebiten.Image)
}
type defaultCombatState struct {
	ecs *ecs.ECS
}

func (c *defaultCombatState) Update() {
	c.ecs.Update()
}
func (c *defaultCombatState) Draw(screen *ebiten.Image) {
	// by default already drawn on main Draw function, this function is just a stub
}

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
	translate1.Translate(float64(s.menuPosX+40), 280)
	op.DrawImageOptions.GeoM = translate1
	text.Draw(screen, "E to Select\nSpace to Cancel\nQ to peek", assets.PixelOperatorFace, op)
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

func (c *CombatScene) Update() error {
	if c.loopMusic && !c.musicPlayer.AudioPlayer().IsPlaying() {
		c.musicPlayer.AudioPlayer().Rewind()
		c.musicPlayer.AudioPlayer().Play()
	}
	if c.musicPlayer != nil {
		c.musicPlayer.Update()
	}
	c.currentCombatSubState.Update()
	//c.ecs.Update()
	return nil
}
func (c *CombatScene) Draw(screen *ebiten.Image) {
	screen.Clear()
	c.ecs.DrawLayer(LayerCharacter, screen)
	c.ecs.DrawLayer(LayerHP, screen)
	c.ecs.DrawLayer(LayerUI, screen)
	c.ecs.DrawLayer(LayerDebug, screen)
	c.currentCombatSubState.Draw(screen)
}
func (s *CombatScene) Load(state *SceneData, manager stagehand.SceneController[*SceneData]) {
	s.sm = manager.(*stagehand.SceneDirector[*SceneData]) // This type assertion is important
	s.world = donburi.NewWorld()
	s.entitygrid = [4][8]int64{}
	s.ecs = ecs.NewECS(s.world)
	s.data = state
	LoadGrid(s.world)
	s.player = LoadPlayer(s.world, state)
	s.defaultSubState = &defaultCombatState{ecs: s.ecs}
	s.powerupSubState = &powerupState{
		menuStack:   []powerupmenu{},
		currentMenu: &poweruptopmenu{},
		menuPosX:    -320,
		combatScene: s,
	}
	s.currentCombatSubState = s.powerupSubState
	//LoadBlock(s.world, state, 2, 6)
	//LoadBlock(s.world, state, 2, 7)

	// tankEntity := s.world.Create(
	// 	component.Position,
	// 	component.Sprite,
	// )
	// tankEntry := s.world.Entry(tankEntity)
	// component.Sprite.Set(tankEntry, &component.SpriteData{
	// 	Image: assets.TankSprite1,
	// })
	// gridStartX := 5
	// gridStartY := 2
	// component.Position.Set(tankEntry, &component.PositionComponentData{
	// 	X: startX + float64(gridStartX)*gridLength,
	// 	Z: startX + float64(gridStartY)*gridLength,
	// 	Y: startX + float64(gridStartY)*gridLength,
	// })

	LoadBlock(s.world, state, 2, 3)
	enemies.GridLength = gridLength
	enemies.GridWidth = gridWidth
	enemies.StartX = startX
	enemies.StartY = startY
	enemies.LoadTank(s.ecs, 2, 5)

	system.WeaponSlot = state.Weapons
	for _, w := range state.Weapons {
		ent := s.world.Create(component.Ticker)
		e := s.world.Entry(ent)
		component.Ticker.SetValue(e, component.DummyTicker{w})
	}
	system.LastTick = time.Now()

	s.ecs.AddSystem(system.PlayerMovementHandler)
	s.ecs.AddSystem(system.NonPlayerMovementHandler)
	s.ecs.AddSystem(system.PlayerAttackHandler)
	s.ecs.AddSystem(system.DamageSystemHandler)
	s.ecs.AddSystem(system.PositionCheckerSystem)
	s.ecs.AddSystem(system.Tick)
	s.ecs.AddRenderer(LayerCharacter, system.UnifiedRenderer)
	s.ecs.AddRenderer(LayerDebug, system.DrawDebug)
	s.ecs.AddRenderer(LayerHP, system.DrawHP)
	s.ecs.AddRenderer(LayerUI, system.RenderWeapon)

}
func LoadPlayer(world donburi.World, state *SceneData) donburi.Entity {
	playerEntity := world.Create(
		component.Health,
		component.Position,
		component.Sprite,
		component.PlayerTag,
		component.Velocity,
	)
	playerEntry := world.Entry(playerEntity)
	component.Health.Set(playerEntry, &component.HealthData{
		HP:    100,
		Name:  "Player",
		MaxHP: 100,
	})
	component.Sprite.Set(playerEntry, &component.SpriteData{
		Image: assets.SvenSprite1,
	})
	gridStartX := 1
	gridStartY := 1
	component.Position.Set(playerEntry, &component.PositionComponentData{
		X: startX + float64(gridStartX)*gridLength,
		Z: startY + float64(gridStartY)*gridWidth,
		Y: startY + float64(gridStartY)*gridWidth,
	})
	component.Velocity.Set(playerEntry, &component.VelocityComponentData{
		X: 0, Y: 0, Z: 0,
	})

	return playerEntity
}

func LoadBlock(world donburi.World, state *SceneData, row, col int) {
	playerEntity := world.Create(
		component.Health,
		component.Position,
		component.Sprite,
	)
	playerEntry := world.Entry(playerEntity)
	component.Health.Set(playerEntry, &component.HealthData{
		HP:    100,
		MaxHP: 100,
		Name:  "rock",
	})
	component.Sprite.Set(playerEntry, &component.SpriteData{
		Image: assets.CubeSprite,
	})
	gridStartX := col
	gridStartY := row
	component.Position.Set(playerEntry, &component.PositionComponentData{
		X: startX + float64(gridStartX)*gridLength,
		Z: startY + float64(gridStartY)*gridWidth,
		Y: startY + float64(gridStartY)*gridWidth,
	})
}

// startX and startY is the top left grid coordinate
var startX = 40.0
var startY = 180.0

var gridLength = 80.0
var gridWidth = 40.0

func LoadGrid(world donburi.World) {
	component.GridLength = int(gridLength)
	component.GridWidth = int(gridWidth)

	component.GridStartPointX = int(startX)
	component.GridStartPointY = int(startY)
	for i := 0; i < 4; i++ {
		for j := 0; j < 8; j++ {
			idx := world.Create(component.Position, component.Sprite, component.TileTag)
			entId := world.Entry(idx)
			// for the grid we treat Y on the grid as Y on the screen
			component.Position.Set(entId, &component.PositionComponentData{X: startX + float64(j*80), Y: startY + float64(i*40)})
			if j < 4 {
				component.Sprite.Set(entId, &component.SpriteData{Image: assets.GridBlue})
			} else {
				component.Sprite.Set(entId, &component.SpriteData{Image: assets.GridRed})
			}

		}
	}
}
func (s *CombatScene) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 360
}
func (s *CombatScene) Unload() *SceneData {

	s.data.MusicSeek = s.musicPlayer.AudioPlayer().Position()
	s.musicPlayer.AudioPlayer().Rewind()
	s.musicPlayer.AudioPlayer().Pause()
	return s.data
}
