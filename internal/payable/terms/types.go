package terms

import "time"

// PayableSummary holds aggregated metrics for payable/terms analysis.
type PayableSummary struct {
    Count      int
    Net        float64
    DebitTotal float64
    CreditTotal float64
    From       time.Time
    To         time.Time
}
