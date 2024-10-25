package pipeline

import (
	"testing"
	"time"

	"github.com/joshuataye/ledgerpipeline/internal/categorize/rules"
	"github.com/joshuataye/ledgerpipeline/internal/parser"
	"github.com/joshuataye/ledgerpipeline/internal/validate"
)

func TestRun_Dedupe(t *testing.T) {
	d := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	txns := []parser.Transaction{
		{Date: d, Description: "Coffee", Category: "Food", Amount: -3},
		{Date: d, Description: "Coffee", Category: "Food", Amount: -3},
	}
	res, err := Run(txns, Config{Dedupe: true, Validate: validate.Validator{}})
	if err != nil || len(res.Transactions) != 1 {
		t.Fatalf("err %v len %d", err, len(res.Transactions))
	}
}

func TestRun_CategorizeRules(t *testing.T) {
	d := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	txns := []parser.Transaction{
		{Date: d, Description: "UBER TRIP", Amount: -12},
	}
	res, err := Run(txns, Config{
		CategorizeRules: []rules.Rule{{Category: "Transport", Contains: "UBER", Priority: 1}},
	})
	if err != nil || res.Transactions[0].Category != "Transport" {
		t.Fatalf("err %v category %q", err, res.Transactions[0].Category)
	}
}
