package prepay

import "time"

// LoanSummary holds aggregated metrics for loan/prepay analysis.
type LoanSummary struct {
    Count      int
    Net        float64
    DebitTotal float64
    CreditTotal float64
    From       time.Time
    To         time.Time
}
