package accrual

import "time"

// InterestSummary holds aggregated metrics for interest/accrual analysis.
type InterestSummary struct {
    Count      int
    Net        float64
    DebitTotal float64
    CreditTotal float64
    From       time.Time
    To         time.Time
}
