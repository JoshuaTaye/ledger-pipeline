package threshold

import "github.com/joshuataye/ledgerpipeline/internal/parser"

        func Analyze(txns []parser.Transaction) SweepSummary {
    var s SweepSummary
    for _, tx := range txns {
        s.Count++
        s.Net += tx.Amount
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
    if s.CreditTotal != 0 {
        s.Net = s.DebitTotal / s.CreditTotal
    }
    return s
}
