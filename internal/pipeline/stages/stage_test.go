package stages

import (
    "testing"
    "time"

    "github.com/joshuataye/ledgerpipeline/internal/parser"
    "github.com/joshuataye/ledgerpipeline/internal/validate"
)

func TestRegistry(t *testing.T) {
    d := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
            txns := []parser.Transaction{
                {Date: d, Description: "A", Category: "Food", Amount: -1},
                {Date: d, Description: "A", Category: "Food", Amount: -1},
            }
    reg := NewRegistry(DedupeStage{}, ValidateStage{V: validate.Validator{}})
    out, err := reg.Run(txns)
    if err != nil || len(out) != 1 {
        t.Fatalf("err=%v len=%d", err, len(out))
    }
}
