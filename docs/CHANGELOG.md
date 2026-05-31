# Changelog

All notable changes to ledger-pipeline are documented here.

## 0.2.0 — 2026-05-31

- Removed unused generated analysis modules; focused the codebase on wired pipeline packages
- Added `-rules` flag to `summarize` for JSON categorization rules
- Added `batch`, `import-ofx`, and `import-qif` CLI commands
- Added `rules.LoadFile` and file-based OFX/QIF import helpers
- Documented architecture in `docs/ARCHITECTURE.md`

## 0.1.0 — 2024-11-04

- CSV parser with validation and amount helpers
- Configurable processing pipeline (dedupe, filter, normalize, tags, reconcile)
- CLI: summarize, reconcile, budget, recurring, export
- Batch orchestrator with account registry
- OFX, QIF, and fixed-width import parsers
- Priority-based categorization rules engine
- Budget comparison and recurring charge detection
