package rule

import "time"

// NotifySummary holds aggregated metrics for notify/rule analysis.
type NotifySummary struct {
    Count      int
    Net        float64
    DebitTotal float64
    CreditTotal float64
    From       time.Time
    To         time.Time
}
