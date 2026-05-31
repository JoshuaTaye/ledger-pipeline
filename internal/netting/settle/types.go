package settle

import "time"

// NettingSummary holds aggregated metrics for netting/settle analysis.
type NettingSummary struct {
    Count      int
    Net        float64
    DebitTotal float64
    CreditTotal float64
    From       time.Time
    To         time.Time
}
