package focusmode

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type assertionsFile struct {
	Data []struct {
		StoreAssertionRecords []struct {
			AssertionDetails struct {
				AssertionDetailsModeIdentifier string `json:"assertionDetailsModeIdentifier"`
			} `json:"assertionDetails"`
		} `json:"storeAssertionRecords"`
	} `json:"data"`
}

type modeConfigurationsFile struct {
	Data []struct {
		ModeConfigurations map[string]struct {
			Mode struct {
				Name string `json:"name"`
			} `json:"mode"`
		} `json:"modeConfigurations"`
	} `json:"data"`
}

// Current returns the current macOS focus mode.
//
// If the focus mode is not enabled, an empty non-nil result is returned.
func Current() (*Mode, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	assertions, err := os.ReadFile(filepath.Join(home, "Library/DoNotDisturb/DB/Assertions.json"))
	if err != nil {
		return nil, err
	}

	modeConfigurations, err := os.ReadFile(filepath.Join(home, "Library/DoNotDisturb/DB/ModeConfigurations.json"))
	if err != nil {
		return nil, err
	}

	return parse(assertions, modeConfigurations)
}

func parse(assertions, modeConfigurations []byte) (*Mode, error) {
	var assertionsFile assertionsFile
	if err := json.Unmarshal(assertions, &assertionsFile); err != nil {
		return nil, err
	}

	var modeConfigurationsFile modeConfigurationsFile
	if err := json.Unmarshal(modeConfigurations, &modeConfigurationsFile); err != nil {
		return nil, err
	}

	if l := len(assertionsFile.Data); l != 1 {
		return nil, fmt.Errorf("assertions: expected 1 data, got %d", l)
	}

	storeInvalidationRecords := assertionsFile.Data[0].StoreAssertionRecords
	switch l := len(storeInvalidationRecords); l {
	case 0:
		return new(Mode), nil
	case 1:
		if l := len(modeConfigurationsFile.Data); l != 1 {
			return nil, fmt.Errorf("modeConfigurations: expected 1 data, got %d", l)
		}

		id := storeInvalidationRecords[0].AssertionDetails.AssertionDetailsModeIdentifier
		name := modeConfigurationsFile.Data[0].ModeConfigurations[id].Mode.Name

		return &Mode{
			ID:   id,
			Name: name,
		}, nil
	default:
		return nil, fmt.Errorf("assertions: expected 1 storeAssertionRecords, got %d", l)
	}
}
