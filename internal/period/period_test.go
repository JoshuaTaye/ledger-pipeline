package period

import (
	"testing"
	"time"

	"github.com/joshuataye/ledgerpipeline/internal/parser"
)

func TestByMonth(t *testing.T) {
	txns := []parser.Transaction{
		{Date: time.Date(2026, 1, 5, 0, 0, 0, 0, time.UTC), Amount: 10},
		{Date: time.Date(2026, 1, 20, 0, 0, 0, 0, time.UTC), Amount: -3},
	}
	rollups := ByMonth(txns)
	if len(rollups) != 1 || rollups[0].Net != 7 {
		t.Fatalf("got %+v", rollups)
	}
}
