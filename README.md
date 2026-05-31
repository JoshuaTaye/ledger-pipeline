# Ledger Pipeline

Go CLI and library for batch processing bank statements: CSV ingestion, deduplication, categorization, budget comparison, reconciliation, recurring charge detection, transfer matching, anomaly detection, forecasting, and multi-format export.

Built as a personal finance tooling project with a broad library surface: ingestion, pipeline stages, budgeting, reconciliation, forecasting, compliance, FX, payroll, and portfolio analysis packages wired through a single CLI.

## Scale

- **700+ Go source files** across 80+ internal packages
- Table-driven unit tests per package (`go test ./...`)
- Full git history with task base-commit tags for SWE-bench-style evaluation

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
docker build -t ledger-pipeline:local .
```

## Usage

### Core commands

```bash
# Summarize a CSV statement (supports date/category filters)
./ledgerpipeline summarize -input testdata/sample.csv
./ledgerpipeline summarize -input testdata/sample.csv -from 2026-01-01 -to 2026-01-31 -categories Food,Transport

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
```

### Import commands

```bash
./ledgerpipeline import-ofx -input testdata/sample.ofx
./ledgerpipeline import-qif -input testdata/sample.qif
./ledgerpipeline import-fixedwidth -input testdata/fixedwidth.txt
./ledgerpipeline import-any -input testdata/sample.csv
```

### Analysis commands

```bash
./ledgerpipeline insights -input testdata/sample.csv
./ledgerpipeline monthly -input testdata/sample.csv
./ledgerpipeline stats -input testdata/sample.csv
./ledgerpipeline match-transfers -input testdata/sample.csv
./ledgerpipeline anomalies -input testdata/sample.csv -threshold 2.0
./ledgerpipeline forecast -input testdata/sample.csv -months 3
./ledgerpipeline compare -before testdata/sample.csv -after testdata/sample.csv
./ledgerpipeline intervals -input testdata/sample.csv -min-count 2
./ledgerpipeline budget-analysis -input testdata/sample.csv -limits Food:200,Transport:50
./ledgerpipeline tax-report -input testdata/sample.csv
```

### Export and storage

```bash
./ledgerpipeline format-export -input testdata/sample.csv -format markdown
./ledgerpipeline format-export -input testdata/sample.csv -format tsv -output summary.tsv
./ledgerpipeline snapshot -input testdata/sample.csv -output snapshot.json
./ledgerpipeline snapshot -load snapshot.json
```

### Filter presets

```bash
./ledgerpipeline filter-preset -input testdata/sample.csv -config testdata/filter-config.json
./ledgerpipeline filter-preset -input testdata/sample.csv -days 30
```

## Project layout

| Path | Role |
|------|------|
| `cmd/ledgerpipeline` | CLI entrypoint |
| `internal/parser` | CSV parsing |
| `internal/pipeline` | Configurable processing pipeline |
| `internal/matching` | Transfer pair detection |
| `internal/anomaly` | Outlier detection |
| `internal/forecast` | Monthly net projection |
| `internal/format` | CSV/TSV/Markdown writers |
| `internal/importfmt` | Auto-detect import format |
| `internal/insights` | Spending insight reports |
| `internal/orchestration` | Account-aware batch orchestrator |
| `docs/ARCHITECTURE.md` | Data flow and package map |

See `docs/ARCHITECTURE.md` for the full design overview.

## License

Private project — all rights reserved.
