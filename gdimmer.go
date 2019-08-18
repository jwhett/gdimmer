package gdimmer

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

// Dimmer struct to represent the state of a screen dimmer.
type Dimmer struct {
	maxfile     string
	currentfile string
}

const (
	// ProviderDir sets the location for backlight providers
	ProviderDir = "/sys/class/backlight"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// GetProviders returns a list of providers discovered
// in /sys/class/backlight and an error, if needed.
func GetProviders() ([]string, error) {
	ProviderDir, err := os.Open(ProviderDir)

	if err != nil {
		fmt.Printf("Failed to open file: %s\n", err)
		return []string{}, err
	}
	defer ProviderDir.Close()

	providers, err := ProviderDir.Readdirnames(0)
	if err != nil {
		fmt.Printf("Failed to read dir names: %s\n", err)
		return []string{}, err
	}

	return providers, nil
}

// New initializes a new dimmer with the system's current settings.
func New(provider string) *Dimmer {

	var (
		basedir string
		current string
		max     string
	)

	// Build path to important files.
	basedir = ProviderDir + "/" + provider
	current = basedir + "/brightness"
	max = basedir + "/max_brightness"

	return &Dimmer{maxfile: max, currentfile: current}
}

func getIntFromFile(fp string) int {
	// Take the []bytes from the file path provided
	// and turn it into an integer.
	i, err := ioutil.ReadFile(fp)
	check(err)
	istring := string(i[:len(i)-1])
	fullint, err := strconv.Atoi(istring)
	check(err)
	return fullint
}

// GetMax returns the maximum brightness.
func (d *Dimmer) GetMax() int {
	return getIntFromFile(d.maxfile)
}

// GetCurrent returns the current brightness.
func (d *Dimmer) GetCurrent() int {
	return getIntFromFile(d.currentfile)
}

// GetStep returns the step value.
func (d *Dimmer) GetStep() int {
	var max = float64(d.GetMax())
	var step = 0.1
	var stepby = int(max * step)
	return stepby
}

// SetBrightness sets the brightness to specified value.
func (d *Dimmer) SetBrightness(b int) {
	var (
		max int
	)

	max = d.GetMax()

	if b > max {
		// Don't let someone assign something higher
		// than the max value.
		b = max
	}

	if b < 0 {
		// Don't let someone assign something lower
		// than zero.
		b = 0
	}

	// Convert int to ASCII...
	bstr := strconv.Itoa(b)
	// ...then write a byte slice to the brightness file.
	err := ioutil.WriteFile(d.currentfile, []byte(bstr), 0444)
	check(err)
}

// StepDown decreases the brightness by step.
func (d *Dimmer) StepDown() {
	d.SetBrightness(d.GetCurrent() - d.GetStep())
}

// StepUp increases brightness by step.
func (d *Dimmer) StepUp() {
	d.SetBrightness(d.GetCurrent() + d.GetStep())
}
