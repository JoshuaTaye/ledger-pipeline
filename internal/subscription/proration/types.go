package proration

import "time"

// SubscriptionSummary holds aggregated metrics for subscription/proration analysis.
type SubscriptionSummary struct {
    Count      int
    Net        float64
    DebitTotal float64
    CreditTotal float64
    From       time.Time
    To         time.Time
}
