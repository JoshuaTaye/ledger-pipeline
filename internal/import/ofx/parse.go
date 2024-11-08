package ofx

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/joshuataye/ledgerpipeline/internal/parser"
)

// ParseFile reads an OFX snippet from path.
func ParseFile(path string) ([]parser.Transaction, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read ofx: %w", err)
	}
	return ParseSnippet(string(data))
}

func ParseSnippet(text string) ([]parser.Transaction, error) {
	var out []parser.Transaction
	var date time.Time
	var amount float64
	var desc string
	for _, line := range strings.Split(text, "\n") {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "DTPOSTED:") && len(line) >= 17 {
			d, err := time.Parse("20060102", strings.TrimPrefix(line, "DTPOSTED:")[:8])
			if err != nil {
				return nil, err
			}
			date = d
		}
		if strings.HasPrefix(line, "TRNAMT:") {
			_, _ = fmt.Sscanf(strings.TrimPrefix(line, "TRNAMT:"), "%f", &amount)
		}
		if strings.HasPrefix(line, "NAME:") {
			desc = strings.TrimPrefix(line, "NAME:")
			out = append(out, parser.Transaction{Date: date, Description: desc, Amount: amount})
		}
	}
	return out, nil
}
