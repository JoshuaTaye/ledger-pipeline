package batch

import "time"

// IngestSummary holds aggregated metrics for ingest/batch analysis.
type IngestSummary struct {
    Count      int
    Net        float64
    DebitTotal float64
    CreditTotal float64
    From       time.Time
    To         time.Time
}
