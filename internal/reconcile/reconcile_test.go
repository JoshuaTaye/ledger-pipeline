package reconcile

import (
	"testing"

	"github.com/joshuataye/ledgerpipeline/internal/parser"
)

func TestReconcile(t *testing.T) {
	txns := []parser.Transaction{{Amount: -10}, {Amount: 25}}
	r := Reconcile(100, 115, txns)
	if r.Delta != 0 {
		t.Fatalf("delta %v", r.Delta)
	}
}
