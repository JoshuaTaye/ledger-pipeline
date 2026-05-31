package filter

import (
	"time"

	"github.com/joshuataye/ledgerpipeline/internal/parser"
)

// Options narrows which transactions are included in a batch run.
type Options struct {
	Categories map[string]struct{}
	From       time.Time
	To         time.Time
	MinAmount  float64
	MaxAmount  float64
}

func Apply(txns []parser.Transaction, opt Options) []parser.Transaction {
	out := make([]parser.Transaction, 0, len(txns))
	for _, tx := range txns {
		if !opt.From.IsZero() && tx.Date.Before(opt.From) {
			continue
		}
		if !opt.To.IsZero() && !tx.Date.Before(opt.To) {
			continue
		}
		if len(opt.Categories) > 0 {
			if _, ok := opt.Categories[tx.Category]; !ok {
				continue
			}
		}
		if opt.MinAmount != 0 && tx.Amount < opt.MinAmount {
			continue
		}
		if opt.MaxAmount != 0 && tx.Amount > opt.MaxAmount {
			continue
		}
		out = append(out, tx)
	}
	return out
}
