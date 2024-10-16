package filter

import (
	"testing"
	"time"

	"github.com/joshuataye/ledgerpipeline/internal/parser"
)

func TestApply_DateRange(t *testing.T) {
	txns := []parser.Transaction{
		{Date: time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC), Category: "Food", Amount: -1},
		{Date: time.Date(2026, 2, 1, 0, 0, 0, 0, time.UTC), Category: "Food", Amount: -2},
	}
	got := Apply(txns, Options{From: time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC)})
	if len(got) != 1 {
		t.Fatalf("len %d", len(got))
	}
}
