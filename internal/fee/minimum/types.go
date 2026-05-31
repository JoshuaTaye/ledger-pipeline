package minimum

import "time"

// FeeSummary holds aggregated metrics for fee/minimum analysis.
type FeeSummary struct {
    Count      int
    Net        float64
    DebitTotal float64
    CreditTotal float64
    From       time.Time
    To         time.Time
}
