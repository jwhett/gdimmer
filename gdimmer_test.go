package gdimmer_test

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"testing"

	"github.com/jwhett/gdimmer"
)

func TestInit(t *testing.T) {
	providers, err := gdimmer.GetProviders()
	if err != nil {
		fmt.Printf("Error getting providers: %s\n", err)
	}
	firstProvider := providers[0]
	fullProvider := gdimmer.ProviderDir + "/" + firstProvider
	t.Logf("fullProvider: %s\n", fullProvider)

	d := gdimmer.New(firstProvider)

	m, err := ioutil.ReadFile(fullProvider + "/max_brightness")
	if err != nil {
		t.Skip("Unable to read brightness files.")
	}
	mx := strings.TrimSpace(string(m))
	max, _ := strconv.Atoi(mx)

	c, err := ioutil.ReadFile(fullProvider + "/brightness")
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
