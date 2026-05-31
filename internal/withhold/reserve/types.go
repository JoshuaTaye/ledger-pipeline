package reserve

import "time"

// WithholdSummary holds aggregated metrics for withhold/reserve analysis.
type WithholdSummary struct {
    Count      int
    Net        float64
    DebitTotal float64
    CreditTotal float64
    From       time.Time
    To         time.Time
}
