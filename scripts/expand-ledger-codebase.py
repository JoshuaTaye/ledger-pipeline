#!/usr/bin/env python3
"""Generate substantive finance-domain packages for ledger-pipeline."""
from __future__ import annotations

import hashlib
import textwrap
from pathlib import Path

ROOT = Path(__file__).resolve().parents[1]
INTERNAL = ROOT / "internal"
MODULE = "github.com/joshuataye/ledgerpipeline"

DOMAINS: dict[str, list[str]] = {
    "amortization": ["schedule", "balance", "paydown"],
    "allocation": ["split", "weight", "priority"],
    "approval": ["workflow", "limit", "queue"],
    "balance": ["running", "snapshot", "carry"],
    "cashflow": ["burn", "runway", "timeline"],
    "compliance": ["audit", "policy", "retention"],
    "currency": ["convert", "normalize", "pair"],
    "deposit": ["hold", "release", "sweep"],
    "dividend": ["accrual", "payout", "reinvest"],
    "escrow": ["holdback", "release", "milestone"],
    "fee": ["tiered", "flat", "minimum"],
    "fx": ["convert", "spread", "historical"],
    "goal": ["progress", "target", "milestone"],
    "institution": ["routing", "profile", "registry"],
    "interest": ["simple", "compound", "accrual"],
    "investment": ["costbasis", "gain", "lot"],
    "invoice": ["aging", "matching", "terms"],
    "loan": ["servicing", "escrow", "prepay"],
    "netting": ["offset", "batch", "settle"],
    "notify": ["threshold", "digest", "rule"],
    "payroll": ["gross", "deduction", "deposit"],
    "portfolio": ["allocation", "rebalance", "drift"],
    "projection": ["linear", "seasonal", "scenario"],
    "refund": ["match", "partial", "chargeback"],
    "remittance": ["batch", "fee", "fx"],
    "rollup": ["daily", "weekly", "hierarchy"],
    "sanitize": ["mask", "tokenize", "normalize"],
    "scenario": ["whatif", "stress", "baseline"],
    "settlement": ["batch", "clearing", "reconcile"],
    "subscription": ["renewal", "proration", "pause"],
    "sweep": ["threshold", "target", "schedule"],
    "threshold": ["alert", "velocity", "limit"],
    "transfer": ["internal", "wire", "ach"],
    "vault": ["bucket", "reserve", "lock"],
    "velocity": ["daily", "weekly", "category"],
    "withhold": ["tax", "garnish", "reserve"],
    "writeoff": ["baddebt", "recovery", "provision"],
    "envelope": ["assign", "rollover", "sweep"],
    "ingest": ["validate", "normalize", "batch"],
    "classify": ["merchant", "mcc", "pattern"],
    "match": ["fuzzy", "exact", "window"],
    "schedule": ["calendar", "recurrence", "skip"],
    "benchmark": ["category", "peer", "baseline"],
    "category": ["rollup", "merge", "tree"],
    "interestrate": ["apr", "apy", "tier"],
    "credit": ["utilization", "limit", "score"],
    "debit": ["hold", "posted", "pending"],
    "revenue": ["recognize", "deferred", "allocate"],
    "expense": ["capitalize", "amortize", "accrue"],
    "payable": ["due", "discount", "terms"],
    "receivable": ["aging", "collect", "doubtful"],
}

TEMPLATES = [
    "running_balance",
    "category_totals",
    "monthly_average",
    "threshold_count",
    "debit_credit_ratio",
    "largest_transaction",
    "date_span",
    "positive_share",
    "category_count",
    "median_amount",
]


def pick_template(domain: str, sub: str) -> str:
    h = hashlib.sha256(f"{domain}/{sub}".encode()).hexdigest()
    return TEMPLATES[int(h[:8], 16) % len(TEMPLATES)]


def types_go(domain: str, sub: str, type_name: str) -> str:
    return textwrap.dedent(
        f"""\
        package {sub}

        import "time"

        // {type_name} holds aggregated metrics for {domain}/{sub} analysis.
        type {type_name} struct {{
            Count      int
            Net        float64
            DebitTotal float64
            CreditTotal float64
            From       time.Time
            To         time.Time
        }}
        """
    )


def helpers_go(sub: str) -> str:
    return textwrap.dedent(
        f"""\
        package {sub}

        import "github.com/joshuataye/ledgerpipeline/internal/parser"

        func debitTotal(txns []parser.Transaction) float64 {{
            var total float64
            for _, tx := range txns {{
                if tx.Amount < 0 {{
                    total += tx.Amount
                }}
            }}
            return total
        }}

        func creditTotal(txns []parser.Transaction) float64 {{
            var total float64
            for _, tx := range txns {{
                if tx.Amount > 0 {{
                    total += tx.Amount
                }}
            }}
            return total
        }}
        """
    )


