package schedule

import "time"

// SweepSummary holds aggregated metrics for sweep/schedule analysis.
type SweepSummary struct {
    Count      int
    Net        float64
    DebitTotal float64
    CreditTotal float64
    From       time.Time
    To         time.Time
}
