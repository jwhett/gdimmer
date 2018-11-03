package gdimmer_test

import (
	"github.com/jwhett/gdimmer"
	"io/ioutil"
	"strconv"
	"strings"
	"testing"
)

func TestInit(t *testing.T) {
	d := gdimmer.New("gmux_backlight")

	m, err := ioutil.ReadFile("/sys/class/backlight/gmux_backlight/max_brightness")
	if err != nil {
		t.Skip("Unable to read brightness files.")
	}
	mx := strings.TrimSpace(string(m))
	max, _ := strconv.Atoi(mx)

	c, err := ioutil.ReadFile("/sys/class/backlight/gmux_backlight/brightness")
	if err != nil {
		t.Skip("Unable to read brightness files.")
	}
	cur := strings.TrimSpace(string(c))
	current, _ := strconv.Atoi(cur)

	if d.GetMax() != max {
		t.Error("Max not set properly...")
	}

	if d.GetCurrent() != current {
		t.Error("Current not set properly...")
	}
}
