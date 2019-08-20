package main

import (
	"flag"
	"fmt"
	"os"

	gd "github.com/jwhett/gdimmer"
)

func main() {

	var (
		up, down, max, force bool
		provider             string
		set                  int
	)

	providers, err := gd.GetProviders()
	if err != nil {
		fmt.Printf("Failed to get providers: %s", err)
		os.Exit(1)
	}

	if len(os.Getenv("BLPROVIDER")) > 0 {
		provider = os.Getenv("BLPROVIDER")
	} else if len(providers) > 0 {
		// just take the first one always
		provider = providers[0]
	}

	flag.StringVar(&provider, "provider", provider, "Chose backlight provider. Defaults to first provider found. Find your provider in /sys/class/backlight/. BLPROVIDER environment variable can also be set.")
	flag.StringVar(&provider, "p", provider, "Chose backlight provider. Defaults to first provider found. Find your provider in /sys/class/backlight/. BLPROVIDER environment variable can also be set.")

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

	sp, err := gd.NewSysfsProvider(gd.ProviderDir + "/" + provider)
	if err != nil {
		fmt.Printf("Cannot create new sysfs provider: %s", err)
		os.Exit(2)
	}
	d := gd.New(sp)

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
