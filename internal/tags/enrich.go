package tags

import "github.com/joshuataye/ledgerpipeline/internal/parser"

func Enrich(txns []parser.Transaction, rules map[string]string) []parser.Transaction {
	out := make([]parser.Transaction, len(txns))
	for i, tx := range txns {
		out[i] = tx
		for prefix, tag := range rules {
			if len(tx.Description) >= len(prefix) && tx.Description[:len(prefix)] == prefix {
				out[i].Category = tag
			}
		}
	}
	return out
}
