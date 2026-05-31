package ledger

import (
	"testing"
	"time"

	"github.com/joshuataye/ledgerpipeline/internal/parser"
)

func TestJournalFromTransactions(t *testing.T) {
	j := NewJournal()
	d := time.Date(2026, 1, 5, 0, 0, 0, 0, time.UTC)
	j.FromTransactions([]parser.Transaction{
		{Date: d, Description: "Groceries", Category: "Food", Amount: -40},
	}, "Cash")
	entries := j.Entries()
	if len(entries) != 2 {
		t.Fatalf("Entries() len = %d", len(entries))
	}
}
