package matching

import "time"

// InvoiceSummary holds aggregated metrics for invoice/matching analysis.
type InvoiceSummary struct {
    Count      int
    Net        float64
    DebitTotal float64
    CreditTotal float64
    From       time.Time
    To         time.Time
}
