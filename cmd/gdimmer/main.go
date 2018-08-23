package main

import (
    "flag"
    gd "github.com/jwhett/gdimmer"
)

func main() {
	up   := flag.Bool("up", false, "Turn up the brightness.")
	down := flag.Bool("down", false, "Turn down the brightness.")
	max  := flag.Bool("max", false, "Maximum brightness.")
	flag.Parse()

    d := gd.New("gmux_backlight")

	if *up {
    	d.StepUp()
	}
	if *down {
    	d.StepDown()
	}
	if *max {
    	d.SetBrightness(d.GetMax())
	}
}
