package save

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/dcxforge/bugcraft/internal/model"
)

func Load(path string) (model.Save, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return model.Save{},
			fmt.Errorf("Coudnot read save: %w", err)
	}

	var s model.Save

	if err := json.Unmarshal(data, &s); err != nil {
		return model.Save{},
			fmt.Errorf("Coudnot parse json save: %w", err)
	}

	return s, nil
}

func Write(path string, s model.Save) error {
	data, err := json.MarshalIndent(s, "", "   ")
	if err != nil {
		return fmt.Errorf("Couldnot encode save to json: %w", err)
	}

	// Create the directories with the following permisisons
	// Owner - rw- (6)
	// Group & Others - r-- (4)
	if err := os.WriteFile(path, data, 0o644); err != nil {
		return fmt.Errorf("Couldnot save: %w", err)
	}
	return nil
}
