package qif

import (
	"fmt"
	"os"

	"github.com/joshuataye/ledgerpipeline/internal/parser"
)

// ParseFile reads a QIF file from path.
func ParseFile(path string) ([]parser.Transaction, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open qif: %w", err)
	}
	defer f.Close()
	txns, err := ParseReader(f)
	if err != nil {
		return nil, err
	}
	return txns, nil
}
