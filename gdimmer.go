package gdimmer

import (
    "io/ioutil"
    "strconv"
    "strings"
)

// Struct to represent the state of a screen dimmer.
type Dimmer struct {
    max     int
    current int
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

// Initialize a new dimmer with the system's current settings.
func New() (*Dimmer) {

    // Take the []bytes from the "current brightness" file
    // and turn it into an integer.
    c, err := ioutil.ReadFile("/sys/class/backlight/gmux_backlight/brightness")
    check(err)
    cur := strings.TrimSpace(string(c))
    current, _ := strconv.Atoi(cur)

    // Take the []bytes from the "max brightness" file
    // and turn it into an integer.
    m, err := ioutil.ReadFile("/sys/class/backlight/gmux_backlight/max_brightness")
    check(err)
    mx := strings.TrimSpace(string(m))
    max, _ := strconv.Atoi(mx)

    return &Dimmer{max: max, current: current}
}

// Return the maximum brightness.
func (d *Dimmer) Max() int {
    return d.max
}

// Return the current brightness.
func (d *Dimmer) Current() int {
    return d.current
}

