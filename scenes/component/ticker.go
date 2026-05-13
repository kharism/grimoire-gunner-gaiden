package component

import "github.com/yohamta/donburi"

type Tick interface {
	Tick()
}

type DummyTicker struct {
	H Tick
}

func (d *DummyTicker) Tick() {
	d.H.Tick()
}

var Ticker = donburi.NewComponentType[DummyTicker]()
