package budget

import "fmt"

// Line is a budget target for one category.
type Line struct {
	Category string
	Limit    float64
}

// Variance is actual spend vs budget.
type Variance struct {
	Category string
	Limit    float64
	Actual   float64
	Delta    float64
}

func Compare(actual map[string]float64, lines []Line) []Variance {
	out := make([]Variance, 0, len(lines))
	for _, line := range lines {
		spent := actual[line.Category]
		out = append(out, Variance{
			Category: line.Category,
			Limit:    line.Limit,
			Actual:   spent,
			Delta:    line.Limit - spent,
		})
	}
	return out
}

func FormatVariance(v Variance) string {
	return fmt.Sprintf("%s: limit %.2f actual %.2f remaining %.2f", v.Category, v.Limit, v.Actual, v.Delta)
}
