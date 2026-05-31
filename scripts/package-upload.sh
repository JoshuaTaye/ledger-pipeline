#!/usr/bin/env bash
set -euo pipefail
ROOT="$(cd "$(dirname "$0")/.." && pwd)"
OUT="${1:-$ROOT/../../dist/ledger-pipeline.zip}"
cd "$ROOT"

if [[ ! -d .git ]]; then
  echo "ERROR: missing .git" >&2
  exit 1
fi

if [[ -n "$(git status --porcelain)" ]]; then
  echo "ERROR: uncommitted changes — commit before packaging" >&2
  git status --short >&2
  exit 1
fi

if [[ -f scripts/expand-ledger-codebase.py ]]; then
  echo "ERROR: scaffold scripts still present — run scripts/slim-codebase.sh" >&2
  exit 1
fi

COMMITS="$(git rev-list --count HEAD)"
if [[ "${COMMITS}" -lt 25 ]]; then
  echo "ERROR: expected at least 25 commits, found ${COMMITS}" >&2
  exit 1
fi

mkdir -p "$(dirname "$OUT")"
rm -f "$OUT"
export COPYFILE_DISABLE=1
zip -r "$OUT" . -x ".DS_Store" -x "**/.DS_Store" -x "__MACOSX/*"

if ! grep -Fxq '.git/HEAD' < <(zipinfo -1 "$OUT"); then
  echo "ERROR: zip is missing .git/HEAD" >&2
  exit 1
fi

echo "Created $OUT ($(du -h "$OUT" | cut -f1), ${COMMITS} commits)"
