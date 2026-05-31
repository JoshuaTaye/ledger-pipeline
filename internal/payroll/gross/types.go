package gross

import "time"

// PayrollSummary holds aggregated metrics for payroll/gross analysis.
type PayrollSummary struct {
    Count      int
    Net        float64
    DebitTotal float64
    CreditTotal float64
    From       time.Time
    To         time.Time
}
