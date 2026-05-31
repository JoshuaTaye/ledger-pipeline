package scenario

import "time"

// ProjectionSummary holds aggregated metrics for projection/scenario analysis.
type ProjectionSummary struct {
    Count      int
    Net        float64
    DebitTotal float64
    CreditTotal float64
    From       time.Time
    To         time.Time
}
