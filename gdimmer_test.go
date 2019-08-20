package gdimmer_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/jwhett/gdimmer"
)

type FakeProvider struct {
	max     int
	current int
}

func (fp *FakeProvider) GetMax() int {
	return fp.max
}

func (fp *FakeProvider) GetCurrent() (int, error) {
	return fp.current, nil
}

func (fp *FakeProvider) SetCurrent(newlvl int) error {
	fp.current = newlvl
	return nil
}

var _ gdimmer.Provider = &FakeProvider{}

func TestDimmer(t *testing.T) {
	testMax := 100
	testCurrent := 30
	stepby := 10
	fp := &FakeProvider{max: testMax, current: testCurrent}
	d := gdimmer.New(fp)

	if stepby != d.GetStep() {
		t.Error("Step set incorrectly...")
	}

	_ = d.StepDown()
	if fp.current != 20 {
		t.Error("Step down failed...")
	}

	_ = d.StepUp()
	if fp.current != 30 {
		t.Error("Step up failed...")
	}

	_ = d.SetBrightness(-10)
	if fp.current != 0 {
		t.Error("Failed to set minimum brightness when going below zero...")
	}

	_ = d.SetBrightness(110)
	if fp.current != fp.max {
		t.Error("Failed to set max brightness when going over max...")
	}
}

func TestSysfsProvider(t *testing.T) {
	tmpDir, err := ioutil.TempDir("", "")
	if err != nil {
		t.Fatalf("Failed to create temporary directory for SysfsProvider: %s", err)
	}
	defer os.RemoveAll(tmpDir)

	err = ioutil.WriteFile(tmpDir+"/max_brightness", []byte("100\n"), 0660)
	if err != nil {
		t.Fatalf("Failed to create max_brightness file for SysfsProvider: %s", err)
	}
	err = ioutil.WriteFile(tmpDir+"/brightness", []byte("30\n"), 0660)
	if err != nil {
		t.Fatalf("Failed to create brightness file for SysfsProvider: %s", err)
	}

	sp, err := gdimmer.NewSysfsProvider(tmpDir)
	if err != nil {
		t.Error("SysfsProvider unable to init...")
	}

	curr, err := sp.GetCurrent()
	if err != nil {
		t.Errorf("SysfsProvider error retrieving current value: %s", err)
	}

	if curr != 30 {
		t.Errorf("SysfsProvider failed to return correct current val... Returned: %v", curr)
	}

	if sp.GetMax() != 100 {
		t.Error("Failed to retrieve max value on sysfs provider.")
	}
}
