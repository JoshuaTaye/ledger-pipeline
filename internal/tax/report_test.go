package tax

import (
	"testing"
	"time"

	"github.com/joshuataye/ledgerpipeline/internal/parser"
)

func TestReport(t *testing.T) {
	d := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	txns := []parser.Transaction{
		{Date: d, Category: "Business", Amount: -100},
		{Date: d, Category: "Food", Amount: -50},
		{Date: d, Category: "Medical", Amount: -30},
	}
	lines := Report(txns)
	if len(lines) != 2 {
		t.Fatalf("Report() len = %d", len(lines))
	}
}
