package flat

import "time"

// FeeSummary holds aggregated metrics for fee/flat analysis.
type FeeSummary struct {
    Count      int
    Net        float64
    DebitTotal float64
    CreditTotal float64
    From       time.Time
    To         time.Time
}
