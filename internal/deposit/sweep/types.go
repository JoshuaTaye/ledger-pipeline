package sweep

import "time"

// DepositSummary holds aggregated metrics for deposit/sweep analysis.
type DepositSummary struct {
    Count      int
    Net        float64
    DebitTotal float64
    CreditTotal float64
    From       time.Time
    To         time.Time
}
