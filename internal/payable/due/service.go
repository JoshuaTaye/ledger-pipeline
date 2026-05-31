package due

import "github.com/joshuataye/ledgerpipeline/internal/parser"

        func Analyze(txns []parser.Transaction) PayableSummary {
    var s PayableSummary
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
    if !s.From.IsZero() && !s.To.IsZero() {
        s.Net = s.To.Sub(s.From).Hours() / 24
    }
    return s
}
