package compare

import (
	"testing"
	"time"

	"github.com/joshuataye/ledgerpipeline/internal/period"
)

func TestPeriods(t *testing.T) {
	before := []period.MonthlyRollup{
		{Year: 2026, Month: time.January, Net: -100},
	}
	after := []period.MonthlyRollup{
		{Year: 2026, Month: time.January, Net: -150},
	}
	deltas := Periods(before, after)
	if len(deltas) != 1 || deltas[0].Change != -50 {
		t.Fatalf("Periods() = %+v", deltas)
	}
}
