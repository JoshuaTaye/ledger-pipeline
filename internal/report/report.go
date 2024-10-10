package report

import (
	"fmt"
	"io"
	"strings"

	"github.com/joshuataye/ledgerpipeline/internal/aggregate"
)

// WriteSummary writes a human-readable category summary to w.
func WriteSummary(w io.Writer, summaries []aggregate.CategorySummary, netTotal float64) error {
	var b strings.Builder

	b.WriteString("Bank Statement Summary\n")
	b.WriteString(strings.Repeat("=", 22))
	b.WriteString("\n\n")

	if len(summaries) == 0 {
		b.WriteString("No transactions found.\n")
		_, err := io.WriteString(w, b.String())
		return err
	}

	b.WriteString(fmt.Sprintf("%-20s %8s %12s %12s %12s\n",
		"Category", "Count", "Debits", "Credits", "Net"))
	b.WriteString(strings.Repeat("-", 68))
	b.WriteString("\n")

	var totalDebits, totalCredits float64
	var totalCount int

	for _, s := range summaries {
		b.WriteString(fmt.Sprintf("%-20s %8d %12.2f %12.2f %12.2f\n",
			s.Category, s.TransactionCount, s.DebitTotal, s.CreditTotal, s.TotalAmount))
		totalCount += s.TransactionCount
		totalDebits += s.DebitTotal
		totalCredits += s.CreditTotal
	}

	b.WriteString(strings.Repeat("-", 68))
	b.WriteString("\n")
	b.WriteString(fmt.Sprintf("%-20s %8d %12.2f %12.2f %12.2f\n",
		"TOTAL", totalCount, totalDebits, totalCredits, netTotal))
	b.WriteString("\n")

	_, err := io.WriteString(w, b.String())
	return err
}
