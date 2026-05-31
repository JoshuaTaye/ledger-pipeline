package importfmt

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDefaultFixedWidthLayout(t *testing.T) {
	layout := DefaultFixedWidthLayout()
	if layout.Date.End != 10 || layout.Amount.Start != 30 {
		t.Fatalf("DefaultFixedWidthLayout() = %+v", layout)
	}
}

func TestDetect(t *testing.T) {
	dir := t.TempDir()
	csvPath := filepath.Join(dir, "sample.csv")
	if err := os.WriteFile(csvPath, []byte("date,description,category,amount\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	format, err := Detect(csvPath)
	if err != nil || format != FormatCSV {
		t.Fatalf("Detect() = %v, %v", format, err)
	}
}

func TestParseFileFixedWidth(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "export.txt")
	line := "2026-01-05Grocery             -12.50    \n"
	if err := os.WriteFile(path, []byte(line), 0o644); err != nil {
		t.Fatal(err)
	}
	txns, err := ParseFile(path)
	if err != nil || len(txns) != 1 || txns[0].Amount != -12.5 {
		t.Fatalf("ParseFile() = %v, %v", txns, err)
	}
}
