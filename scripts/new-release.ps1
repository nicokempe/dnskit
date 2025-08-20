#!/usr/bin/env pwsh
<#
Creates and pushes the next tag in the scheme vYYYY.MM.VV and triggers the CI release.
- Ensures clean working tree
- Computes VV = last tag this month + 1
- Optionally previews changelog with changelogen (no repo changes)
#>

$ErrorActionPreference = "Stop"

# 1) Safety checks
# Require git + a clean working tree
git rev-parse --is-inside-work-tree *> $null
if ($LASTEXITCODE -ne 0) { throw "Not a git repository." }

$dirty = (git status --porcelain)
if ($dirty) { throw "Working tree is not clean. Commit or stash changes first." }

# 2) Compute next tag vYYYY.MM.VV
git fetch --tags --prune

$now   = Get-Date
$year  = $now.ToString('yyyy')
$month = $now.ToString('MM')
$prefix = "v$year.$month."

# List tags for current month and extract VV
$existing = git tag --list "$prefix*" | Sort-Object -Descending
$nextVV = 1
if ($existing.Count -gt 0) {
  if ($existing[0] -match '^v\d{4}\.\d{2}\.(\d+)$') {
    $nextVV = [int]$Matches[1] + 1
  }
}
$tag = "$prefix$nextVV"

Write-Host "Preparing release tag: $tag" -ForegroundColor Cyan

# 3) (Optional) Preview release notes with changelogen (no repo changes)
try {
  # If changelogen is available, preview to temp file in UTF-8
  npx changelogen@latest --noAuthors --no-output `
    | Set-Content -Encoding utf8 "$env:TEMP\RELEASE_NOTES_PREVIEW.md"
  Write-Host "Preview notes: $env:TEMP\RELEASE_NOTES_PREVIEW.md"
} catch {
  Write-Host "Note preview skipped (changelogen not available): $($_.Exception.Message)"
}

# 4) Create annotated tag (message is helpful in git UI; CI doesn’t require it)
$tagMessage = "Release $tag"
git tag -a "$tag" -m "$tagMessage"

# 5) Push tag to trigger GitHub Actions (GoReleaser)
git push origin "$tag"

Write-Host "Pushed $tag. The Release workflow will publish binaries/packages shortly." -ForegroundColor Green
