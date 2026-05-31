package discount

import "time"

// PayableSummary holds aggregated metrics for payable/discount analysis.
type PayableSummary struct {
    Count      int
    Net        float64
    DebitTotal float64
    CreditTotal float64
    From       time.Time
    To         time.Time
}
