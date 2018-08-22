package gdimmer_test

import (
    "github.com/jwhett/gdimmer"
    "io/ioutil"
    "strings"
    "strconv"
    "testing"
)

func TestInit(t *testing.T) {
    d := gdimmer.New("gmux_backlight")

    m, _ := ioutil.ReadFile("/sys/class/backlight/gmux_backlight/max_brightness")
    mx := strings.TrimSpace(string(m))
    max, _ := strconv.Atoi(mx)

    c, _ := ioutil.ReadFile("/sys/class/backlight/gmux_backlight/brightness")
    cur := strings.TrimSpace(string(c))
    current, _ := strconv.Atoi(cur)

    if d.GetMax() != max {
        t.Error("Max not set properly...")
    }

    if d.GetCurrent() != current {
        t.Error("Current not set properly...")
    }
}

