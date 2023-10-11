package main

import (
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep"
	"time"
	"fmt"
)

func main() {
	sr := beep.SampleRate(44100)
	speaker.Init(sr, sr.N(time.Second/10))

	for {
		fmt.Print("Press [ENTER] to quit. ")
		fmt.Scanln()
	}
}
