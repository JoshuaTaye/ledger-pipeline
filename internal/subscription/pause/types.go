package pause

import "time"

// SubscriptionSummary holds aggregated metrics for subscription/pause analysis.
type SubscriptionSummary struct {
    Count      int
    Net        float64
    DebitTotal float64
    CreditTotal float64
    From       time.Time
    To         time.Time
}
