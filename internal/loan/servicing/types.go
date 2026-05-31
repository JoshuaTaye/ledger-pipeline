package servicing

import "time"

// LoanSummary holds aggregated metrics for loan/servicing analysis.
type LoanSummary struct {
    Count      int
    Net        float64
    DebitTotal float64
    CreditTotal float64
    From       time.Time
    To         time.Time
}
