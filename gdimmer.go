package gdimmer

import (
    // "fmt"
    // "io"
)

type Dimmer struct {
    max     int
    current int
}

func New() (*Dimmer) {
    return &Dimmer{max: 1024, current: 512}
}

func (d *Dimmer) Max() int {
    return d.max
}

func (d *Dimmer) Current() int {
    return d.current
}

