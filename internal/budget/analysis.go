package budget

import "math"

// Utilization is budget usage for one category.
type Utilization struct {
	Category string
	Limit    float64
	Actual   float64
	UsedPct  float64
}

// Analyze computes utilization from budget variances using absolute spend.
func Analyze(variances []Variance) []Utilization {
	out := make([]Utilization, 0, len(variances))
	for _, v := range variances {
		spent := math.Abs(v.Actual)
		var pct float64
		if v.Limit > 0 {
			pct = spent / v.Limit * 100
		}
		out = append(out, Utilization{
			Category: v.Category,
			Limit:    v.Limit,
			Actual:   v.Actual,
			UsedPct:  pct,
		})
	}
	return out
}

// OverBudget returns categories where absolute actual exceeds the limit.
func OverBudget(variances []Variance) []Variance {
	var out []Variance
	for _, v := range variances {
		if v.Limit > 0 && math.Abs(v.Actual) > v.Limit {
			out = append(out, v)
		}
	}
	return out
}
