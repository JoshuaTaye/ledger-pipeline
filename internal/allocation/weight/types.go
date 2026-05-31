package weight

import "time"

// AllocationSummary holds aggregated metrics for allocation/weight analysis.
type AllocationSummary struct {
    Count      int
    Net        float64
    DebitTotal float64
    CreditTotal float64
    From       time.Time
    To         time.Time
}
