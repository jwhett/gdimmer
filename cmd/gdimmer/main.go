package main

import (
    "flag"
    gd "github.com/jwhett/gdimmer"
)

func main() {

    var (
        up, down, max, force bool
        set int
    )

	flag.BoolVar(&up, "up", false, "Turn up the brightness.")
	flag.BoolVar(&down, "down", false, "Turn down the brightness.")
	flag.BoolVar(&max, "max", false, "Maximum brightness.")
	flag.BoolVar(&force, "f", false, "Force; to be used with 'set'.")
	flag.IntVar(&set, "set", 1, "Explicitly set brightness to VALUE.")
	flag.Parse()

    d := gd.New("gmux_backlight")

	if up {
    	d.StepUp()
	}
	if down {
    	d.StepDown()
	}
	if max {
    	d.SetBrightness(d.GetMax())
	}
	if force {
    	d.SetBrightness(set)
	}
}
