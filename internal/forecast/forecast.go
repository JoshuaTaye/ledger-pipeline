package forecast

import (
	"github.com/joshuataye/ledgerpipeline/internal/parser"
	"github.com/joshuataye/ledgerpipeline/internal/period"
)

// Point is a projected monthly net total.
type Point struct {
	Label string
	Net   float64
}

// Linear projects the next month from recent monthly rollups.
func Linear(txns []parser.Transaction, months int) []Point {
	rollups := period.ByMonth(txns)
	if len(rollups) == 0 {
		return nil
	}
	start := 0
	if len(rollups) > months {
		start = len(rollups) - months
	}
	window := rollups[start:]
	var sum float64
	for _, r := range window {
		sum += r.Net
	}
	avg := sum / float64(len(window))
	out := make([]Point, len(window))
	for i, r := range window {
		out[i] = Point{Label: r.Label(), Net: r.Net}
	}
	out = append(out, Point{Label: "forecast", Net: avg})
	return out
}
