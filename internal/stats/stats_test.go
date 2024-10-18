package stats

import (
	"testing"

	"github.com/joshuataye/ledgerpipeline/internal/parser"
)

func TestCompute(t *testing.T) {
	s := Compute([]parser.Transaction{{Amount: -50}, {Amount: 100}})
	if s.Count != 2 || s.DebitSum != -50 {
		t.Fatalf("%+v", s)
	}
}
