package gdimmer

import (
    "io/ioutil"
    "strconv"
    "strings"
)

type Dimmer struct {
    max     string
    current string
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func New() (*Dimmer) {
    c, err := ioutil.ReadFile("/sys/class/backlight/gmux_backlight/brightness")
    check(err)

    m, err := ioutil.ReadFile("/sys/class/backlight/gmux_backlight/max_brightness")
    check(err)

    cur := strings.TrimSpace(string(c))
    check(err)

    mx := strings.TrimSpace(string(m))
    check(err)

    return &Dimmer{max: mx, current: cur}
}

func (d *Dimmer) Max() string {
    return d.max
}

func (d *Dimmer) Current() string {
    return d.current
}

func (d *Dimmer) Maxasi() int {
    imax, _ := strconv.Atoi(d.max)
    return imax
}

func (d *Dimmer) Currentasi() int {
    icur, _ := strconv.Atoi(d.current)
    return icur
}
