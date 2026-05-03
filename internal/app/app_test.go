package app

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/dcxforge/bugcraft/internal/save"
)

func TestDirUsesBugcraftHomeOverride(t *testing.T) {
	want := filepath.Join(t.TempDir(), "custom-home")
	t.Setenv("BUGCRAFT_HOME", want)

	got, err := Dir()
	if err != nil {
		t.Fatalf("Dir() error = %v", err)
	}
	if got != want {
		t.Fatalf("Dir() = %q, want %q", got, want)
	}
}

func TestPathHelpers(t *testing.T) {
	t.Parallel()

	dir := filepath.Join("tmp", "bugcraft")

	if got := ConfigPath(dir); got != filepath.Join(dir, "config.yaml") {
		t.Fatalf("ConfigPath() = %q", got)
	}
	if got := SavePath(dir); got != filepath.Join(dir, "save.json") {
		t.Fatalf("SavePath() = %q", got)
	}
	if got := PackPath(dir); got != filepath.Join(dir, "packs", "core.yaml") {
		t.Fatalf("PackPath() = %q", got)
	}
}

func TestInitCreatesGameDirectoryAndDefaultSave(t *testing.T) {
	dir := filepath.Join(t.TempDir(), "bugcraft-home")
	t.Setenv("BUGCRAFT_HOME", dir)

	got, err := Init()
	if err != nil {
		t.Fatalf("Init() error = %v", err)
	}
	if got != dir {
		t.Fatalf("Init() dir = %q, want %q", got, dir)
	}
	if info, err := os.Stat(filepath.Join(dir, "packs")); err != nil {
		t.Fatalf("packs directory stat error = %v", err)
	} else if !info.IsDir() {
		t.Fatal("packs path is not a directory")
	}

	s, err := save.Load(SavePath(dir))
	if err != nil {
		t.Fatalf("Load(default save) error = %v", err)
	}
	if s.Player.Name != "Sir GREPalot" {
		t.Fatalf("default player name = %q", s.Player.Name)
	}
}

func TestInitPreservesExistingSave(t *testing.T) {
	dir := filepath.Join(t.TempDir(), "bugcraft-home")
	t.Setenv("BUGCRAFT_HOME", dir)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatalf("MkdirAll() error = %v", err)
	}
	existing := []byte(`{"player":{"name":"Grace","hp":3,"max_hp":10}}`)
	if err := os.WriteFile(SavePath(dir), existing, 0o644); err != nil {
		t.Fatalf("WriteFile() error = %v", err)
	}

	if _, err := Init(); err != nil {
		t.Fatalf("Init() error = %v", err)
	}

	data, err := os.ReadFile(SavePath(dir))
	if err != nil {
		t.Fatalf("ReadFile() error = %v", err)
	}
	if string(data) != string(existing) {
		t.Fatalf("save was overwritten: got %q, want %q", string(data), string(existing))
	}
}
