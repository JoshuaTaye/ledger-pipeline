package split

import "github.com/joshuataye/ledgerpipeline/internal/parser"

// Part is one share of a split transaction.
type Part struct {
	Category string
	Amount   float64
}

// ByCategory divides a transaction amount across categories by weight.
func ByCategory(tx parser.Transaction, weights map[string]float64) []Part {
	var totalWeight float64
	for _, w := range weights {
		totalWeight += w
	}
	if totalWeight == 0 {
		return nil
	}
	var parts []Part
	for cat, w := range weights {
		share := tx.Amount * (w / totalWeight)
		parts = append(parts, Part{Category: cat, Amount: share})
	}
	return parts
}

// Even divides a transaction equally across categories.
func Even(tx parser.Transaction, categories []string) []Part {
	if len(categories) == 0 {
		return nil
	}
	share := tx.Amount / float64(len(categories))
	parts := make([]Part, len(categories))
	for i, cat := range categories {
		parts[i] = Part{Category: cat, Amount: share}
	}
	return parts
}
