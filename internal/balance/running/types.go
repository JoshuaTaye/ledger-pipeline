package running

import "time"

// BalanceSummary holds aggregated metrics for balance/running analysis.
type BalanceSummary struct {
    Count      int
    Net        float64
    DebitTotal float64
    CreditTotal float64
    From       time.Time
    To         time.Time
}
