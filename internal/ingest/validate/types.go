package validate

import "time"

// IngestSummary holds aggregated metrics for ingest/validate analysis.
type IngestSummary struct {
    Count      int
    Net        float64
    DebitTotal float64
    CreditTotal float64
    From       time.Time
    To         time.Time
}
