package reserve

import "time"

// VaultSummary holds aggregated metrics for vault/reserve analysis.
type VaultSummary struct {
    Count      int
    Net        float64
    DebitTotal float64
    CreditTotal float64
    From       time.Time
    To         time.Time
}
