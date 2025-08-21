#!/usr/bin/env pwsh
<#
Creates and pushes the next tag in the scheme vYYYY.MM.VV and triggers the CI release.
- Ensures clean working tree
- Computes VV = last tag this month + 1
- Optionally previews changelog with changelogen (no repo changes)
#>

$ErrorActionPreference = "Stop"

# 1) Safety checks: Require git + a clean working tree
git rev-parse --is-inside-work-tree *> $null
if ($LASTEXITCODE -ne 0) { throw "Not a git repository." }

$dirty = (git status --porcelain)
if ($dirty) { throw "Working tree is not clean. Commit or stash changes first." }

# 2) Compute next tag vYYYY.MM.VV
git fetch --tags --prune

$currentDate  = Get-Date
$currentYear  = $currentDate.ToString('yyyy')
$currentMonth = $currentDate.ToString('MM')
$tagPrefix    = "v$currentYear.$currentMonth."

$existingTags = git tag --list "$tagPrefix*" | Sort-Object -Descending
$nextSequence = 1
if ($existingTags.Count -gt 0) {
  if ($existingTags[0] -match '^v\d{4}\.\d{2}\.(\d+)$') {
    $nextSequence = [int]$Matches[1] + 1
  }
}
$tag = "$tagPrefix$nextSequence"

Write-Host "Preparing release tag: $tag" -ForegroundColor Cyan

# 3) Create annotated tag (message is helpful in git UI; CI doesn’t require it)
$tagMessage = "Release $tag"
git tag -a "$tag" -m "$tagMessage"

# 4) Push tag to trigger GitHub Actions (GoReleaser)
git push origin "$tag"

Write-Host "Pushed $tag. The Release workflow will publish binaries/packages shortly." -ForegroundColor Green
