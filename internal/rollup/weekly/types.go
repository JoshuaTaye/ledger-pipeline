package weekly

import "time"

// RollupSummary holds aggregated metrics for rollup/weekly analysis.
type RollupSummary struct {
    Count      int
    Net        float64
    DebitTotal float64
    CreditTotal float64
    From       time.Time
    To         time.Time
}
