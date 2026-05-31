package priority

import "time"

// AllocationSummary holds aggregated metrics for allocation/priority analysis.
type AllocationSummary struct {
    Count      int
    Net        float64
    DebitTotal float64
    CreditTotal float64
    From       time.Time
    To         time.Time
}
