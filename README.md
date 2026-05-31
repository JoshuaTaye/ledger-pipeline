# Ledger Pipeline

Go CLI and library for batch processing bank statements: CSV ingestion, deduplication, categorization, budget comparison, reconciliation, recurring charge detection, and JSON export.

Built as a personal finance tooling project — a focused codebase with real pipeline integration, not a generated module tree.

## Requirements

- Go 1.22+
- Docker (optional, for containerized builds)

## Build and test

```bash
go test ./...
go build -o ledgerpipeline ./cmd/ledgerpipeline
```

## Docker

```bash
docker build -t ledger-pipeline .
```

## Usage

```bash
# Summarize a CSV statement
./ledgerpipeline summarize -input testdata/sample.csv

# Apply categorization rules from JSON
./ledgerpipeline summarize -input testdata/sample.csv -rules testdata/rules.json -normalize

# Reconcile opening and closing balances
./ledgerpipeline reconcile -input testdata/sample.csv -opening 500 -closing 480

# Compare against budget limits
./ledgerpipeline budget -input testdata/sample.csv -limits Food:200,Transport:50

# Detect recurring charges
./ledgerpipeline recurring -input testdata/sample.csv -min-count 2

# Export JSON summary
./ledgerpipeline export -input testdata/sample.csv -output summary.json

# Process every CSV in a directory
./ledgerpipeline batch -dir testdata/statements -dedupe

# Convert OFX or QIF to stdout rows
./ledgerpipeline import-ofx -input testdata/sample.ofx
./ledgerpipeline import-qif -input testdata/sample.qif
```

## Project layout

| Path | Role |
|------|------|
| `cmd/ledgerpipeline` | CLI entrypoint |
| `internal/parser` | CSV parsing |
| `internal/pipeline` | Configurable processing pipeline |
| `internal/orchestration` | Account-aware batch orchestrator |
| `internal/categorize/rules` | Priority-based categorization |
| `internal/import/` | OFX, QIF, and fixed-width parsers |
| `docs/ARCHITECTURE.md` | Data flow and package map |

See `docs/ARCHITECTURE.md` for the full design overview.

## License

Private project — all rights reserved.
