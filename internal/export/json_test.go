package export

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/joshuataye/ledgerpipeline/internal/aggregate"
)

func TestWriteJSON(t *testing.T) {
	var buf bytes.Buffer
	err := WriteJSON(&buf, []aggregate.CategorySummary{{Category: "Food", TotalAmount: -5}}, -5)
	if err != nil {
		t.Fatal(err)
	}
	var m map[string]any
	if json.Unmarshal(buf.Bytes(), &m) != nil {
		t.Fatal("invalid json")
	}
}
