package rules

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadFile(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "rules.json")
	if err := os.WriteFile(path, []byte(`[
		{"category":"Food","contains":"COFFEE","priority":10},
		{"category":"Transport","contains":"UBER","priority":5}
	]`), 0o644); err != nil {
		t.Fatal(err)
	}
	got, err := LoadFile(path)
	if err != nil || len(got) != 2 || got[0].Category != "Food" {
		t.Fatalf("LoadFile() = %v, %v", got, err)
	}
}
