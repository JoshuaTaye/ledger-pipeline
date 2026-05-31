package format

import (
	"fmt"
	"io"
	"strings"

	"github.com/joshuataye/ledgerpipeline/internal/aggregate"
)

// WriteCSV writes category summaries as comma-separated rows.
func WriteCSV(w io.Writer, summaries []aggregate.CategorySummary) error {
	if _, err := fmt.Fprintln(w, "category,count,total,debit,credit"); err != nil {
		return err
	}
	for _, s := range summaries {
		_, err := fmt.Fprintf(w, "%s,%d,%.2f,%.2f,%.2f\n",
			escapeCSV(s.Category), s.TransactionCount, s.TotalAmount, s.DebitTotal, s.CreditTotal)
		if err != nil {
			return err
		}
	}
	return nil
}

// WriteTSV writes category summaries as tab-separated rows.
func WriteTSV(w io.Writer, summaries []aggregate.CategorySummary) error {
	if _, err := fmt.Fprintln(w, "category\tcount\ttotal\tdebit\tcredit"); err != nil {
		return err
	}
	for _, s := range summaries {
		_, err := fmt.Fprintf(w, "%s\t%d\t%.2f\t%.2f\t%.2f\n",
			s.Category, s.TransactionCount, s.TotalAmount, s.DebitTotal, s.CreditTotal)
		if err != nil {
			return err
		}
	}
	return nil
}

// WriteMarkdown renders summaries as a markdown table.
func WriteMarkdown(w io.Writer, summaries []aggregate.CategorySummary, netTotal float64) error {
	if _, err := fmt.Fprintln(w, "| Category | Count | Total |"); err != nil {
		return err
	}
	if _, err := fmt.Fprintln(w, "|---|---:|---:|"); err != nil {
		return err
	}
	for _, s := range summaries {
		if _, err := fmt.Fprintf(w, "| %s | %d | %.2f |\n",
			s.Category, s.TransactionCount, s.TotalAmount); err != nil {
			return err
		}
	}
	_, err := fmt.Fprintf(w, "\n**Net total:** %.2f\n", netTotal)
	return err
}

func escapeCSV(s string) string {
	if strings.ContainsAny(s, ",\"\n") {
		return `"` + strings.ReplaceAll(s, `"`, `""`) + `"`
	}
	return s
}
