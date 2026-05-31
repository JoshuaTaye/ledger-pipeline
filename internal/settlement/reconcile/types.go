package reconcile

import "time"

// SettlementSummary holds aggregated metrics for settlement/reconcile analysis.
type SettlementSummary struct {
    Count      int
    Net        float64
    DebitTotal float64
    CreditTotal float64
    From       time.Time
    To         time.Time
}
