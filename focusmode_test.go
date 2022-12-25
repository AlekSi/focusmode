//go:build darwin
// +build darwin

package focusmode

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	t.Parallel()

	for dir, expected := range map[string]*Mode{
		"dnd": {
			ID:   ModeIDDoNotDisturb,
			Name: "Do Not Disturb",
		},
		"no": {},
		"workout": {
			ID:   ModeIDWorkout,
			Name: "Спорт",
		},
	} {
		dir, expected := dir, expected
		t.Run(dir, func(t *testing.T) {
			t.Parallel()

			assertions, err := os.ReadFile(filepath.Join("testdata", dir, "Assertions.json"))
			if err != nil {
				t.Fatal(err)
			}

			modeConfigurations, err := os.ReadFile(filepath.Join("testdata", dir, "ModeConfigurations.json"))
			if err != nil {
				t.Fatal(err)
			}

			actual, err := parse(assertions, modeConfigurations)
			if err != nil {
				t.Fatal(err)
			}

			if !reflect.DeepEqual(expected, actual) {
				t.Fatalf("expected = %+v, actual = %+v", expected, actual)
			}
		})
	}
}
