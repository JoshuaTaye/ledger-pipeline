package queue

import "time"

// ApprovalSummary holds aggregated metrics for approval/queue analysis.
type ApprovalSummary struct {
    Count      int
    Net        float64
    DebitTotal float64
    CreditTotal float64
    From       time.Time
    To         time.Time
}
