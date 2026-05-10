package scene

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/joelschutz/stagehand"
	"github.com/kharism/GrimoireGunner2/scenes/assets"
	"github.com/kharism/GrimoireGunner2/scenes/component"
	"github.com/kharism/GrimoireGunner2/scenes/system"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

type CombatScene struct {
	data   *SceneData
	sm     *stagehand.SceneDirector[*SceneData]
	world  donburi.World
	ecs    *ecs.ECS
	player donburi.Entity

	entitygrid  [4][8]int64
	musicPlayer *assets.AudioPlayer
	loopMusic   bool
}

func (c *CombatScene) Update() error {
	if c.loopMusic && !c.musicPlayer.AudioPlayer().IsPlaying() {
		c.musicPlayer.AudioPlayer().Rewind()
		c.musicPlayer.AudioPlayer().Play()
	}
	if c.musicPlayer != nil {
		c.musicPlayer.Update()
	}

	c.ecs.Update()
	return nil
}
func (c *CombatScene) Draw(screen *ebiten.Image) {
	screen.Clear()
	c.ecs.DrawLayer(LayerCharacter, screen)
	c.ecs.DrawLayer(LayerHP, screen)
	c.ecs.DrawLayer(LayerDebug, screen)
}
func (s *CombatScene) Load(state *SceneData, manager stagehand.SceneController[*SceneData]) {
	s.sm = manager.(*stagehand.SceneDirector[*SceneData]) // This type assertion is important
	s.world = donburi.NewWorld()
	s.entitygrid = [4][8]int64{}
	s.ecs = ecs.NewECS(s.world)
	s.data = state
	LoadGrid(s.world)
	s.player = LoadPlayer(s.world, state)
	LoadBlock(s.world, state, 2, 6)
	LoadBlock(s.world, state, 2, 3)
	s.ecs.AddSystem(system.PlayerMovementHandler)
	s.ecs.AddSystem(system.NonPlayerMovementHandler)
	s.ecs.AddSystem(system.PlayerAttackHandler)
	s.ecs.AddSystem(system.DamageSystemHandler)
	s.ecs.AddSystem(system.PositionCheckerSystem)
	s.ecs.AddRenderer(LayerCharacter, system.UnifiedRenderer)
	s.ecs.AddRenderer(LayerDebug, system.DrawDebug)
	s.ecs.AddRenderer(LayerHP, system.DrawHP)

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
