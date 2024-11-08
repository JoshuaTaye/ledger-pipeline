package qif

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/joshuataye/ledgerpipeline/internal/parser"
)

// ParseReader reads Quicken Interchange Format records.
func ParseReader(r io.Reader) ([]parser.Transaction, error) {
	sc := bufio.NewScanner(r)
	var out []parser.Transaction
	var date time.Time
	var desc string
	var amount float64
	flush := func() {
		if desc == "" {
			return
		}
		out = append(out, parser.Transaction{Date: date, Description: desc, Amount: amount})
		desc = ""
		amount = 0
	}
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line == "^" {
			flush()
			continue
		}
		if len(line) < 2 {
			continue
		}
		switch line[0] {
		case 'D':
			t, err := time.Parse("1/2/2006", line[1:])
			if err != nil {
				return nil, fmt.Errorf("date: %w", err)
			}
			date = t
		case 'T':
			v, err := strconv.ParseFloat(strings.TrimPrefix(line, "T"), 64)
			if err != nil {
				return nil, err
			}
			amount = v
		case 'P':
			desc = line[1:]
		}
	}
	flush()
	return out, sc.Err()
}
