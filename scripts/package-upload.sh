#!/usr/bin/env bash
set -euo pipefail
ROOT="$(cd "$(dirname "$0")/.." && pwd)"
OUT="${1:-$ROOT/../../dist/ledger-pipeline.zip}"
cd "$ROOT"

if [[ ! -d .git ]]; then
  echo "ERROR: missing .git" >&2
  exit 1
fi

for required in .git/HEAD .git/config .git/refs/heads/main .git/objects; do
  if [[ ! -e "$required" ]]; then
    echo "ERROR: incomplete .git directory (missing ${required})" >&2
    exit 1
  fi
done

if ! git rev-parse --verify HEAD >/dev/null 2>&1; then
  echo "ERROR: git HEAD is not a valid commit" >&2
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

# Stale worktrees reference absolute paths on the packager machine and break
# git history reads after the archive is extracted elsewhere (e.g. Silver).
git worktree prune 2>/dev/null || true
rm -rf .git/worktrees

COMMITS="$(git rev-list --count HEAD)"
if [[ "${COMMITS}" -lt 15 ]]; then
  echo "ERROR: expected at least 15 commits, found ${COMMITS}" >&2
  exit 1
fi

git fsck --no-progress >/dev/null
git repack -adf >/dev/null

mkdir -p "$(dirname "$OUT")"
rm -f "$OUT"
export COPYFILE_DISABLE=1
# Package from repo root so Dockerfile and .git/ sit at the zip root (required by Silver).
zip -r "$OUT" . \
  -x ".git/worktrees/*" \
  -x ".history-snapshot/*" \
  -x ".DS_Store" \
  -x "**/.DS_Store" \
  -x "__MACOSX/*"

if ! grep -Fxq '.git/HEAD' < <(zipinfo -1 "$OUT"); then
  echo "ERROR: zip is missing .git/HEAD — archive would fail Silver git validation" >&2
  exit 1
fi

if grep -q '^\.git/worktrees/' < <(zipinfo -1 "$OUT"); then
  echo "ERROR: zip still contains .git/worktrees — remove stale worktrees first" >&2
  exit 1
fi

VERIFY_DIR="$(mktemp -d "${TMPDIR:-/tmp}/ledger-pipeline-verify.XXXXXX")"
trap 'rm -rf "$VERIFY_DIR"' EXIT
unzip -q "$OUT" -d "$VERIFY_DIR"
(
  cd "$VERIFY_DIR"
  git rev-parse --verify HEAD >/dev/null
  EXTRACTED="$(git rev-list --count HEAD)"
  if [[ "${EXTRACTED}" -ne "${COMMITS}" ]]; then
    echo "ERROR: extracted archive has ${EXTRACTED} commits, expected ${COMMITS}" >&2
    exit 1
  fi
  git rev-parse --verify task/01-base >/dev/null
  git rev-parse --verify task/22-base >/dev/null
)

GIT_ENTRIES="$(zipinfo -1 "$OUT" | grep -c '^\.git/' || true)"
echo "Created $OUT ($(du -h "$OUT" | cut -f1), ${COMMITS} commits, ${GIT_ENTRIES} .git entries)"
echo "Upload this zip only — do NOT use GitHub \"Download ZIP\" (it excludes .git)."
echo "Run this script from inside the repo; do NOT zip the parent folder."
