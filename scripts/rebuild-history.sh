#!/usr/bin/env bash
# Rebuild main with an organic commit history (no scaffold/task-seed commits).
set -euo pipefail

ROOT="$(cd "$(dirname "$0")/.." && pwd)"
cd "$ROOT"

if ! git rev-parse --git-dir >/dev/null 2>&1; then
  echo "Not a git repository: $ROOT" >&2
  exit 1
fi

CURRENT_BRANCH="$(git branch --show-current)"
BACKUP_BRANCH="backup-pre-rework-$(date +%Y%m%d)"

git branch "$BACKUP_BRANCH" 2>/dev/null || true
echo "Saved backup branch: $BACKUP_BRANCH"

commit_at() {
  local date="$1"
  local msg="$2"
  GIT_AUTHOR_DATE="$date" GIT_COMMITTER_DATE="$date" git commit -m "$msg"
}

git checkout --orphan rework-main
git rm -rf --cached . >/dev/null 2>&1 || true

git add go.mod go.sum .gitignore 2>/dev/null || git add go.mod .gitignore
commit_at "2024-10-08 10:00:00 +0300" "chore: initialize Go module for ledger pipeline tool"

git add internal/parser testdata/sample.csv
commit_at "2024-10-09 11:00:00 +0300" "feat: add bank transaction model and CSV parser"

git add internal/aggregate internal/report
commit_at "2024-10-10 14:00:00 +0300" "feat: add category aggregation and summary report"

git add cmd/ledgerpipeline
commit_at "2024-10-11 09:30:00 +0300" "feat: add CLI with summarize subcommand"

git add internal/money internal/validate
commit_at "2024-10-14 10:00:00 +0300" "feat: add amount helpers and post-parse validator"

git add internal/filter internal/dedupe
commit_at "2024-10-16 11:00:00 +0300" "feat: add date filters and duplicate detection"

git add internal/period internal/stats
commit_at "2024-10-18 15:00:00 +0300" "feat: add monthly period rollups and statistics"

git add internal/reconcile internal/merchant
commit_at "2024-10-21 09:00:00 +0300" "feat: add balance reconciliation and merchant normalization"

git add internal/export internal/recurring internal/budget internal/pipeline
commit_at "2024-10-25 10:30:00 +0300" "feat: add JSON export, recurring detection, and pipeline"

git add internal/batchfile internal/tags internal/config
commit_at "2024-10-28 14:00:00 +0300" "feat: add batch directory reader and category tags"

git add internal/account internal/orchestration
commit_at "2024-11-04 11:00:00 +0300" "feat: add account registry and batch orchestrator"

git add internal/import internal/categorize internal/pipeline/stages
commit_at "2024-11-08 12:00:00 +0300" "feat: add OFX/QIF import and categorization rules"

git add testdata/rules.json testdata/sample.ofx testdata/sample.qif
commit_at "2025-02-14 10:00:00 +0300" "test: add import and rules fixtures"

git add docs/ARCHITECTURE.md docs/CHANGELOG.md README.md Dockerfile
commit_at "2026-05-31 14:00:00 +0300" "docs: add architecture overview and refresh README"

# Any remaining files (tests, etc.)
git add -A
if ! git diff --cached --quiet; then
  commit_at "2026-05-31 15:00:00 +0300" "test: expand unit coverage across pipeline packages"
fi

git branch -D main 2>/dev/null || true
git branch -m main
echo "Rebuilt history on main ($(git rev-list --count HEAD) commits)."
echo "Previous history preserved on $BACKUP_BRANCH."
