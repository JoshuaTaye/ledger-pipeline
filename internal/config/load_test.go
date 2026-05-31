package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoad(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "config.json")
	content := `{"min_date":"2026-01-01","max_date":"2026-01-31","categories":["Food"]}`
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
	cfg, err := Load(path)
	if err != nil || cfg.MinDate != "2026-01-01" || len(cfg.Categories) != 1 {
		t.Fatalf("Load() = %+v, %v", cfg, err)
	}
}
