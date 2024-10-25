package export

import (
	"encoding/json"
	"io"

	"github.com/joshuataye/ledgerpipeline/internal/aggregate"
)

func WriteJSON(w io.Writer, summaries []aggregate.CategorySummary, net float64) error {
	payload := struct {
		NetTotal   float64                      `json:"net_total"`
		Categories []aggregate.CategorySummary `json:"categories"`
	}{NetTotal: net, Categories: summaries}
	data, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		return err
	}
	_, err = w.Write(data)
	return err
}
