package accountsfile

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoad(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "accounts.json")
	content := `[
		{"id":"checking","name":"Checking","type":"checking","opening":1000},
		{"id":"savings","name":"Savings","type":"savings","opening":5000}
	]`
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
	accounts, err := Load(path)
	if err != nil || len(accounts) != 2 || accounts[0].ID != "checking" {
		t.Fatalf("Load() = %v, %v", accounts, err)
	}
	reg, err := LoadRegistry(path)
	if err != nil || reg == nil {
		t.Fatalf("LoadRegistry() = %v, %v", reg, err)
	}
}
