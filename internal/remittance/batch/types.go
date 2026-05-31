package batch

import "time"

// RemittanceSummary holds aggregated metrics for remittance/batch analysis.
type RemittanceSummary struct {
    Count      int
    Net        float64
    DebitTotal float64
    CreditTotal float64
    From       time.Time
    To         time.Time
}
