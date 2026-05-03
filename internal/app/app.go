package app

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/dcxforge/bugcraft/internal/model"
	"github.com/dcxforge/bugcraft/internal/save"
)

const DefaultDirName string = ".bugcraft"

func Dir() (string, error) {
	if override := os.Getenv("BUGCRAFT_HOME"); override != "" {
		return override, nil
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("Cannot find home directory: %w", err)
	}
	return filepath.Join(home, DefaultDirName), nil
}

func ConfigPath(dir string) string { return filepath.Join(dir, "config.yaml") }
func SavePath(dir string) string   { return filepath.Join(dir, "save.json") }
func PackPath(dir string) string   { return filepath.Join(dir, "packs", "core.yaml") }

func Init() (string, error) {
	dir, err := Dir()
	if err != nil {
		return "", err
	}

	// Create the directories with the following permisisons
	// Owner - rwx (7)
	// Group & Others - r-x (5)
	if err := os.MkdirAll(filepath.Join(dir, "packs"), 0o755); err != nil {
		return "", fmt.Errorf("Cannot create game directory: %w", err)
	}

	if _, err := os.Stat(SavePath(dir)); os.IsNotExist(err) {
		if err := save.Write(SavePath(dir), model.DefaultSave()); err != nil {
			return "", err
		}
	}

	return dir, nil
}
