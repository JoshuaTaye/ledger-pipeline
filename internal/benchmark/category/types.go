package category

import "time"

// BenchmarkSummary holds aggregated metrics for benchmark/category analysis.
type BenchmarkSummary struct {
    Count      int
    Net        float64
    DebitTotal float64
    CreditTotal float64
    From       time.Time
    To         time.Time
}
