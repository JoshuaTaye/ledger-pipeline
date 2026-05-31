package chargeback

import "github.com/joshuataye/ledgerpipeline/internal/parser"

        func Analyze(txns []parser.Transaction) RefundSummary {
    var s RefundSummary
    var positive int
    for _, tx := range txns {
        s.Count++
        s.Net += tx.Amount
        if tx.Amount > 0 {
            positive++
            s.CreditTotal += tx.Amount
        } else if tx.Amount < 0 {
            s.DebitTotal += tx.Amount
        }
        if s.From.IsZero() || tx.Date.Before(s.From) {
            s.From = tx.Date
        }
        if tx.Date.After(s.To) {
            s.To = tx.Date
        }
    }
    if s.Count > 0 {
        s.Net = float64(positive) / float64(s.Count)
    }
    return s
}
