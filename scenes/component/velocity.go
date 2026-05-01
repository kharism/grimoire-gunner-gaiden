package component

import "github.com/yohamta/donburi"

type VelocityComponentData struct {
	X, Y, Z float64
}

var Velocity = donburi.NewComponentType[VelocityComponentData]()

type AccellerationComponentData struct {
	DX, DY, DZ float64
}

var Acceleration = donburi.NewComponentType[AccellerationComponentData]()
