package gdimmer_test

import (
    "github.com/jwhett/gdimmer"
    "io/ioutil"
    "strings"
    "testing"
)

func TestInit(t *testing.T) {
    d := gdimmer.New()
    m, _ := ioutil.ReadFile("/sys/class/backlight/gmux_backlight/max_brightness")
    c, _ := ioutil.ReadFile("/sys/class/backlight/gmux_backlight/brightness")

    if d.Max() != strings.TrimSpace(string(m)) {
        t.Error("Max not set properly...")
    }

    if d.Current() != strings.TrimSpace(string(c)) {
        t.Error("Current not set properly...")
    }
}

