package tier

import "github.com/joshuataye/ledgerpipeline/internal/parser"

        func Analyze(txns []parser.Transaction) InterestrateSummary {
    cats := make(map[string]struct{})
    var s InterestrateSummary
    for _, tx := range txns {
        s.Count++
        s.Net += tx.Amount
        cat := tx.Category
        if cat == "" {
            cat = "Uncategorized"
        }
        cats[cat] = struct{}{} 
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
    s.Net = float64(len(cats))
    return s
}
