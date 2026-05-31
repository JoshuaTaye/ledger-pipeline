package internal

import "time"

// TransferSummary holds aggregated metrics for transfer/internal analysis.
type TransferSummary struct {
    Count      int
    Net        float64
    DebitTotal float64
    CreditTotal float64
    From       time.Time
    To         time.Time
}
