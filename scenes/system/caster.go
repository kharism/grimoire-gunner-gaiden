package system

import "github.com/yohamta/donburi/ecs"

type Caster interface {
	Cast(e *ecs.ECS)
}
