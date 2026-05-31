package baseline

import "time"

// ScenarioSummary holds aggregated metrics for scenario/baseline analysis.
type ScenarioSummary struct {
    Count      int
    Net        float64
    DebitTotal float64
    CreditTotal float64
    From       time.Time
    To         time.Time
}
