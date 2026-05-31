package compound

import "time"

// InterestSummary holds aggregated metrics for interest/compound analysis.
type InterestSummary struct {
    Count      int
    Net        float64
    DebitTotal float64
    CreditTotal float64
    From       time.Time
    To         time.Time
}
