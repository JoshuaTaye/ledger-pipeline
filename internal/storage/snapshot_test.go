package storage

import (
	"path/filepath"
	"testing"
	"time"

	"github.com/joshuataye/ledgerpipeline/internal/parser"
)

func TestSaveLoad(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "snap.json")
	d := time.Date(2026, 1, 3, 0, 0, 0, 0, time.UTC)
	orig := []parser.Transaction{
		{Date: d, Description: "Coffee", Category: "Food", Amount: -5},
	}
	if err := Save(path, orig); err != nil {
		t.Fatal(err)
	}
	loaded, err := Load(path)
	if err != nil || len(loaded) != 1 || loaded[0].Amount != -5 {
		t.Fatalf("Load() = %v, %v", loaded, err)
	}
}
