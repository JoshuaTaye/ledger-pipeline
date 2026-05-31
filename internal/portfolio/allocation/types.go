package allocation

import "time"

// PortfolioSummary holds aggregated metrics for portfolio/allocation analysis.
type PortfolioSummary struct {
    Count      int
    Net        float64
    DebitTotal float64
    CreditTotal float64
    From       time.Time
    To         time.Time
}
