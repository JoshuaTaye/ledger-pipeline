package stats

import "github.com/joshuataye/ledgerpipeline/internal/parser"

type Summary struct {
	Count      int
	DebitSum   float64
	CreditSum  float64
	LargestDebit float64
}

func Compute(txns []parser.Transaction) Summary {
	var s Summary
	for _, tx := range txns {
		s.Count++
		if tx.Amount < 0 {
			s.DebitSum += tx.Amount
			if tx.Amount > s.LargestDebit {
				s.LargestDebit = tx.Amount
			}
		} else {
			s.CreditSum += tx.Amount
		}
	}
	return s
}
