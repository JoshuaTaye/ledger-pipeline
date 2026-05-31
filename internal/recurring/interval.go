package recurring

import (
	"sort"
	"time"

	"github.com/joshuataye/ledgerpipeline/internal/parser"
)

// Interval describes cadence between recurring charges.
type Interval struct {
	Description string
	Category    string
	Days        int
	Count       int
}

// Intervals estimates day gaps between same-description transactions.
func Intervals(txns []parser.Transaction, minCount int) []Interval {
	type key struct{ desc, cat string }
	dates := map[key][]time.Time{}
	for _, tx := range txns {
		k := key{tx.Description, tx.Category}
		dates[k] = append(dates[k], tx.Date)
	}
	var out []Interval
	for k, ds := range dates {
		if len(ds) < minCount {
			continue
		}
		sort.Slice(ds, func(i, j int) bool { return ds[i].Before(ds[j]) })
		var totalDays int
		for i := 1; i < len(ds); i++ {
			totalDays += int(ds[i].Sub(ds[i-1]).Hours() / 24)
		}
		avg := totalDays / (len(ds) - 1)
		out = append(out, Interval{
			Description: k.desc,
			Category:    k.cat,
			Days:        avg,
			Count:       len(ds),
		})
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Days < out[j].Days })
	return out
}
