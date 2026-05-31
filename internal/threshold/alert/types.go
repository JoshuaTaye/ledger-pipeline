package alert

import "time"

// ThresholdSummary holds aggregated metrics for threshold/alert analysis.
type ThresholdSummary struct {
    Count      int
    Net        float64
    DebitTotal float64
    CreditTotal float64
    From       time.Time
    To         time.Time
}
