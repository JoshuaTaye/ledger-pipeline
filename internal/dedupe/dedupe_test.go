package dedupe

import (
	"testing"
	"time"

	"github.com/joshuataye/ledgerpipeline/internal/parser"
)

func TestRemoveDuplicates(t *testing.T) {
	d := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	txns := []parser.Transaction{
		{Date: d, Description: "A", Category: "Food", Amount: -1},
		{Date: d, Description: "A", Category: "Food", Amount: -1},
	}
	got := RemoveDuplicates(txns)
	if len(got) != 1 {
		t.Fatalf("len %d", len(got))
	}
}
