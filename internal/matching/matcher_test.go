package matching

import (
	"testing"
	"time"

	"github.com/joshuataye/ledgerpipeline/internal/parser"
)

func TestDefaultOptions(t *testing.T) {
	opt := DefaultOptions()
	if opt.MaxDays != 3 || opt.Tolerance != 0.01 {
		t.Fatalf("DefaultOptions() = %+v", opt)
	}
}

func TestFindTransfers(t *testing.T) {
	d := time.Date(2026, 1, 10, 0, 0, 0, 0, time.UTC)
	txns := []parser.Transaction{
		{Date: d, Description: "Transfer to savings", Amount: -500},
		{Date: d.Add(24 * time.Hour), Description: "Transfer from checking", Amount: 500},
		{Date: d, Description: "Coffee", Amount: -5},
	}
	pairs := FindTransfers(txns, DefaultOptions())
	if len(pairs) != 1 {
		t.Fatalf("FindTransfers() len = %d", len(pairs))
	}
	unmatched := Unmatched(txns, pairs)
	if len(unmatched) != 1 || unmatched[0].Description != "Coffee" {
		t.Fatalf("Unmatched() = %+v", unmatched)
	}
}
