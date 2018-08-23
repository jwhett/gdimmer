package gdimmer

import (
    "io/ioutil"
    "strconv"
    "strings"
)

// Struct to represent the state of a screen dimmer.
type Dimmer struct {
    maxfile     string
    currentfile string
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

// Initialize a new dimmer with the system's current settings.
func New(provider string) (*Dimmer) {

    var (
        basedir string
        current string
        max     string
    )

	// Build path to important files.
    basedir = strings.Join([]string{"/sys/class/backlight/", provider}, "")
    current = strings.Join([]string{basedir, "/brightness"}, "")
    max     = strings.Join([]string{basedir, "/max_brightness"}, "")

    return &Dimmer{maxfile: max, currentfile: current}
}

func getIntFromFile(fp string) int {
    // Take the []bytes from the file path provided
    // and turn it into an integer.
    i, err := ioutil.ReadFile(fp)
    check(err)
    istring := strings.TrimSpace(string(i))
    fullint, err := strconv.Atoi(istring)
    check(err)
    return fullint
}

// Return the maximum brightness.
func (d *Dimmer) GetMax() int {
	return getIntFromFile(d.maxfile)
}

// Return the current brightness.
func (d *Dimmer) GetCurrent() int {
	return getIntFromFile(d.currentfile)
}

// Return the step value.
func (d *Dimmer) GetStep() int {
    var max float64 = float64(d.GetMax())
    var step float64 = 0.1
    var stepby int = int(max * step)
    return stepby
}

// Set brightness to specified value.
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

// Step brightness down.
func (d *Dimmer) StepDown() {
    d.SetBrightness(d.GetCurrent() - d.GetStep())
}

// Step brightness up.
func (d *Dimmer) StepUp() {
    d.SetBrightness(d.GetCurrent() + d.GetStep())
}
