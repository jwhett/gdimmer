package gdimmer_test

import (
    "testing"
    "github.com/jwhett/gdimmer"
    "fmt"
    //"io/ioutil"
)

func TestInit(t *testing.T) {
    d := gdimmer.New()
    m := 1023
    c := 1023

    if d.Max() != m {
        fmt.Printf("Got: %d, Expected: %d", d.Max(), m)
        t.Error("Max not set properly...")
    }

    if d.Current() != c {
        t.Error("Current not set properly...")
    }
}
