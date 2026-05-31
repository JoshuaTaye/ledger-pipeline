package assign

import "time"

// EnvelopeSummary holds aggregated metrics for envelope/assign analysis.
type EnvelopeSummary struct {
    Count      int
    Net        float64
    DebitTotal float64
    CreditTotal float64
    From       time.Time
    To         time.Time
}
