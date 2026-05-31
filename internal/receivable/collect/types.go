package collect

import "time"

// ReceivableSummary holds aggregated metrics for receivable/collect analysis.
type ReceivableSummary struct {
    Count      int
    Net        float64
    DebitTotal float64
    CreditTotal float64
    From       time.Time
    To         time.Time
}
