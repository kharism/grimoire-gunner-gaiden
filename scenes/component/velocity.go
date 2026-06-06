package component

import (
	"math"

	"github.com/yohamta/donburi"
)

type VelocityComponentData struct {
	X, Y, Z float64
}

func (v *VelocityComponentData) IsMoving() bool {
	return math.Abs(v.X)+math.Abs(v.Y)+math.Abs(v.Z) != 0
}

var Velocity = donburi.NewComponentType[VelocityComponentData]()
var SingleGridMovementTag = donburi.NewTag("SingleGridMove")

type AccellerationComponentData struct {
	DX, DY, DZ float64
}

var Acceleration = donburi.NewComponentType[AccellerationComponentData]()
