package parser

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

const dateLayout = "2006-01-02"

// ParseFile reads and parses a CSV bank statement from path.
func ParseFile(path string) ([]Transaction, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open statement: %w", err)
	}
	defer f.Close()
	return ParseReader(f)
}

// ParseReader parses CSV bank statement rows from r.
// Expected header: date,description,category,amount
func ParseReader(r io.Reader) ([]Transaction, error) {
	reader := csv.NewReader(r)
	reader.TrimLeadingSpace = true

	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("read csv: %w", err)
	}
	if len(records) == 0 {
		return nil, nil
	}

	start := 0
	if isHeader(records[0]) {
		start = 1
	}

	out := make([]Transaction, 0, len(records)-start)
	for i := start; i < len(records); i++ {
		row := records[i]
		if len(row) == 0 || strings.TrimSpace(strings.Join(row, "")) == "" {
			continue
		}
		if len(row) < 4 {
			return nil, fmt.Errorf("row %d: expected 4 columns, got %d", i+1, len(row))
		}

		date, err := time.Parse(dateLayout, strings.TrimSpace(row[0]))
		if err != nil {
			return nil, fmt.Errorf("row %d: invalid date %q: %w", i+1, row[0], err)
		}

		amount, err := strconv.ParseFloat(strings.TrimSpace(row[3]), 64)
		if err != nil {
			return nil, fmt.Errorf("row %d: invalid amount %q: %w", i+1, row[3], err)
		}

		out = append(out, Transaction{
			Date:        date,
			Description: strings.TrimSpace(row[1]),
			Category:    strings.TrimSpace(row[2]),
			Amount:      amount,
		})
	}

	return out, nil
}

func isHeader(row []string) bool {
	if len(row) < 4 {
		return false
	}
	return strings.EqualFold(strings.TrimSpace(row[0]), "date") &&
		strings.EqualFold(strings.TrimSpace(row[1]), "description") &&
		strings.EqualFold(strings.TrimSpace(row[2]), "category") &&
		strings.EqualFold(strings.TrimSpace(row[3]), "amount")
}

// ParseCSV is an alias for ParseReader.
func ParseCSV(r io.Reader) ([]Transaction, error) {
	return ParseReader(r)
}
