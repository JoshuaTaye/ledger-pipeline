package accrue

import "time"

// ExpenseSummary holds aggregated metrics for expense/accrue analysis.
type ExpenseSummary struct {
    Count      int
    Net        float64
    DebitTotal float64
    CreditTotal float64
    From       time.Time
    To         time.Time
}
