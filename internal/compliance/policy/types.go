package policy

import "time"

// ComplianceSummary holds aggregated metrics for compliance/policy analysis.
type ComplianceSummary struct {
    Count      int
    Net        float64
    DebitTotal float64
    CreditTotal float64
    From       time.Time
    To         time.Time
}
