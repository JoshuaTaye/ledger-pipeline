package payout

import "time"

// DividendSummary holds aggregated metrics for dividend/payout analysis.
type DividendSummary struct {
    Count      int
    Net        float64
    DebitTotal float64
    CreditTotal float64
    From       time.Time
    To         time.Time
}
