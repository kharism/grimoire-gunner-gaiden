package component

import "github.com/yohamta/donburi"

const STATE_STANDING = "STANDING"
const STATE_ATTACK = "ATTACK"

var State = donburi.NewComponentType[string]()
