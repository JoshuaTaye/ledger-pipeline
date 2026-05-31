package allocate

import "time"

// RevenueSummary holds aggregated metrics for revenue/allocate analysis.
type RevenueSummary struct {
    Count      int
    Net        float64
    DebitTotal float64
    CreditTotal float64
    From       time.Time
    To         time.Time
}
