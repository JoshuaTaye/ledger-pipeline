package burn

import "time"

// CashflowSummary holds aggregated metrics for cashflow/burn analysis.
type CashflowSummary struct {
    Count      int
    Net        float64
    DebitTotal float64
    CreditTotal float64
    From       time.Time
    To         time.Time
}
