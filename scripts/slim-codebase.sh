#!/usr/bin/env bash
# Remove generator scripts before platform upload.
set -euo pipefail
ROOT="$(cd "$(dirname "$0")/.." && pwd)"
cd "$ROOT"

rm -f scripts/expand-ledger-codebase.py
rm -rf scripts/__pycache__

echo "Slim-down complete."
