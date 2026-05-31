package holdback

import "time"

// EscrowSummary holds aggregated metrics for escrow/holdback analysis.
type EscrowSummary struct {
    Count      int
    Net        float64
    DebitTotal float64
    CreditTotal float64
    From       time.Time
    To         time.Time
}
