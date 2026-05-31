package weekly

import "time"

// VelocitySummary holds aggregated metrics for velocity/weekly analysis.
type VelocitySummary struct {
    Count      int
    Net        float64
    DebitTotal float64
    CreditTotal float64
    From       time.Time
    To         time.Time
}
