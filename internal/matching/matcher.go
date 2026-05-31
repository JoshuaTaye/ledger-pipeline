package matching

import (
	"math"
	"strings"
	"time"

	"github.com/joshuataye/ledgerpipeline/internal/parser"
)

// Options controls transfer pair matching.
type Options struct {
	MaxDays   int
	Tolerance float64
}

// DefaultOptions returns sensible matching defaults.
func DefaultOptions() Options {
	return Options{MaxDays: 3, Tolerance: 0.01}
}

// Pair links an outgoing debit with an incoming credit transfer.
type Pair struct {
	From parser.Transaction
	To   parser.Transaction
}

// FindTransfers pairs opposite-sign transactions with matching amounts.
func FindTransfers(txns []parser.Transaction, opt Options) []Pair {
	if opt.MaxDays <= 0 {
		opt.MaxDays = DefaultOptions().MaxDays
	}
	if opt.Tolerance <= 0 {
		opt.Tolerance = DefaultOptions().Tolerance
	}
	var pairs []Pair
	used := make([]bool, len(txns))
	for i := range txns {
		if used[i] || txns[i].Amount >= 0 {
			continue
		}
		for j := range txns {
			if used[j] || i == j || txns[j].Amount <= 0 {
				continue
			}
			if !amountsMatch(txns[i].Amount, txns[j].Amount, opt.Tolerance) {
				continue
			}
			days := txns[i].Date.Sub(txns[j].Date)
			if days < 0 {
				days = -days
			}
			if days > time.Duration(opt.MaxDays)*24*time.Hour {
				continue
			}
			if !descriptionsRelated(txns[i].Description, txns[j].Description) {
				continue
			}
			pairs = append(pairs, Pair{From: txns[i], To: txns[j]})
			used[i], used[j] = true, true
			break
		}
	}
	return pairs
}

// Unmatched returns transactions not included in any pair.
func Unmatched(txns []parser.Transaction, pairs []Pair) []parser.Transaction {
	used := make([]bool, len(txns))
	for _, p := range pairs {
		for i, tx := range txns {
			if used[i] {
				continue
			}
			if sameTransaction(tx, p.From) || sameTransaction(tx, p.To) {
				used[i] = true
			}
		}
	}
	var out []parser.Transaction
	for i, tx := range txns {
		if !used[i] {
			out = append(out, tx)
		}
	}
	return out
}

func sameTransaction(a, b parser.Transaction) bool {
	return a.Date.Equal(b.Date) && a.Description == b.Description &&
		a.Category == b.Category && a.Amount == b.Amount
}

func amountsMatch(a, b, tol float64) bool {
	return math.Abs(a+b) <= tol
}

func descriptionsRelated(a, b string) bool {
	a = strings.ToLower(strings.TrimSpace(a))
	b = strings.ToLower(strings.TrimSpace(b))
	if a == "" || b == "" {
		return true
	}
	return strings.Contains(a, "transfer") || strings.Contains(b, "transfer") ||
		strings.Contains(a, b) || strings.Contains(b, a)
}
