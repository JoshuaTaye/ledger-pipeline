package tier

import "time"

// InterestrateSummary holds aggregated metrics for interestrate/tier analysis.
type InterestrateSummary struct {
    Count      int
    Net        float64
    DebitTotal float64
    CreditTotal float64
    From       time.Time
    To         time.Time
}
