package release

import "time"

// EscrowSummary holds aggregated metrics for escrow/release analysis.
type EscrowSummary struct {
    Count      int
    Net        float64
    DebitTotal float64
    CreditTotal float64
    From       time.Time
    To         time.Time
}
