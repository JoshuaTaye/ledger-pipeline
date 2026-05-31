package release

import "time"

// DepositSummary holds aggregated metrics for deposit/release analysis.
type DepositSummary struct {
    Count      int
    Net        float64
    DebitTotal float64
    CreditTotal float64
    From       time.Time
    To         time.Time
}
