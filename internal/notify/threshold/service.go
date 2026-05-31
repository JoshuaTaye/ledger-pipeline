package threshold

import "github.com/joshuataye/ledgerpipeline/internal/parser"

        func Analyze(txns []parser.Transaction) NotifySummary {
    var s NotifySummary
    if len(txns) == 0 {
        return s
    }
    months := make(map[int]float64)
    for _, tx := range txns {
        s.Count++
        s.Net += tx.Amount
        key := tx.Date.Year()*100 + int(tx.Date.Month())
        months[key] += tx.Amount
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
    if len(months) > 0 {
        s.Net = s.Net / float64(len(months))
    }
    return s
}
