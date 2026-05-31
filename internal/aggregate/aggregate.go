package aggregate

import (
	"sort"

	"github.com/joshuataye/ledgerpipeline/internal/parser"
)

// CategorySummary holds rolled-up totals for one category.
type CategorySummary struct {
	Category         string
	TransactionCount int
	TotalAmount      float64
	DebitTotal       float64
	CreditTotal      float64
}

// ByCategory groups transactions and computes per-category totals.
func ByCategory(transactions []parser.Transaction) []CategorySummary {
	counts := make(map[string]*CategorySummary)

	for _, tx := range transactions {
		category := tx.Category
		summary, ok := counts[category]
		if !ok {
			summary = &CategorySummary{Category: category}
			counts[category] = summary
		}

		summary.TransactionCount++
		summary.TotalAmount += tx.Amount
		if tx.Amount < 0 {
			summary.DebitTotal += tx.Amount
		} else {
			summary.CreditTotal += tx.Amount
		}
	}

	out := make([]CategorySummary, 0, len(counts))
	for _, summary := range counts {
		out = append(out, *summary)
	}

	sort.Slice(out, func(i, j int) bool {
		return out[i].Category < out[j].Category
	})

	return out
}

// NetTotal returns the sum of all transaction amounts.
func NetTotal(transactions []parser.Transaction) float64 {
	var total float64
	for _, tx := range transactions {
		total += tx.Amount
	}
	return total
}
