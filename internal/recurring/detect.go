package recurring

import (
	"sort"

	"github.com/joshuataye/ledgerpipeline/internal/parser"
)

// Candidate is a possible subscription or bill.
type Candidate struct {
	Description string
	Category    string
	Count       int
	AvgAmount   float64
}

func Detect(txns []parser.Transaction, minCount int) []Candidate {
	type key struct{ desc, cat string }
	buckets := map[key][]float64{}
	for _, tx := range txns {
		k := key{tx.Description, tx.Category}
		buckets[k] = append(buckets[k], tx.Amount)
	}
	var out []Candidate
	for k, amounts := range buckets {
		if len(amounts) < minCount {
			continue
		}
		var sum float64
		for _, a := range amounts {
			sum += a
		}
		out = append(out, Candidate{
			Description: k.desc,
			Category:    k.cat,
			Count:       len(amounts),
			AvgAmount:   sum / float64(len(amounts)),
		})
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Count > out[j].Count })
	return out
}
