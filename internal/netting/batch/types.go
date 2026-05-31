package batch

import "time"

// NettingSummary holds aggregated metrics for netting/batch analysis.
type NettingSummary struct {
    Count      int
    Net        float64
    DebitTotal float64
    CreditTotal float64
    From       time.Time
    To         time.Time
}
