package partial

import "time"

// RefundSummary holds aggregated metrics for refund/partial analysis.
type RefundSummary struct {
    Count      int
    Net        float64
    DebitTotal float64
    CreditTotal float64
    From       time.Time
    To         time.Time
}
