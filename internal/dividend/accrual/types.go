package accrual

import "time"

// DividendSummary holds aggregated metrics for dividend/accrual analysis.
type DividendSummary struct {
    Count      int
    Net        float64
    DebitTotal float64
    CreditTotal float64
    From       time.Time
    To         time.Time
}
