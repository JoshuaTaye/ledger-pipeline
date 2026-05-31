package enrich

import (
	"testing"
	"time"

	"github.com/joshuataye/ledgerpipeline/internal/parser"
)

func TestApply(t *testing.T) {
	d := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	txns := []parser.Transaction{
		{Date: d, Description: "AMZN MARKETPLACE", Amount: -20},
	}
	out := Apply(txns, Options{Normalize: true})
	if out[0].Description != "amazon" {
		t.Fatalf("Apply() desc = %q", out[0].Description)
	}
}
