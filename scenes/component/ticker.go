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

type TransientTicker struct {
	World    donburi.World
	Entry    *donburi.Entry
	TimeTick int
	curTick  int
}

func (t *TransientTicker) Tick() {
	t.curTick += 1
	if t.curTick == t.TimeTick {
		t.World.Remove(t.Entry.Entity())
	}
}
