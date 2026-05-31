package insights

import (
	"fmt"
	"io"
	"sort"

	"github.com/joshuataye/ledgerpipeline/internal/aggregate"
	"github.com/joshuataye/ledgerpipeline/internal/stats"
	"github.com/joshuataye/ledgerpipeline/internal/parser"
)

// Summary combines category rollups with transaction statistics.
type Summary struct {
	TopCategory   string
	TopSpend      float64
	NetTotal      float64
	DebitCount    int
	LargestDebit  float64
}

// Build derives insights from processed transactions.
func Build(txns []parser.Transaction, summaries []aggregate.CategorySummary, netTotal float64) Summary {
	st := stats.Compute(txns)
	var topCat string
	var topSpend float64
	for _, s := range summaries {
		if s.TotalAmount < topSpend {
			topSpend = s.TotalAmount
			topCat = s.Category
		}
	}
	debitCount := 0
	for _, tx := range txns {
		if tx.Amount < 0 {
			debitCount++
		}
	}
	return Summary{
		TopCategory:  topCat,
		TopSpend:     topSpend,
		NetTotal:     netTotal,
		DebitCount:   debitCount,
		LargestDebit: st.LargestDebit,
	}
}

// WriteReport prints a human-readable insights summary.
func WriteReport(w io.Writer, s Summary) error {
	lines := []string{
		fmt.Sprintf("top_category: %s (%.2f)", s.TopCategory, s.TopSpend),
		fmt.Sprintf("net_total: %.2f", s.NetTotal),
		fmt.Sprintf("debit_count: %d", s.DebitCount),
		fmt.Sprintf("largest_debit: %.2f", s.LargestDebit),
	}
	sort.Strings(lines)
	for _, line := range lines {
		if _, err := fmt.Fprintln(w, line); err != nil {
			return err
		}
	}
	return nil
}
