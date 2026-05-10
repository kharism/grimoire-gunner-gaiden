package component

import (
	"fmt"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

// the grid plane is X and Z (Z is Z-buffer), Y is the height
// the closer to the camera the higher Z-buffer is
// assume the Z=0 at the top row
type PositionComponentData struct {
	X, Y, Z float64
}

func (p PositionComponentData) Order() int {
	return int(p.Z * 100)
}
func (p *PositionComponentData) String() string {
	return fmt.Sprintf("{X:%f,Y:%f,Z:%f}", p.X, p.Y, p.Z)
}

type PositionCheckerComponent func(ecs *ecs.ECS) bool

var Position = donburi.NewComponentType[PositionComponentData]()

var PositionChecker = donburi.NewComponentType[PositionCheckerComponent]()
