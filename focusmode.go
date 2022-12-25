// Package focusmode returns information about macOS focus modes.
package focusmode

// ModeID represents a focus mode ID.
type ModeID string

// List of pre-defined focus mode IDs.
const (
	ModeIDDoNotDisturb ModeID = "com.apple.donotdisturb.mode.default"
	ModeIDPersonal     ModeID = "com.apple.focus.personal-time"
	ModeIDWork         ModeID = "com.apple.focus.work"
	ModeIDSleep        ModeID = "com.apple.sleep.sleep-mode"
	ModeIDWorkout      ModeID = "com.apple.donotdisturb.mode.workout"
	ModeIDMindfulness  ModeID = "com.apple.focus.mindfulness"
	ModeIDReading      ModeID = "com.apple.focus.reading"
	ModeIDDriving      ModeID = "com.apple.donotdisturb.mode.driving"
	ModeIDGaming       ModeID = "com.apple.focus.gaming"
)

// Mode represents a macOS focus mode.
type Mode struct {
	// Focus mode ID, e.g. "com.apple.donotdisturb.mode.default" (ModeIDDoNotDisturb).
	ID ModeID
	// Focus mode name, e.g. "Do Not Disturb".
	Name string
}
