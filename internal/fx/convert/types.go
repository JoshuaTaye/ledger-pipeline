package convert

import "time"

// FxSummary holds aggregated metrics for fx/convert analysis.
type FxSummary struct {
    Count      int
    Net        float64
    DebitTotal float64
    CreditTotal float64
    From       time.Time
    To         time.Time
}