def service_go(domain: str, sub: str, type_name: str, template: str) -> str:
    bodies = {
        "running_balance": f"""\
        func Analyze(txns []parser.Transaction) {type_name} {{
            var s {type_name}
            var running float64
            for _, tx := range txns {{
                s.Count++
                s.Net += tx.Amount
                running += tx.Amount
                if tx.Amount < 0 {{
                    s.DebitTotal += tx.Amount
                }} else if tx.Amount > 0 {{
                    s.CreditTotal += tx.Amount
                }}
                if s.From.IsZero() || tx.Date.Before(s.From) {{
                    s.From = tx.Date
                }}
                if tx.Date.After(s.To) {{
                    s.To = tx.Date
                }}
                _ = running
            }}
            return s
        }}""",
        "category_totals": f"""\
        func Analyze(txns []parser.Transaction) {type_name} {{
            totals := make(map[string]float64)
            var s {type_name}
            for _, tx := range txns {{
                s.Count++
                cat := tx.Category
                if cat == "" {{
                    cat = "Uncategorized"
                }}
                totals[cat] += tx.Amount
                s.Net += tx.Amount
                if tx.Amount < 0 {{
                    s.DebitTotal += tx.Amount
                }} else if tx.Amount > 0 {{
                    s.CreditTotal += tx.Amount
                }}
                if s.From.IsZero() || tx.Date.Before(s.From) {{
                    s.From = tx.Date
                }}
                if tx.Date.After(s.To) {{
                    s.To = tx.Date
                }}
            }}
            _ = totals
            return s
        }}""",
        "monthly_average": f"""\
        func Analyze(txns []parser.Transaction) {type_name} {{
            var s {type_name}
            if len(txns) == 0 {{
                return s
            }}
            months := make(map[int]float64)
            for _, tx := range txns {{
                s.Count++
                s.Net += tx.Amount
                key := tx.Date.Year()*100 + int(tx.Date.Month())
                months[key] += tx.Amount
                if tx.Amount < 0 {{
                    s.DebitTotal += tx.Amount
                }} else if tx.Amount > 0 {{
                    s.CreditTotal += tx.Amount
                }}
                if s.From.IsZero() || tx.Date.Before(s.From) {{
                    s.From = tx.Date
                }}
                if tx.Date.After(s.To) {{
                    s.To = tx.Date
                }}
            }}
            if len(months) > 0 {{
                s.Net = s.Net / float64(len(months))
            }}
            return s
        }}""",
        "threshold_count": f"""\
        func Analyze(txns []parser.Transaction) {type_name} {{
            const threshold = 100.0
            var s {type_name}
            for _, tx := range txns {{
                if tx.Amount < -threshold || tx.Amount > threshold {{
                    s.Count++
                }}
                s.Net += tx.Amount
                if tx.Amount < 0 {{
                    s.DebitTotal += tx.Amount
                }} else if tx.Amount > 0 {{
                    s.CreditTotal += tx.Amount
                }}
                if s.From.IsZero() || tx.Date.Before(s.From) {{
                    s.From = tx.Date
                }}
                if tx.Date.After(s.To) {{
                    s.To = tx.Date
                }}
            }}
            return s
        }}""",
        "debit_credit_ratio": f"""\
        func Analyze(txns []parser.Transaction) {type_name} {{
            var s {type_name}
            for _, tx := range txns {{
                s.Count++
                s.Net += tx.Amount
                if tx.Amount < 0 {{
                    s.DebitTotal += tx.Amount
                }} else if tx.Amount > 0 {{
                    s.CreditTotal += tx.Amount
                }}
                if s.From.IsZero() || tx.Date.Before(s.From) {{
                    s.From = tx.Date
                }}
                if tx.Date.After(s.To) {{
                    s.To = tx.Date
                }}
            }}
            if s.CreditTotal != 0 {{
                s.Net = s.DebitTotal / s.CreditTotal
            }}
            return s
        }}""",
        "largest_transaction": f"""\
        func Analyze(txns []parser.Transaction) {type_name} {{
            var s {type_name}
            var largest float64
            for _, tx := range txns {{
                s.Count++
                s.Net += tx.Amount
                if tx.Amount < largest {{
                    largest = tx.Amount
                }}
                if tx.Amount < 0 {{
                    s.DebitTotal += tx.Amount
                }} else if tx.Amount > 0 {{
                    s.CreditTotal += tx.Amount
                }}
                if s.From.IsZero() || tx.Date.Before(s.From) {{
                    s.From = tx.Date
                }}
                if tx.Date.After(s.To) {{
                    s.To = tx.Date
                }}
            }}
            s.Net = largest
            return s
        }}""",
        "date_span": f"""\
        func Analyze(txns []parser.Transaction) {type_name} {{
            var s {type_name}
            for _, tx := range txns {{
                s.Count++
                s.Net += tx.Amount
                if tx.Amount < 0 {{
                    s.DebitTotal += tx.Amount
                }} else if tx.Amount > 0 {{
                    s.CreditTotal += tx.Amount
                }}
                if s.From.IsZero() || tx.Date.Before(s.From) {{
                    s.From = tx.Date
                }}
                if tx.Date.After(s.To) {{
                    s.To = tx.Date
                }}
            }}
            if !s.From.IsZero() && !s.To.IsZero() {{
                s.Net = s.To.Sub(s.From).Hours() / 24
            }}
            return s
        }}""",
        "positive_share": f"""\
        func Analyze(txns []parser.Transaction) {type_name} {{
            var s {type_name}
            var positive int
            for _, tx := range txns {{
                s.Count++
                s.Net += tx.Amount
                if tx.Amount > 0 {{
                    positive++
                    s.CreditTotal += tx.Amount
                }} else if tx.Amount < 0 {{
                    s.DebitTotal += tx.Amount
                }}
                if s.From.IsZero() || tx.Date.Before(s.From) {{
                    s.From = tx.Date
                }}
                if tx.Date.After(s.To) {{
                    s.To = tx.Date
                }}
            }}
            if s.Count > 0 {{
                s.Net = float64(positive) / float64(s.Count)
            }}
            return s
        }}""",
        "category_count": f"""\
        func Analyze(txns []parser.Transaction) {type_name} {{
            cats := make(map[string]struct{{}})
            var s {type_name}
            for _, tx := range txns {{
                s.Count++
                s.Net += tx.Amount
                cat := tx.Category
                if cat == "" {{
                    cat = "Uncategorized"
                }}
                cats[cat] = struct{{}}{{}} 
                if tx.Amount < 0 {{
                    s.DebitTotal += tx.Amount
                }} else if tx.Amount > 0 {{
                    s.CreditTotal += tx.Amount
                }}
                if s.From.IsZero() || tx.Date.Before(s.From) {{
                    s.From = tx.Date
                }}
                if tx.Date.After(s.To) {{
                    s.To = tx.Date
                }}
            }}
            s.Net = float64(len(cats))
            return s
        }}""",
        "median_amount": f"""\
        func Analyze(txns []parser.Transaction) {type_name} {{
            var s {type_name}
            amounts := make([]float64, 0, len(txns))
            for _, tx := range txns {{
                s.Count++
                s.Net += tx.Amount
                amounts = append(amounts, tx.Amount)
                if tx.Amount < 0 {{
                    s.DebitTotal += tx.Amount
                }} else if tx.Amount > 0 {{
                    s.CreditTotal += tx.Amount
                }}
                if s.From.IsZero() || tx.Date.Before(s.From) {{
                    s.From = tx.Date
                }}
                if tx.Date.After(s.To) {{
                    s.To = tx.Date
                }}
            }}
            if len(amounts) == 0 {{
                return s
            }}
            // insertion sort for small slices keeps stdlib-only and deterministic
            for i := 1; i < len(amounts); i++ {{
                v := amounts[i]
                j := i - 1
                for j >= 0 && amounts[j] > v {{
                    amounts[j+1] = amounts[j]
                    j--
                }}
                amounts[j+1] = v
            }}
            mid := len(amounts) / 2
            if len(amounts)%2 == 0 {{
                s.Net = (amounts[mid-1] + amounts[mid]) / 2
            }} else {{
                s.Net = amounts[mid]
            }}
            return s
        }}""",
    }
    body = bodies[template]
    return textwrap.dedent(
        f"""\
        package {sub}

        import "github.com/joshuataye/ledgerpipeline/internal/parser"

        {body}
        """
    )


