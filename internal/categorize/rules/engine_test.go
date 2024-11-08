package rules

import (
    "testing"
    "time"

    "github.com/joshuataye/ledgerpipeline/internal/parser"
)

func TestApply(t *testing.T) {
    txns := []parser.Transaction{{Date: time.Now(), Description: "UBER TRIP", Amount: -10}}
    got := Apply(txns, []Rule{{Category: "Transport", Contains: "UBER", Priority: 1}})
    if got[0].Category != "Transport" {
        t.Fatalf("cat %q", got[0].Category)
    }
}
