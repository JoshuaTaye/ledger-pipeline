package reinvest

import "time"

// DividendSummary holds aggregated metrics for dividend/reinvest analysis.
type DividendSummary struct {
    Count      int
    Net        float64
    DebitTotal float64
    CreditTotal float64
    From       time.Time
    To         time.Time
}
