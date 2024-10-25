package recurring

import (
	"testing"
	"time"

	"github.com/joshuataye/ledgerpipeline/internal/parser"
)

func TestDetect(t *testing.T) {
	d := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	txns := []parser.Transaction{
		{Date: d, Description: "Netflix", Category: "Subs", Amount: -15},
		{Date: d.AddDate(0, 1, 0), Description: "Netflix", Category: "Subs", Amount: -15},
		{Date: d.AddDate(0, 2, 0), Description: "Netflix", Category: "Subs", Amount: -15},
	}
	c := Detect(txns, 3)
	if len(c) != 1 {
		t.Fatalf("got %d", len(c))
	}
}
