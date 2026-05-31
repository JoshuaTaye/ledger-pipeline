package timeline

import "github.com/joshuataye/ledgerpipeline/internal/parser"

func debitTotal(txns []parser.Transaction) float64 {
    var total float64
    for _, tx := range txns {
        if tx.Amount < 0 {
            total += tx.Amount
        }
    }
    return total
}

func creditTotal(txns []parser.Transaction) float64 {
    var total float64
    for _, tx := range txns {
        if tx.Amount > 0 {
            total += tx.Amount
        }
    }
    return total
}
