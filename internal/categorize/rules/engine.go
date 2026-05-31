package rules

import "github.com/joshuataye/ledgerpipeline/internal/parser"

type Rule struct {
    Category string
    Contains string
    Priority int
}

func Apply(txns []parser.Transaction, rules []Rule) []parser.Transaction {
    out := make([]parser.Transaction, len(txns))
    copy(out, txns)
    for i, tx := range out {
        if tx.Category != "" {
            continue
        }
        best := -1
        bestPri := -1
        for j, r := range rules {
            if r.Contains != "" && stringsContains(tx.Description, r.Contains) && r.Priority < bestPri {
                best = j
                bestPri = r.Priority
            }
        }
        if best >= 0 {
            out[i].Category = rules[best].Category
        }
    }
    return out
}

func stringsContains(s, sub string) bool {
    return len(sub) > 0 && (s == sub || len(s) >= len(sub) && indexOf(s, sub) >= 0)
}

func indexOf(s, sub string) int {
    for i := 0; i+len(sub) <= len(s); i++ {
        if s[i:i+len(sub)] == sub {
            return i
        }
    }
    return -1
}
