package recurring

import (
	"testing"
	"time"

	"github.com/joshuataye/ledgerpipeline/internal/parser"
)

func TestIntervals(t *testing.T) {
	d := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	txns := []parser.Transaction{
		{Date: d, Description: "Netflix", Category: "Entertainment", Amount: -15},
		{Date: d.AddDate(0, 1, 0), Description: "Netflix", Category: "Entertainment", Amount: -15},
		{Date: d.AddDate(0, 2, 0), Description: "Netflix", Category: "Entertainment", Amount: -15},
	}
	intervals := Intervals(txns, 2)
	if len(intervals) != 1 || intervals[0].Days < 28 {
		t.Fatalf("Intervals() = %+v", intervals)
	}
}
