package runway

import "time"

// CashflowSummary holds aggregated metrics for cashflow/runway analysis.
type CashflowSummary struct {
    Count      int
    Net        float64
    DebitTotal float64
    CreditTotal float64
    From       time.Time
    To         time.Time
}
