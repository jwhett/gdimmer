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

// Return the maximum brightness.
func (d *Dimmer) GetMax() int {
    // Take the []bytes from the "max brightness" file
    // and turn it into an integer.
    m, err := ioutil.ReadFile(d.maxfile)
    check(err)
    mx := strings.TrimSpace(string(m))
    max, _ := strconv.Atoi(mx)
    return max
}

// Return the current brightness.
func (d *Dimmer) GetCurrent() int {
    // Take the []bytes from the "current brightness" file
    // and turn it into an integer.
    c, err := ioutil.ReadFile(d.currentfile)
    check(err)
    cur := strings.TrimSpace(string(c))
    current, _ := strconv.Atoi(cur)
    return current
}
