package fuzzy

import "time"

// MatchSummary holds aggregated metrics for match/fuzzy analysis.
type MatchSummary struct {
    Count      int
    Net        float64
    DebitTotal float64
    CreditTotal float64
    From       time.Time
    To         time.Time
}
