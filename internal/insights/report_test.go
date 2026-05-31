package insights

import (
	"testing"
	"time"

	"github.com/joshuataye/ledgerpipeline/internal/aggregate"
	"github.com/joshuataye/ledgerpipeline/internal/parser"
)

func TestBuild(t *testing.T) {
	d := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	txns := []parser.Transaction{
		{Date: d, Category: "Food", Amount: -100},
		{Date: d, Category: "Transport", Amount: -20},
	}
	summaries := aggregate.ByCategory(txns)
	s := Build(txns, summaries, aggregate.NetTotal(txns))
	if s.TopCategory != "Food" {
		t.Fatalf("TopCategory = %q", s.TopCategory)
	}
}
