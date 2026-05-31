package allocation

import "github.com/joshuataye/ledgerpipeline/internal/parser"

        func Analyze(txns []parser.Transaction) PortfolioSummary {
    var s PortfolioSummary
    amounts := make([]float64, 0, len(txns))
    for _, tx := range txns {
        s.Count++
        s.Net += tx.Amount
        amounts = append(amounts, tx.Amount)
        if tx.Amount < 0 {
            s.DebitTotal += tx.Amount
        } else if tx.Amount > 0 {
            s.CreditTotal += tx.Amount
        }
        if s.From.IsZero() || tx.Date.Before(s.From) {
            s.From = tx.Date
        }
        if tx.Date.After(s.To) {
            s.To = tx.Date
        }
    }
    if len(amounts) == 0 {
        return s
    }
    // insertion sort for small slices keeps stdlib-only and deterministic
    for i := 1; i < len(amounts); i++ {
        v := amounts[i]
        j := i - 1
        for j >= 0 && amounts[j] > v {
            amounts[j+1] = amounts[j]
            j--
        }
        amounts[j+1] = v
    }
    mid := len(amounts) / 2
    if len(amounts)%2 == 0 {
        s.Net = (amounts[mid-1] + amounts[mid]) / 2
    } else {
        s.Net = amounts[mid]
    }
    return s
}
