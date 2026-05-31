package pair

import "time"

// CurrencySummary holds aggregated metrics for currency/pair analysis.
type CurrencySummary struct {
    Count      int
    Net        float64
    DebitTotal float64
    CreditTotal float64
    From       time.Time
    To         time.Time
}
