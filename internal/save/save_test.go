package save

import (
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/dcxforge/bugcraft/internal/model"
)

func TestWriteAndLoadRoundTrip(t *testing.T) {
	t.Parallel()

	path := filepath.Join(t.TempDir(), "save.json")
	want := model.DefaultSave()
	want.Player.Name = "Ada"
	want.Inventory["coffee"] = 7

	if err := Write(path, want); err != nil {
		t.Fatalf("Write() error = %v", err)
	}

	got, err := Load(path)
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}

	if got.Player.Name != want.Player.Name {
		t.Fatalf("Player.Name = %q, want %q", got.Player.Name, want.Player.Name)
	}
	if got.Inventory["coffee"] != want.Inventory["coffee"] {
		t.Fatalf("coffee inventory = %d, want %d", got.Inventory["coffee"], want.Inventory["coffee"])
	}
	if len(got.Farm.Plots) != len(want.Farm.Plots) {
		t.Fatalf("farm plot count = %d, want %d", len(got.Farm.Plots), len(want.Farm.Plots))
	}
}

func TestLoadMissingFileReturnsError(t *testing.T) {
	t.Parallel()

	_, err := Load(filepath.Join(t.TempDir(), "missing.json"))
	if err == nil {
		t.Fatal("Load() error = nil, want error")
	}
	if !errors.Is(err, os.ErrNotExist) {
		t.Fatalf("Load() error = %v, want os.ErrNotExist", err)
	}
}

func TestLoadInvalidJSONReturnsError(t *testing.T) {
	t.Parallel()

	path := filepath.Join(t.TempDir(), "save.json")
	if err := os.WriteFile(path, []byte("{"), 0o644); err != nil {
		t.Fatalf("WriteFile() error = %v", err)
	}

	if _, err := Load(path); err == nil {
		t.Fatal("Load() error = nil, want error")
	}
}

func TestWriteMissingParentReturnsError(t *testing.T) {
	t.Parallel()

	err := Write(filepath.Join(t.TempDir(), "missing", "save.json"), model.DefaultSave())
	if err == nil {
		t.Fatal("Write() error = nil, want error")
	}
}
