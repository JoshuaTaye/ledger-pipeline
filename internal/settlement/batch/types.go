package batch

import "time"

// SettlementSummary holds aggregated metrics for settlement/batch analysis.
type SettlementSummary struct {
    Count      int
    Net        float64
    DebitTotal float64
    CreditTotal float64
    From       time.Time
    To         time.Time
}
