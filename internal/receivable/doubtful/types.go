package doubtful

import "time"

// ReceivableSummary holds aggregated metrics for receivable/doubtful analysis.
type ReceivableSummary struct {
    Count      int
    Net        float64
    DebitTotal float64
    CreditTotal float64
    From       time.Time
    To         time.Time
}
