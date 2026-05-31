# Architecture

Ledger Pipeline is a Go CLI and library for ingesting personal bank statements, normalizing transactions, and producing summaries for budgeting and reconciliation.

## Data flow

```
CSV / OFX / QIF / batch directory
        │
        ▼
   internal/parser          ← canonical Transaction model
        │
        ▼
   internal/pipeline        ← dedupe → filter → normalize → rules → tags → validate
        │
        ├── internal/aggregate   → category summaries
        ├── internal/reconcile   → opening/closing balance check
        ├── internal/recurring   → recurring charge candidates
        └── internal/budget      → budget variance
        │
        ▼
   internal/report / internal/export
```

The orchestrator (`internal/orchestration`) wraps the pipeline with account lookup and optional budget comparison for multi-account batch runs.

## Package map

| Package | Responsibility |
|---------|----------------|
| `internal/parser` | CSV ingestion and row validation |
| `internal/import/ofx` | OFX snippet parsing |
| `internal/import/qif` | Quicken Interchange Format parsing |
| `internal/import/fixedwidth` | Fixed-width bank export lines |
| `internal/dedupe` | Fingerprint-based duplicate removal |
| `internal/filter` | Date and amount filters |
| `internal/merchant` | Merchant description normalization |
| `internal/categorize/rules` | Priority-based auto-categorization |
| `internal/tags` | Tag enrichment from rule maps |
| `internal/validate` | Post-parse invariant checks |
| `internal/aggregate` | Per-category rollups |
| `internal/stats` | Debit/credit statistics helpers |
| `internal/period` | Monthly and boundary rollups |
| `internal/money` | Amount parsing and helpers |
| `internal/pipeline` | Stage orchestration |
| `internal/pipeline/stages` | Pluggable stage interface |
| `internal/orchestration` | Account-aware batch runs |
| `internal/account` | Account registry |
| `internal/batchfile` | Multi-file CSV directory ingestion |
| `internal/config` | JSON profile parsing |
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
