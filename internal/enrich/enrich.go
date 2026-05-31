package enrich

import (
	"strings"

	"github.com/joshuataye/ledgerpipeline/internal/merchant"
	"github.com/joshuataye/ledgerpipeline/internal/parser"
)

// Options controls enrichment behavior.
type Options struct {
	Normalize bool
	Suffixes  map[string]string
}

// Apply enriches transaction descriptions and categories.
func Apply(txns []parser.Transaction, opt Options) []parser.Transaction {
	out := make([]parser.Transaction, len(txns))
	for i, tx := range txns {
		out[i] = tx
		if opt.Normalize {
			out[i].Description = merchant.Normalize(tx.Description)
		}
		for suffix, category := range opt.Suffixes {
			desc := strings.ToLower(out[i].Description)
			if strings.HasSuffix(desc, suffix) && out[i].Category == "" {
				out[i].Category = category
			}
		}
	}
	return out
}
