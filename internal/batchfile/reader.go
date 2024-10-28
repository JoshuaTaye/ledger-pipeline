package batchfile

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/joshuataye/ledgerpipeline/internal/parser"
)

// ReadDir parses every CSV file in a directory.
func ReadDir(dir string) ([]parser.Transaction, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	var all []parser.Transaction
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(strings.ToLower(e.Name()), ".csv") {
			continue
		}
		path := filepath.Join(dir, e.Name())
		txns, err := parser.ParseFile(path)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", e.Name(), err)
		}
		all = append(all, txns...)
	}
	return all, nil
}
