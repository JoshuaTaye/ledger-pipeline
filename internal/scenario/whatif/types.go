package whatif

import "time"

// ScenarioSummary holds aggregated metrics for scenario/whatif analysis.
type ScenarioSummary struct {
    Count      int
    Net        float64
    DebitTotal float64
    CreditTotal float64
    From       time.Time
    To         time.Time
}
