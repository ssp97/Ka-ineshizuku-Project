package zero

import (
	ZeroBot "github.com/wdvxdr1123/ZeroBot"
)

var defaultZero = ZeroBot.New()


func RunDefault(op ZeroBot.Config) {
	ZeroBot.Run(op)
}

func Default() *ZeroBot.Engine {
	return defaultZero
}


