package validate

import (
	"testing"
	"time"

	"github.com/joshuataye/ledgerpipeline/internal/parser"
)

func TestValidateOne_RejectsEmptyDescription(t *testing.T) {
	v := Validator{}
	err := v.ValidateOne(parser.Transaction{Date: time.Now(), Amount: -1})
	if err == nil {
		t.Fatal("expected error")
	}
}
