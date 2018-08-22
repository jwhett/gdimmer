package gdimmer

import (
    "strings"
    "strconv"
    "io/ioutil"
)

type Dimmer struct {
    max     int
    current int
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

    // current, err := strconv.Atoi(c)
    cur := strings.TrimSpace(string(c))
    check(err)

    // max, err := strconv.Atoi(m)
    mx := strings.TrimSpace(string(m))
    check(err)

    max, _ := strconv.Atoi(mx)
    current, _ := strconv.Atoi(cur)

    return &Dimmer{max: max, current: current}
}

func (d *Dimmer) Max() int {
    return d.max
}

func (d *Dimmer) Current() int {
    return d.current
}

