[![DNSKit GitHub Banner](/.github/media/banner.svg)](https://www.nicokempe.de)

## ✨ Features

* 🔎 **DNS Lookup** for all major record types (`A`, `AAAA`, `CNAME`, `MX`, `TXT`, `NS`, `SRV`, …)
* 🚀 **Subdomain Enumeration** with concurrency for speed
* 🛡 **Zone Transfer** (`AXFR`) checks to find misconfigurations
* ↩️ **Reverse DNS** lookups (IPv4 + IPv6)
* 🧩 **Custom DNS Resolver** (e.g., `8.8.8.8` or `8.8.8.8:53`)
* 🎨 **JSON or Colorized** output
* 💻 **Cross-Platform** (Windows, macOS, Linux)

## 📦 Installation

### Download prebuilt binaries (recommended)

Grab the latest ZIP/TAR for your OS/arch from the [Releases](https://github.com/nicokempe/dnskit/releases) page and put the `dnskit` binary on your `PATH`.

* **Windows:** download `dnskit_*_windows_amd64.zip` (or `arm64`), unzip, run `dnskit.exe`
* **macOS:** download `dnskit_*_darwin_amd64.tar.gz` (or `arm64`), extract, run `./dnskit`
* **Linux:** download `dnskit_*_linux_amd64.tar.gz` (or `arm64`), extract, run `./dnskit`
  *(Deb/RPM packages are also available under Assets.)*

### Build from source

```bash
git clone https://github.com/nicokempe/dnskit.git
cd dnskit
go build -o dnskit .
# or install to $GOBIN
go install ./...
```

> Official releases are built via GoReleaser and embed version/commit metadata.
> Local builds show `dev` unless you set `-ldflags`.

## 🚀 Usage

```bash
dnskit --help
```

**Global Flags**

* `--json` — Output results in JSON format
* `--resolver <ip[:port]>` — Use a custom DNS resolver

### Subcommands

* **lookup** — Query a specific record type

  ```bash
  dnskit lookup <hostname> --type A|AAAA|MX|TXT|NS|CNAME|SRV
  ```

* **enum** — Enumerate subdomains with concurrency

  ```bash
  dnskit enum <domain> --wordlist subdomains.txt --concurrency 10
  ```

* **transfer** — Attempt DNS zone transfer (AXFR)

  ```bash
  dnskit transfer <domain> --nameserver <ns.host>
  ```

* **reverse** — Reverse DNS lookups (IPv4/IPv6)

  ```bash
  dnskit reverse <ip>
  ```

## 🧾 Versioning

Releases follow the format **`vYYYY.MM.VV`**:

* `YYYY` — year (e.g., `2025`)
* `MM` — month (`01`–`12`)
* `VV` — sequential release number within the month (resets each month)

Examples:

* `v2025.08.1` → first release in August 2025
* `v2025.08.2` → second release in August 2025

## 📝 Release Workflow

1. **Update Changelog**

   Run [changelogen](https://github.com/unjs/changelogen) to update `CHANGELOG.md`:

   ```bash
   npx changelogen --release
   ```

   Commit the updated `CHANGELOG.md`.

2. **Create a new release tag**

   Use the provided script for your OS:

   * **Windows (PowerShell)**

     ```powershell
     ./scripts/new-release.ps1 v2025.08.1
     ```

   * **Linux/macOS (Bash)**

     ```bash
     ./scripts/new-release.sh v2025.08.1
     ```

   These scripts:

   * Validate the version format
   * Commit any pending changes
   * Create a Git tag
   * Push tag → triggers GitHub Actions GoReleaser

3. **GitHub Actions builds & publishes**

   Once the tag is pushed, the CI pipeline:

   * Builds binaries for Linux, macOS, Windows (amd64 + arm64)
   * Packages as `.tar.gz`, `.zip`, `.deb`, `.rpm`
   * Publishes assets to GitHub Releases
   * Updates Winget manifests (if configured)

## 🧰 Development

* Requires **Go 1.24+**

* Quick build:

  ```bash
  go build -o dnskit .
  ```

* Cross-compile + package (local snapshot):

  ```bash
  goreleaser release --snapshot --skip=publish --clean
  ```

## 📝 Changelog

See [CHANGELOG.md](./CHANGELOG.md).
Generated automatically from Conventional Commits via [changelogen](https://github.com/unjs/changelogen).

## 📜 License

MIT — see [LICENSE](./LICENSE).
