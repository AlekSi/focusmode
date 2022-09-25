// Package focusmode returns information about macOS focus modes.
package focusmode

// Mode represents a macOS focus mode.
type Mode struct {
	// Mode ID, e.g. "com.apple.donotdisturb.mode.default".
	ID string
	// Mode name, e.g. "Do Not Disturb".
	Name string
}
