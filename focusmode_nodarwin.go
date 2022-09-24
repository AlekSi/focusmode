//go:build !darwin
// +build !darwin

package focusmode

// Current returns the current macOS focus mode.
func Current() (*Mode, error) {
	return nil, fmt.Errorf("not implemented")
}
