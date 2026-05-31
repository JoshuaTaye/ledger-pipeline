package compare

import "github.com/joshuataye/ledgerpipeline/internal/period"

// Delta is the change between two monthly rollups with the same label.
type Delta struct {
	Label    string
	Previous float64
	Current  float64
	Change   float64
}

// Periods compares monthly nets between two transaction sets.
func Periods(before, after []period.MonthlyRollup) []Delta {
	prev := map[string]float64{}
	for _, r := range before {
		prev[r.Label()] = r.Net
	}
	cur := map[string]float64{}
	for _, r := range after {
		cur[r.Label()] = r.Net
	}
	seen := map[string]struct{}{}
	var out []Delta
	for label, p := range prev {
		c := cur[label]
		out = append(out, Delta{Label: label, Previous: p, Current: c, Change: c - p})
		seen[label] = struct{}{}
	}
	for label, c := range cur {
		if _, ok := seen[label]; ok {
			continue
		}
		out = append(out, Delta{Label: label, Current: c, Change: c})
	}
	return out
}
