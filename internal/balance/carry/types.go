package carry

import "time"

// BalanceSummary holds aggregated metrics for balance/carry analysis.
type BalanceSummary struct {
    Count      int
    Net        float64
    DebitTotal float64
    CreditTotal float64
    From       time.Time
    To         time.Time
}
