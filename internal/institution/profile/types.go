package profile

import "time"

// InstitutionSummary holds aggregated metrics for institution/profile analysis.
type InstitutionSummary struct {
    Count      int
    Net        float64
    DebitTotal float64
    CreditTotal float64
    From       time.Time
    To         time.Time
}
