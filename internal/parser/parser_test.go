package parser

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestParseReader_ValidCSV(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantLen int
		want    Transaction
	}{
		{
			name: "with header",
			input: `date,description,category,amount
2026-01-15,Coffee Shop,Food,-4.50
`,
			wantLen: 1,
			want: Transaction{
				Date:        time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC),
				Description: "Coffee Shop",
				Category:    "Food",
				Amount:      -4.50,
			},
		},
		{
			name: "without header",
			input: `2026-02-01,Payroll,Income,3200.00
`,
			wantLen: 1,
			want: Transaction{
				Date:        time.Date(2026, 2, 1, 0, 0, 0, 0, time.UTC),
				Description: "Payroll",
				Category:    "Income",
				Amount:      3200.00,
			},
		},
		{
			name: "multiple rows",
			input: `date,description,category,amount
2026-01-01,Gas,Transport,-60.00
2026-01-02,Refund,Shopping,15.25
`,
			wantLen: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseReader(strings.NewReader(tt.input))
			if err != nil {
				t.Fatalf("ParseReader() error = %v", err)
			}
			if len(got) != tt.wantLen {
				t.Fatalf("len = %d, want %d", len(got), tt.wantLen)
			}
			if tt.want.Description != "" {
				if got[0] != tt.want {
					t.Fatalf("first transaction = %+v, want %+v", got[0], tt.want)
				}
			}
		})
	}
}

func TestParseReader_EmptyFile(t *testing.T) {
	got, err := ParseReader(strings.NewReader(""))
	if err != nil {
		t.Fatalf("ParseReader() error = %v", err)
	}
	if got != nil {
		t.Fatalf("got %v, want nil", got)
	}
}

func TestParseReader_InvalidAmount(t *testing.T) {
	input := `date,description,category,amount
2026-01-01,Bad,Food,not-a-number
`
	_, err := ParseReader(strings.NewReader(input))
	if err == nil {
		t.Fatal("expected error for invalid amount")
	}
}

func TestParseReader_InvalidDate(t *testing.T) {
	input := `date,description,category,amount
01/15/2026,Coffee,Food,-4.50
`
	_, err := ParseReader(strings.NewReader(input))
	if err == nil {
		t.Fatal("expected error for invalid date")
	}
}

func TestParseReader_MissingFields(t *testing.T) {
	input := `date,description,category,amount
2026-01-01,Incomplete,Food
`
	_, err := ParseReader(strings.NewReader(input))
	if err == nil {
		t.Fatal("expected error for missing fields")
	}
}

func TestParseReader_SkipsBlankRows(t *testing.T) {
	input := `date,description,category,amount

2026-01-01,Valid,Food,-1.00

`
	got, err := ParseReader(strings.NewReader(input))
	if err != nil {
		t.Fatalf("ParseReader() error = %v", err)
	}
	if len(got) != 1 {
		t.Fatalf("len = %d, want 1", len(got))
	}
}

func TestParseFile(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "statement.csv")
	content := `date,description,category,amount
2026-03-01,Rent,Housing,-1200.00
`
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}

	got, err := ParseFile(path)
	if err != nil {
		t.Fatalf("ParseFile() error = %v", err)
	}
	if len(got) != 1 {
		t.Fatalf("len = %d, want 1", len(got))
	}
	if got[0].Category != "Housing" {
		t.Fatalf("category = %q, want Housing", got[0].Category)
	}
}

func TestParseFile_NotFound(t *testing.T) {
	_, err := ParseFile(filepath.Join(t.TempDir(), "missing.csv"))
	if err == nil {
		t.Fatal("expected error for missing file")
	}
}

func TestIsHeader(t *testing.T) {
	tests := []struct {
		name string
		row  []string
		want bool
	}{
		{name: "header row", row: []string{"date", "description", "category", "amount"}, want: true},
		{name: "mixed case", row: []string{"Date", "Description", "Category", "Amount"}, want: true},
		{name: "data row", row: []string{"2026-01-01", "Coffee", "Food", "-3.00"}, want: false},
		{name: "too short", row: []string{"date", "description"}, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isHeader(tt.row); got != tt.want {
				t.Fatalf("isHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}
