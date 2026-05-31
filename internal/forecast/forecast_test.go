package forecast

import (
	"testing"
	"time"

	"github.com/joshuataye/ledgerpipeline/internal/parser"
)

func TestLinear(t *testing.T) {
	txns := []parser.Transaction{
		{Date: time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC), Amount: -100},
		{Date: time.Date(2026, 2, 15, 0, 0, 0, 0, time.UTC), Amount: -200},
	}
	points := Linear(txns, 2)
	if len(points) < 2 {
		t.Fatalf("Linear() = %+v", points)
	}
	last := points[len(points)-1]
	if last.Label != "forecast" {
		t.Fatalf("forecast label = %q", last.Label)
	}
}
