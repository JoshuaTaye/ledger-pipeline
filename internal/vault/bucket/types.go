package bucket

import "time"

// VaultSummary holds aggregated metrics for vault/bucket analysis.
type VaultSummary struct {
    Count      int
    Net        float64
    DebitTotal float64
    CreditTotal float64
    From       time.Time
    To         time.Time
}
