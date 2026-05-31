package split

import (
	"testing"
	"time"

	"github.com/joshuataye/ledgerpipeline/internal/parser"
)

func TestEven(t *testing.T) {
	tx := parser.Transaction{
		Date: time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
		Description: "Split purchase",
		Amount: -90,
	}
	parts := Even(tx, []string{"Food", "Household"})
	if len(parts) != 2 || parts[0].Amount != -45 || parts[1].Amount != -45 {
		t.Fatalf("Even() = %+v", parts)
	}
}

func TestByCategory(t *testing.T) {
	tx := parser.Transaction{Amount: -100}
	parts := ByCategory(tx, map[string]float64{"Food": 3, "Transport": 1})
	if len(parts) != 2 {
		t.Fatalf("ByCategory() len = %d", len(parts))
	}
}
