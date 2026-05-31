package dedupe

import (
	"fmt"

	"github.com/joshuataye/ledgerpipeline/internal/parser"
)

func fingerprint(tx parser.Transaction) string {
	return fmt.Sprintf("%s|%s|%s", tx.Date.Format("2006-01-02"), tx.Description, tx.Category)
}

// RemoveDuplicates keeps first occurrence of identical rows.
func RemoveDuplicates(txns []parser.Transaction) []parser.Transaction {
	seen := make(map[string]struct{})
	out := make([]parser.Transaction, 0, len(txns))
	for _, tx := range txns {
		key := fingerprint(tx)
		if _, ok := seen[key]; ok {
			continue
		}
		seen[key] = struct{}{}
		out = append(out, tx)
	}
	return out
}
