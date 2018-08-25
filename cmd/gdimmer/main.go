package main

import (
	"flag"
	gd "github.com/jwhett/gdimmer"
	"os"
)

func main() {

	var (
		up, down, max, force bool
		provider             string
		set                  int
	)

	if len(os.Getenv("BLPROVIDER")) > 0 {
		provider = os.Getenv("BLPROVIDER")
	} else {
		flag.StringVar(&provider, "provider", "gmux_backlight", "Chose backlight provider. Defaults to GMUX (Macintosh). Find your provider in /sys/class/backlight/. BLPROVIDER environment variable can also be set.")
		flag.StringVar(&provider, "p", "gmux_backlight", "Chose backlight provider. Defaults to GMUX (Macintosh). Find your provider in /sys/class/backlight/. BLPROVIDER environment variable can also be set.")
	}

	flag.BoolVar(&up, "up", false, "Turn up the brightness.")
	flag.BoolVar(&up, "u", false, "Turn up the brightness.")
	flag.BoolVar(&down, "down", false, "Turn down the brightness.")
	flag.BoolVar(&down, "d", false, "Turn down the brightness.")
	flag.BoolVar(&max, "max", false, "Maximum brightness.")
	flag.BoolVar(&max, "m", false, "Maximum brightness.")
	flag.BoolVar(&force, "force", false, "Force; to be used with 'set'.")
	flag.BoolVar(&force, "f", false, "Force; to be used with 'set'.")
	flag.IntVar(&set, "set", 1, "Explicitly set brightness to VALUE.")
	flag.IntVar(&set, "s", 1, "Explicitly set brightness to VALUE.")
	flag.Parse()

	d := gd.New(provider)

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
