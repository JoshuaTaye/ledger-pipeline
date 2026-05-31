package tokenize

import "time"

// SanitizeSummary holds aggregated metrics for sanitize/tokenize analysis.
type SanitizeSummary struct {
    Count      int
    Net        float64
    DebitTotal float64
    CreditTotal float64
    From       time.Time
    To         time.Time
}
