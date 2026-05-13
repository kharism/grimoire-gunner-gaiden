package assets

import (
	_ "embed"
	"fmt"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed shader/cooldown.kage
var cooldownShader []byte

var CooldownShader *ebiten.Shader

func init() {
	if CooldownShader == nil {
		var err error
		CooldownShader, err = ebiten.NewShader(cooldownShader)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(-1)
		}
	}
}
