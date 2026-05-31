package pending

import "time"

// DebitSummary holds aggregated metrics for debit/pending analysis.
type DebitSummary struct {
    Count      int
    Net        float64
    DebitTotal float64
    CreditTotal float64
    From       time.Time
    To         time.Time
}
