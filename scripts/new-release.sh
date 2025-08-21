#!/usr/bin/env bash
set -euo pipefail

git rev-parse --is-inside-work-tree >/dev/null
test -z "$(git status --porcelain)" || { echo "Working tree is not clean."; exit 1; }

git fetch --tags --prune

current_year=$(date -u +%Y)
current_month=$(date -u +%m)
tag_prefix="v${current_year}.${current_month}."
last_sequence=$(git tag -l "${tag_prefix}*" | sed -E 's/^v[0-9]{4}\.[0-9]{2}\.([0-9]+)$/\1/' | sort -n | tail -1)
if [ -z "${last_sequence:-}" ]; then next_sequence=1; else next_sequence=$((last_sequence+1)); fi
tag="${tag_prefix}${next_sequence}"

echo "Preparing release tag: $tag"

git tag -a "$tag" -m "Release $tag"
git push origin "$tag"
echo "Pushed $tag. The Release workflow will publish binaries/packages shortly."
