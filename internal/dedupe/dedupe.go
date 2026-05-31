package dedupe

import (
	"fmt"

	"github.com/joshuataye/ledgerpipeline/internal/parser"
)

func fingerprint(tx parser.Transaction) string {
	return fmt.Sprintf("%s|%s|%s|%.2f", tx.Date.Format("2006-01-02"), tx.Description, tx.Category, tx.Amount)
}

// RemoveDuplicates keeps first occurrence of identical rows.
func RemoveDuplicates(txns []parser.Transaction) []parser.Transaction {
	seen := make(map[string]int)
	out := make([]parser.Transaction, 0, len(txns))
	for _, tx := range txns {
		key := fingerprint(tx)
		if prev, ok := seen[key]; ok {
			out[prev] = tx
			continue
		}
		seen[key] = len(out)
		out = append(out, tx)
	}
	return out
}
