package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	gd "github.com/jwhett/gdimmer"
)

func getProviders() ([]string, error) {
	providerDir, err := os.Open("/sys/class/backlight")

	if err != nil {
		fmt.Printf("Failed to open file: %s\n", err)
		return []string{}, err
	}
	defer providerDir.Close()

	providers, err := providerDir.Readdirnames(0)
	if err != nil {
		fmt.Printf("Failed to read dir names: %s\n", err)
		return []string{}, err
	}

	return providers, nil
}

func main() {

	var (
		up, down, max, force bool
		provider             string
		set                  int
	)

	providers, err := getProviders()
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

	d := gd.New(filepath.Base(provider))

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
