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
