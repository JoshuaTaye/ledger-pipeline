package merge

import "time"

// CategorySummary holds aggregated metrics for category/merge analysis.
type CategorySummary struct {
    Count      int
    Net        float64
    DebitTotal float64
    CreditTotal float64
    From       time.Time
    To         time.Time
}
