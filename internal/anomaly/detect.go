package anomaly

import (
	"math"

	"github.com/joshuataye/ledgerpipeline/internal/parser"
)

// Hit flags a transaction that deviates from the batch mean.
type Hit struct {
	Transaction parser.Transaction
	ZScore      float64
}

// Detect finds transactions whose amount z-score exceeds threshold.
func Detect(txns []parser.Transaction, threshold float64) []Hit {
	if len(txns) == 0 || threshold <= 0 {
		return nil
	}
	var sum float64
	for _, tx := range txns {
		sum += tx.Amount
	}
	mean := sum / float64(len(txns))
	var sq float64
	for _, tx := range txns {
		d := tx.Amount - mean
		sq += d * d
	}
	std := math.Sqrt(sq / float64(len(txns)))
	if std == 0 {
		return nil
	}
	var hits []Hit
	for _, tx := range txns {
		z := (tx.Amount - mean) / std
		if math.Abs(z) >= threshold {
			hits = append(hits, Hit{Transaction: tx, ZScore: z})
		}
	}
	return hits
}
