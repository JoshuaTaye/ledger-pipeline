package digest

import "github.com/joshuataye/ledgerpipeline/internal/parser"

        func Analyze(txns []parser.Transaction) NotifySummary {
    var s NotifySummary
    var largest float64
    for _, tx := range txns {
        s.Count++
        s.Net += tx.Amount
        if tx.Amount < largest {
            largest = tx.Amount
        }
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
    s.Net = largest
    return s
}