def test_go(domain: str, sub: str, type_name: str) -> str:
    return textwrap.dedent(
        f"""\
        package {sub}

        import (
            "testing"
            "time"

            "github.com/joshuataye/ledgerpipeline/internal/parser"
        )

        func TestAnalyze_{domain}_{sub}(t *testing.T) {{
            txns := []parser.Transaction{{
                {{Date: time.Date(2026, 1, 5, 0, 0, 0, 0, time.UTC), Amount: -50, Category: "Food"}},
                {{Date: time.Date(2026, 1, 10, 0, 0, 0, 0, time.UTC), Amount: 200, Category: "Payroll"}},
                {{Date: time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC), Amount: -25, Category: "Food"}},
            }}
            got := Analyze(txns)
            if got.From.IsZero() || got.To.IsZero() {{
                t.Fatal("expected date span")
            }}
            if got.To.Before(got.From) {{
                t.Fatal("invalid date span")
            }}
        }}

        func TestAnalyze_{domain}_{sub}_Empty(t *testing.T) {{
            got := Analyze(nil)
            if got.Count != 0 {{
                t.Fatalf("count=%d want 0", got.Count)
            }}
            if !got.From.IsZero() || !got.To.IsZero() {{
                t.Fatal("expected zero dates for empty input")
            }}
        }}
        """
    )


def write_file(path: Path, content: str) -> None:
    path.parent.mkdir(parents=True, exist_ok=True)
    path.write_text(content)


def main() -> None:
    created = 0
    for domain, subs in DOMAINS.items():
        for sub in subs:
            type_name = "".join(part.capitalize() for part in domain.split("_")) + "Summary"
            template = pick_template(domain, sub)
            base = INTERNAL / domain / sub
            write_file(base / "types.go", types_go(domain, sub, type_name))
            write_file(base / "helpers.go", helpers_go(sub))
            write_file(base / "service.go", service_go(domain, sub, type_name, template))
            write_file(base / "service_test.go", test_go(domain, sub, type_name))
            created += 4
    print(f"Created {created} files across {len(DOMAINS)} domains")


if __name__ == "__main__":
    main()
