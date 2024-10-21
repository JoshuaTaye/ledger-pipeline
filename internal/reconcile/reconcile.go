package reconcile

import "github.com/joshuataye/ledgerpipeline/internal/parser"

// Result compares running balance to parsed transactions.
type Result struct {
	OpeningBalance float64
	ClosingBalance float64
	ComputedClose  float64
	Delta          float64
}

func Reconcile(opening, closing float64, txns []parser.Transaction) Result {
	var sum float64
	for _, tx := range txns {
		sum += tx.Amount
	}
	computed := opening + sum
	return Result{
		OpeningBalance: opening,
		ClosingBalance: closing,
		ComputedClose:  computed,
		Delta:          closing - computed,
	}
}
