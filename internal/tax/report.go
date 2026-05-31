package tax

import (
	"fmt"
	"io"
	"sort"

	"github.com/joshuataye/ledgerpipeline/internal/parser"
)

// Category totals for tax-deductible spend.
var DeductibleCategories = map[string]struct{}{
	"Business": {},
	"Medical":  {},
	"Charity":  {},
}

// Line is one deductible category total.
type Line struct {
	Category string
	Total    float64
}

// Report summarizes deductible category spend.
func Report(txns []parser.Transaction) []Line {
	totals := map[string]float64{}
	for _, tx := range txns {
		if tx.Amount >= 0 {
			continue
		}
		if _, ok := DeductibleCategories[tx.Category]; ok {
			totals[tx.Category] += tx.Amount
		}
	}
	out := make([]Line, 0, len(totals))
	for cat, total := range totals {
		out = append(out, Line{Category: cat, Total: total})
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Category < out[j].Category })
	return out
}

// WriteReport prints deductible lines to w.
func WriteReport(w io.Writer, lines []Line) error {
	for _, line := range lines {
		if _, err := fmt.Fprintf(w, "%s: %.2f\n", line.Category, line.Total); err != nil {
			return err
		}
	}
	return nil
}
