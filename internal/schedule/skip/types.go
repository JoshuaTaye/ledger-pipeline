package skip

import "time"

// ScheduleSummary holds aggregated metrics for schedule/skip analysis.
type ScheduleSummary struct {
    Count      int
    Net        float64
    DebitTotal float64
    CreditTotal float64
    From       time.Time
    To         time.Time
}
