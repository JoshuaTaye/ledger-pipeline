package report

import (
	"strings"
	"testing"

	"github.com/joshuataye/ledgerpipeline/internal/aggregate"
)

func TestWriteSummary_IncludesHeaderAndRows(t *testing.T) {
	summaries := []aggregate.CategorySummary{
		{Category: "Food", TransactionCount: 2, TotalAmount: -30, DebitTotal: -30},
		{Category: "Income", TransactionCount: 1, TotalAmount: 2000, CreditTotal: 2000},
	}

	var buf strings.Builder
	if err := WriteSummary(&buf, summaries, 1970); err != nil {
		t.Fatalf("WriteSummary() error = %v", err)
	}

	out := buf.String()
	checks := []string{
		"Bank Statement Summary",
		"Category",
		"Food",
		"Income",
		"TOTAL",
		"1970.00",
	}
	for _, want := range checks {
		if !strings.Contains(out, want) {
			t.Fatalf("output missing %q:\n%s", want, out)
		}
	}
}

func TestWriteSummary_Empty(t *testing.T) {
	var buf strings.Builder
	if err := WriteSummary(&buf, nil, 0); err != nil {
		t.Fatalf("WriteSummary() error = %v", err)
	}
	if !strings.Contains(buf.String(), "No transactions found.") {
		t.Fatalf("unexpected empty output: %s", buf.String())
	}
}

func TestWriteSummary_TotalsRow(t *testing.T) {
	summaries := []aggregate.CategorySummary{
		{Category: "Utilities", TransactionCount: 1, TotalAmount: -80, DebitTotal: -80},
		{Category: "Salary", TransactionCount: 1, TotalAmount: 3000, CreditTotal: 3000},
	}

	var buf strings.Builder
	net := float64(2920)
	if err := WriteSummary(&buf, summaries, net); err != nil {
		t.Fatalf("WriteSummary() error = %v", err)
	}

	out := buf.String()
	if !strings.Contains(out, "TOTAL") {
		t.Fatal("missing TOTAL row")
	}
	if !strings.Contains(out, "2920.00") {
		t.Fatalf("missing net total in output:\n%s", out)
	}
}

func TestWriteSummary_FormatsCurrency(t *testing.T) {
	summaries := []aggregate.CategorySummary{
		{Category: "Food", TransactionCount: 1, TotalAmount: -4.5, DebitTotal: -4.5},
	}

	var buf strings.Builder
	if err := WriteSummary(&buf, summaries, -4.5); err != nil {
		t.Fatalf("WriteSummary() error = %v", err)
	}

	if !strings.Contains(buf.String(), "-4.50") {
		t.Fatalf("expected formatted amount, got:\n%s", buf.String())
	}
}
