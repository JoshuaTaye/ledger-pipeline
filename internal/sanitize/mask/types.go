package mask

import "time"

// SanitizeSummary holds aggregated metrics for sanitize/mask analysis.
type SanitizeSummary struct {
    Count      int
    Net        float64
    DebitTotal float64
    CreditTotal float64
    From       time.Time
    To         time.Time
}
