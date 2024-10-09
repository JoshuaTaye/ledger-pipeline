package parser

import "time"

// Transaction represents a single bank statement line item.
type Transaction struct {
	Date        time.Time
	Description string
	Category    string
	Amount      float64
}
