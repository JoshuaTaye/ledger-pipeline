package rebalance

import "time"

// PortfolioSummary holds aggregated metrics for portfolio/rebalance analysis.
type PortfolioSummary struct {
    Count      int
    Net        float64
    DebitTotal float64
    CreditTotal float64
    From       time.Time
    To         time.Time
}
