package routing

import "time"

// InstitutionSummary holds aggregated metrics for institution/routing analysis.
type InstitutionSummary struct {
    Count      int
    Net        float64
    DebitTotal float64
    CreditTotal float64
    From       time.Time
    To         time.Time
}
