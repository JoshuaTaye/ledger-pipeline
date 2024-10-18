package period

import (
	"fmt"
	"time"

	"github.com/joshuataye/ledgerpipeline/internal/parser"
)

// MonthlyRollup groups transactions by calendar month.
type MonthlyRollup struct {
	Year  int
	Month time.Month
	Count int
	Net   float64
}

func ByMonth(txns []parser.Transaction) []MonthlyRollup {
	type key struct{ y int; m time.Month }
	buckets := map[key]*MonthlyRollup{}
	for _, tx := range txns {
		k := key{tx.Date.Year(), tx.Date.Month()}
		r, ok := buckets[k]
		if !ok {
			r = &MonthlyRollup{Year: k.y, Month: k.m}
			buckets[k] = r
		}
		r.Count++
		r.Net += tx.Amount
	}
	out := make([]MonthlyRollup, 0, len(buckets))
	for _, r := range buckets {
		out = append(out, *r)
	}
	return out
}

func (m MonthlyRollup) Label() string {
	return fmt.Sprintf("%04d-%02d", m.Year, m.Month)
}
