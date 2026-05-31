package score

import "time"

// CreditSummary holds aggregated metrics for credit/score analysis.
type CreditSummary struct {
    Count      int
    Net        float64
    DebitTotal float64
    CreditTotal float64
    From       time.Time
    To         time.Time
}
