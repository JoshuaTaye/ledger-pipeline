package batchfile

import (
	"os"
	"path/filepath"
	"testing"
)

func TestReadDir(t *testing.T) {
	dir := t.TempDir()
	content := "date,description,category,amount\n2026-01-01,A,Food,-1.00\n"
	if err := os.WriteFile(filepath.Join(dir, "a.csv"), []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
	txns, err := ReadDir(dir)
	if err != nil || len(txns) != 1 {
		t.Fatalf("err %v len %d", err, len(txns))
	}
}
