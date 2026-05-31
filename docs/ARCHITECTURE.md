# Architecture

Ledger Pipeline is a Go CLI and library for ingesting personal bank statements, normalizing transactions, and producing summaries for budgeting and reconciliation.

## Data flow

```
CSV / OFX / QIF / fixed-width / batch directory
        │
        ▼
   internal/parser          ← canonical Transaction model
   internal/importfmt       ← auto-detect import format
        │
        ▼
   internal/pipeline        ← dedupe → filter → normalize → rules → tags → validate
   internal/pipeline/audit  ← audited runs with description checks
        │
        ├── internal/aggregate   → category summaries
        ├── internal/reconcile   → opening/closing balance check
        ├── internal/recurring   → recurring charge candidates
        ├── internal/matching    → transfer pair detection
        ├── internal/anomaly     → outlier detection
        ├── internal/forecast    → monthly net projection
        ├── internal/insights    → spending insight reports
        ├── internal/budget      → budget variance and utilization
        ├── internal/tax         → deductible category totals
        ├── internal/compare     → period-over-period deltas
        ├── internal/ledger      → double-entry journal lines
        ├── internal/storage     → JSON transaction snapshots
        └── internal/enrich      → merchant normalization enrichment
        │
        ▼
   internal/report / internal/export / internal/format
```

The orchestrator (`internal/orchestration`) wraps the pipeline with account lookup and optional budget comparison for multi-account batch runs.

## Package map

| Package | Responsibility |
|---------|----------------|
| `internal/parser` | CSV ingestion and row validation |
| `internal/import/ofx` | OFX snippet parsing |
| `internal/import/qif` | Quicken Interchange Format parsing |
| `internal/import/fixedwidth` | Fixed-width bank export lines |
| `internal/importfmt` | Format detection and unified import |
| `internal/dedupe` | Fingerprint-based duplicate removal |
| `internal/filter` | Date, amount, and category filters |
| `internal/filter/presets` | JSON config and last-N-days presets |
| `internal/merchant` | Merchant description normalization and suffix rules |
| `internal/categorize/rules` | Priority-based auto-categorization |
| `internal/tags` | Tag enrichment from rule maps |
| `internal/enrich` | Combined normalization enrichment |
| `internal/validate` | Post-parse invariant checks |
| `internal/aggregate` | Per-category rollups |
| `internal/stats` | Debit/credit statistics helpers |
| `internal/period` | Monthly and boundary rollups |
| `internal/compare` | Period-over-period comparison |
| `internal/money` | Amount parsing and helpers |
| `internal/matching` | Internal transfer pair matching |
| `internal/split` | Split transactions across categories |
| `internal/anomaly` | Z-score outlier detection |
| `internal/forecast` | Linear monthly net forecast |
| `internal/insights` | Top-category and spending insights |
| `internal/tax` | Deductible category reporting |
| `internal/audit` | Pipeline stage audit log |
| `internal/accountsfile` | JSON account file loading |
| `internal/ledger` | Double-entry journal generation |
| `internal/storage` | JSON transaction snapshots |
| `internal/format` | CSV, TSV, and Markdown export writers |
| `internal/budget` | Budget variance and utilization analysis |
| `internal/recurring` | Recurring detection and interval estimation |
| `internal/pipeline` | Stage orchestration |
| `internal/pipeline/stages` | Pluggable stage interface |
| `internal/pipeline/audit` | Audited pipeline runs |
| `internal/orchestration` | Account-aware batch runs |
| `internal/account` | Account registry |
| `internal/batchfile` | Multi-file CSV directory ingestion |
| `internal/config` | JSON profile and filter config parsing |
| `internal/cliutil` | Shared CLI flag helpers |
| `internal/amortization`, `internal/allocation`, `internal/cashflow`, … | Domain-specific analysis libraries (51 top-level domains) |
| `cmd/ledgerpipeline` | CLI entrypoint |

## Design choices

- **Zero external dependencies** — the tool targets air-gapped or containerized runs with only the Go standard library.
- **Library-first** — each stage is a plain Go package with tests; the CLI is a thin wrapper.
- **Explicit pipeline config** — stages are toggled via `pipeline.Config` rather than implicit globals, making behavior easy to test in isolation.

## Testing

```bash
go test ./...
```

Integration tests use fixtures under `testdata/`. Import parsers and the pipeline have table-driven unit tests per package.
