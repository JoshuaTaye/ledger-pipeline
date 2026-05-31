package peer

import "time"

// BenchmarkSummary holds aggregated metrics for benchmark/peer analysis.
type BenchmarkSummary struct {
    Count      int
    Net        float64
    DebitTotal float64
    CreditTotal float64
    From       time.Time
    To         time.Time
}
