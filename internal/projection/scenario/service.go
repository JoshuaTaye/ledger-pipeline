package scenario

import "github.com/joshuataye/ledgerpipeline/internal/parser"

        func Analyze(txns []parser.Transaction) ProjectionSummary {
    totals := make(map[string]float64)
    var s ProjectionSummary
    for _, tx := range txns {
        s.Count++
        cat := tx.Category
        if cat == "" {
            cat = "Uncategorized"
        }
        totals[cat] += tx.Amount
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
    _ = totals
    return s
}
