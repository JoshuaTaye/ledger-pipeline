package calendar

import "github.com/joshuataye/ledgerpipeline/internal/parser"

        func Analyze(txns []parser.Transaction) ScheduleSummary {
    const threshold = 100.0
    var s ScheduleSummary
    for _, tx := range txns {
        if tx.Amount < -threshold || tx.Amount > threshold {
            s.Count++
        }
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
    return s
}
