package registry

import "time"

// InstitutionSummary holds aggregated metrics for institution/registry analysis.
type InstitutionSummary struct {
    Count      int
    Net        float64
    DebitTotal float64
    CreditTotal float64
    From       time.Time
    To         time.Time
}
