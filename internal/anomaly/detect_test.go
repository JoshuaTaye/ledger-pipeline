package anomaly

import (
	"testing"
	"time"

	"github.com/joshuataye/ledgerpipeline/internal/parser"
)

func TestDetect(t *testing.T) {
	d := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	txns := []parser.Transaction{
		{Date: d, Amount: -10},
		{Date: d, Amount: -12},
		{Date: d, Amount: -11},
		{Date: d, Amount: -500},
	}
	hits := Detect(txns, 1.5)
	if len(hits) != 1 || hits[0].Transaction.Amount != -500 {
		t.Fatalf("Detect() = %+v", hits)
	}
}
