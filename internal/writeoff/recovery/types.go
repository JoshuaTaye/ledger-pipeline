package recovery

import "time"

// WriteoffSummary holds aggregated metrics for writeoff/recovery analysis.
type WriteoffSummary struct {
    Count      int
    Net        float64
    DebitTotal float64
    CreditTotal float64
    From       time.Time
    To         time.Time
}
