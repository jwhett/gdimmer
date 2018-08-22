package main

import (
    gd "github.com/jwhett/gdimmer"
    //"fmt"
)

func main() {
    d := gd.New("gmux_backlight")

	half := d.GetMax()/2

    d.SetBrightness(half)
}
