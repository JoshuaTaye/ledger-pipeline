package retention

import "time"

// ComplianceSummary holds aggregated metrics for compliance/retention analysis.
type ComplianceSummary struct {
    Count      int
    Net        float64
    DebitTotal float64
    CreditTotal float64
    From       time.Time
    To         time.Time
}
