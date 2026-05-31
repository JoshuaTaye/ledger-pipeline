package pipeline

import (
	"testing"
	"time"

	"github.com/joshuataye/ledgerpipeline/internal/parser"
	"github.com/joshuataye/ledgerpipeline/internal/validate"
)

func TestRunWithAudit(t *testing.T) {
	d := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	txns := []parser.Transaction{
		{Date: d, Description: "Coffee", Category: "Food", Amount: -5},
	}
	res, err := RunWithAudit(txns, Config{Validate: validate.Validator{}})
	if err != nil || len(res.Log.Entries()) < 2 {
		t.Fatalf("RunWithAudit() err=%v entries=%d", err, len(res.Log.Entries()))
	}
}

func TestRunWithAuditRequiresDescription(t *testing.T) {
	d := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	txns := []parser.Transaction{{Date: d, Amount: -5}}
	_, err := RunWithAudit(txns, Config{})
	if err == nil {
		t.Fatal("expected error for missing description")
	}
}
