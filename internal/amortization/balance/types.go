package balance

import "time"

// AmortizationSummary holds aggregated metrics for amortization/balance analysis.
type AmortizationSummary struct {
    Count      int
    Net        float64
    DebitTotal float64
    CreditTotal float64
    From       time.Time
    To         time.Time
}
