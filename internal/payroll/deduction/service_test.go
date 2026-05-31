package deduction

import (
    "testing"
    "time"

    "github.com/joshuataye/ledgerpipeline/internal/parser"
)

func TestAnalyze_payroll_deduction(t *testing.T) {
    txns := []parser.Transaction{
        {Date: time.Date(2026, 1, 5, 0, 0, 0, 0, time.UTC), Amount: -50, Category: "Food"},
        {Date: time.Date(2026, 1, 10, 0, 0, 0, 0, time.UTC), Amount: 200, Category: "Payroll"},
        {Date: time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC), Amount: -25, Category: "Food"},
    }
    got := Analyze(txns)
    if got.From.IsZero() || got.To.IsZero() {
        t.Fatal("expected date span")
    }
    if got.To.Before(got.From) {
        t.Fatal("invalid date span")
    }
}

func TestAnalyze_payroll_deduction_Empty(t *testing.T) {
    got := Analyze(nil)
    if got.Count != 0 {
        t.Fatalf("count=%d want 0", got.Count)
    }
    if !got.From.IsZero() || !got.To.IsZero() {
        t.Fatal("expected zero dates for empty input")
    }
}
