package baddebt

import "time"

// WriteoffSummary holds aggregated metrics for writeoff/baddebt analysis.
type WriteoffSummary struct {
    Count      int
    Net        float64
    DebitTotal float64
    CreditTotal float64
    From       time.Time
    To         time.Time
}
