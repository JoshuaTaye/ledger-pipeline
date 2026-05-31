package gain

import "time"

// InvestmentSummary holds aggregated metrics for investment/gain analysis.
type InvestmentSummary struct {
    Count      int
    Net        float64
    DebitTotal float64
    CreditTotal float64
    From       time.Time
    To         time.Time
}
