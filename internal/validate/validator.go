package validate

import (
	"fmt"
	"strings"
	"time"

	"github.com/joshuataye/ledgerpipeline/internal/parser"
)

// Validator applies row-level checks after parsing.
type Validator struct {
	MinDate time.Time
	MaxDate time.Time
}

func (v Validator) ValidateAll(txns []parser.Transaction) error {
	for i, tx := range txns {
		if err := v.ValidateOne(tx); err != nil {
			return fmt.Errorf("row %d: %w", i+1, err)
		}
	}
	return nil
}

func (v Validator) ValidateOne(tx parser.Transaction) error {
	if tx.Description == "" {
		return fmt.Errorf("description required")
	}
	if !v.MinDate.IsZero() && tx.Date.Before(v.MinDate) {
		return fmt.Errorf("date before minimum")
	}
	if !v.MaxDate.IsZero() && tx.Date.After(v.MaxDate) {
		return fmt.Errorf("date after maximum")
	}
	if strings.TrimSpace(tx.Category) == "" && tx.Amount < 0 {
		return fmt.Errorf("debit requires category")
	}
	return nil
}
