package schedule

import "time"

// AmortizationSummary holds aggregated metrics for amortization/schedule analysis.
type AmortizationSummary struct {
    Count      int
    Net        float64
    DebitTotal float64
    CreditTotal float64
    From       time.Time
    To         time.Time
}
