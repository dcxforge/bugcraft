package cmd

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/dcxforge/bugcraft/internal/app"
	"github.com/dcxforge/bugcraft/internal/save"
)

func TestStartCommandInitializesAppAndPrintsWelcome(t *testing.T) {
	dir := filepath.Join(t.TempDir(), "bugcraft-home")
	t.Setenv("BUGCRAFT_HOME", dir)

	output := captureStdout(t, func() {
		if err := startCmd.RunE(startCmd, nil); err != nil {
			t.Fatalf("start RunE error = %v", err)
		}
	})

	for _, want := range []string{
		"Welcome to BugCraft",
		"You inherited a tiny farm at the edge of a haunted codebase.",
		dir,
	} {
		if !strings.Contains(output, want) {
			t.Fatalf("start output = %q, want it to contain %q", output, want)
		}
	}

	if _, err := save.Load(app.SavePath(dir)); err != nil {
		t.Fatalf("Load(initialized save) error = %v", err)
	}
}

func captureStdout(t *testing.T, fn func()) string {
	t.Helper()

	old := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("Pipe() error = %v", err)
	}
	os.Stdout = w

	fn()

	if err := w.Close(); err != nil {
		t.Fatalf("Close() error = %v", err)
	}
	os.Stdout = old

	var buf bytes.Buffer
	if _, err := io.Copy(&buf, r); err != nil {
		t.Fatalf("Copy() error = %v", err)
	}
	if err := r.Close(); err != nil {
		t.Fatalf("Close() error = %v", err)
	}
	return buf.String()
}
