package gdimmer_test

import (
    "testing"
    "github.com/jwhett/gdimmer"
)

func TestInit(t *testing.T) {
    d := gdimmer.New()

    if d.Max() != 1024 {
        t.Error("Max not set properly...")
    }

    if d.Current() != 512 {
        t.Error("Current not set properly...")
    }
}
