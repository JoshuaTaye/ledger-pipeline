package target

import "time"

// GoalSummary holds aggregated metrics for goal/target analysis.
type GoalSummary struct {
    Count      int
    Net        float64
    DebitTotal float64
    CreditTotal float64
    From       time.Time
    To         time.Time
}
