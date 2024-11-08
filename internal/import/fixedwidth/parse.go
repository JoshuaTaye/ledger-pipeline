package fixedwidth

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/joshuataye/ledgerpipeline/internal/parser"
)

// Column describes a fixed-width field slice.
type Column struct {
	Start int
	End   int
}

// Layout maps logical fields to columns.
type Layout struct {
	Date        Column
	Description Column
	Amount      Column
}

func ParseLine(line string, layout Layout) (parser.Transaction, error) {
	if layout.Date.End > len(line) || layout.Amount.End > len(line) {
		return parser.Transaction{}, fmt.Errorf("line too short")
	}
	dateStr := strings.TrimSpace(line[layout.Date.Start:layout.Date.End])
	desc := strings.TrimSpace(line[layout.Description.Start:layout.Description.End])
	amtStr := strings.TrimSpace(line[layout.Amount.Start:layout.Amount.End])
	d, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return parser.Transaction{}, err
	}
	amount, err := strconv.ParseFloat(amtStr, 64)
	if err != nil {
		return parser.Transaction{}, err
	}
	return parser.Transaction{Date: d, Description: desc, Amount: amount}, nil
}

func ParseLines(lines []string, layout Layout) ([]parser.Transaction, error) {
	out := make([]parser.Transaction, 0, len(lines))
	for i, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		tx, err := ParseLine(line, layout)
		if err != nil {
			return nil, fmt.Errorf("line %d: %w", i+1, err)
		}
		out = append(out, tx)
	}
	return out, nil
}
