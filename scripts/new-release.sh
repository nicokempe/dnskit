#!/usr/bin/env bash
set -euo pipefail

# Ensure we're in a repo and clean
git rev-parse --is-inside-work-tree >/dev/null
test -z "$(git status --porcelain)" || { echo "Working tree is not clean."; exit 1; }

git fetch --tags --prune

year=$(date -u +%Y)
month=$(date -u +%m)
prefix="v${year}.${month}."

# Find the last VV and increment
last_vv=$(git tag -l "${prefix}*" | sed -E 's/^v[0-9]{4}\.[0-9]{2}\.([0-9]+)$/\1/' | sort -n | tail -1)
if [ -z "${last_vv:-}" ]; then next_vv=1; else next_vv=$((last_vv+1)); fi
tag="${prefix}${next_vv}"
echo "Preparing release tag: $tag"

# Optional: preview notes (doesn't alter repo)
if command -v npx >/dev/null 2>&1; then
  npx changelogen@latest --noAuthors --no-output > "${TMPDIR:-/tmp}/RELEASE_NOTES_PREVIEW.md" || true
  echo "Preview notes: ${TMPDIR:-/tmp}/RELEASE_NOTES_PREVIEW.md"
fi

git tag -a "$tag" -m "Release $tag"
git push origin "$tag"
echo "Pushed $tag. The Release workflow will publish binaries/packages shortly."
