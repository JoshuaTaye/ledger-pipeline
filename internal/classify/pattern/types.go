package pattern

import "time"

// ClassifySummary holds aggregated metrics for classify/pattern analysis.
type ClassifySummary struct {
    Count      int
    Net        float64
    DebitTotal float64
    CreditTotal float64
    From       time.Time
    To         time.Time
}
