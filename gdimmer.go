package gdimmer

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

// Dimmer struct to represent the state of a screen dimmer.
type Dimmer struct {
	Provider
}

// SysfsProvider is a real provider in sysfs
type SysfsProvider struct {
	path string
	max  int
}

// force all SysfsProviders implement Provider interface
var _ Provider = SysfsProvider{}

// Provider is an interface lol
type Provider interface {
	GetMax() int
	GetCurrent() (int, error)
	SetCurrent(int) error
}

// NewSysfsProvider builds a SysfsProvider with
// proper initialization.
func NewSysfsProvider(path string) (SysfsProvider, error) {
	max, err := getIntFromFile(path + "/max_brightness")
	return SysfsProvider{max: max, path: path}, err
}

// GetMax returns the maximum brightness.
func (sp SysfsProvider) GetMax() int {
	return sp.max
}

// GetCurrent returns the current brightness.
func (sp SysfsProvider) GetCurrent() (int, error) {
	return getIntFromFile(sp.path + "/brightness")
}

// SetCurrent returns the current brightness.
func (sp SysfsProvider) SetCurrent(newlvl int) error {
	// Convert int to ASCII...
	bstr := strconv.Itoa(newlvl)
	// ...then write a byte slice to the brightness file.
	return ioutil.WriteFile(sp.path+"/brightness", []byte(bstr), 0444)
}

const (
	// ProviderDir sets the location for backlight providers
	ProviderDir = "/sys/class/backlight"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// GetProviders returns a list of providers discovered
// in /sys/class/backlight and an error, if needed.
func GetProviders() ([]string, error) {
	ProviderDir, err := os.Open(ProviderDir)

	if err != nil {
		fmt.Printf("Failed to open file: %s\n", err)
		return []string{}, err
	}
	defer ProviderDir.Close()

	providers, err := ProviderDir.Readdirnames(0)
	if err != nil {
		fmt.Printf("Failed to read dir names: %s\n", err)
		return []string{}, err
	}

	return providers, nil
}

// New initializes a new dimmer with the system's current settings.
func New(provider Provider) *Dimmer {
	return &Dimmer{Provider: provider}
}

// getIntFromFile takes the []bytes from the file path provided
// and turn it into an integer.
func getIntFromFile(fp string) (int, error) {
	i, err := ioutil.ReadFile(fp)
	if err != nil {
		return 0, err
	}
	istring := string(i[:len(i)-1])
	fullint, err := strconv.Atoi(istring)
	if err != nil {
		return 0, err
	}
	return fullint, err
}

// GetStep returns the step value.
func (d *Dimmer) GetStep() int {
	max := d.Provider.GetMax()
	return int(float64(max) * 0.1)
}

// SetBrightness sets the brightness to specified value.
func (d *Dimmer) SetBrightness(b int) error {
	var (
		max int
	)

	max = d.Provider.GetMax()

	if b > max {
		// Don't let someone assign something higher
		// than the max value.
		b = max
	}

	if b < 0 {
		// Don't let someone assign something lower
		// than zero.
		b = 0
	}

	return d.Provider.SetCurrent(b)
}

// StepDown decreases the brightness by step.
func (d *Dimmer) StepDown() error {
	current, err := d.GetCurrent()
	if err != nil {
		return err
	}
	return d.SetBrightness(current - d.GetStep())
}

// StepUp increases brightness by step.
func (d *Dimmer) StepUp() error {
	current, err := d.GetCurrent()
	if err != nil {
		return err
	}
	return d.SetBrightness(current + d.GetStep())
}
