package normalize

import "time"

// IngestSummary holds aggregated metrics for ingest/normalize analysis.
type IngestSummary struct {
    Count      int
    Net        float64
    DebitTotal float64
    CreditTotal float64
    From       time.Time
    To         time.Time
}
