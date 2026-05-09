package component

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kharism/hanashi/core"
	"github.com/yohamta/donburi"
)

type ProjectileParam struct {
	Vx, Vy float64
	Pos    PositionComponentData
	//Col, Row       int
	Damage         int
	Sprite         *ebiten.Image
	OnHit          OnAtkHit
	FlipHorizontal bool
}

func NewProjectile(world donburi.World, param ProjectileParam) *donburi.Entity {
	entity := world.Create(
		Position,
		Velocity,
		//Health,
		Damage,
		OnHit,
		Sprite, ProjectileTag)
	entId := world.Entry(entity)
	Position.Set(entId, &param.Pos)
	Damage.SetValue(entId, DamageData{Damage: param.Damage})
	Velocity.Set(entId, &VelocityComponentData{X: param.Vx, Y: param.Vy})
	spriteData := &SpriteData{Image: param.Sprite}
	if param.FlipHorizontal {
		spriteData.Scale = &core.ScaleParam{Sx: -1, Sy: 1}
	}
	Sprite.Set(entId, &SpriteData{Image: param.Sprite})
	OnHit.Set(entId, &param.OnHit)
	return &entity
}
