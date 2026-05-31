package amortize

import "time"

// ExpenseSummary holds aggregated metrics for expense/amortize analysis.
type ExpenseSummary struct {
    Count      int
    Net        float64
    DebitTotal float64
    CreditTotal float64
    From       time.Time
    To         time.Time
}
