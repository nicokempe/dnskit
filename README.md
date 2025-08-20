[![DNSKit GitHub Banner](/.github/banner.webp)](https://www.nicokempe.de)

## ✨ Features

* 🔎 **DNS Lookup** for all major record types (`A`, `AAAA`, `CNAME`, `MX`, `TXT`, `NS`, `SRV`, etc.)
* 🚀 **Subdomain Enumeration** with concurrency for speed
* 🛡 **Zone Transfer** (`AXFR`) checks to find misconfigurations
* ↩️ **Reverse DNS** lookups (IPv4 + IPv6)
* 🧩 **Custom DNS Resolver** support (e.g., `8.8.8.8`)
* 🎨 **JSON or Colorized** output for flexible workflows
* 💻 **Cross-Platform** (macOS, Linux, Windows)

## ⚙️ Installation

1. **Clone** the repository

   ```bash
   git clone git@github.com:nicokempe/dnskit.git
   cd dnskit
   ```

2. **Build** the binary

   ```bash
   go build -o dnskit main.go
   ```

   → Produces an executable named `dnskit` (or `dnskit.exe` on Windows).

3. (Optional) **Install globally**

   ```bash
   go install
   ```

## 🖥 Usage

Run after building:

```bash
./dnskit --help
```

🔧 **Global Flags**

* `--json` → Output results in JSON format
* `--resolver <IP>` → Use a custom resolver (e.g., `8.8.8.8`)

## 🛠 Subcommands

* 🔎 **lookup**
  Perform DNS lookups of a given record type

  ```bash
  dnskit lookup <hostname> --type A|AAAA|MX|TXT|NS|CNAME|SRV
  ```

* 🌐 **enum**
  Enumerate subdomains with concurrency

  ```bash
  dnskit enum <domain> --wordlist subdomains.txt --concurrency 10
  ```

* 🛡 **transfer**
  Attempt a DNS zone transfer (AXFR)

  ```bash
  dnskit transfer <domain> --nameserver <ns.host>
  ```

* ↩️ **reverse**
  Reverse DNS lookups for IPv4/IPv6

  ```bash
  dnskit reverse <ip>
  ```

## 🔧 Advanced Usage

* ⚡ **Concurrency** → Adjust with `--concurrency` for faster/larger scans
* 🧩 **Custom Resolver** → Override OS defaults with `--resolver 8.8.8.8:53`
* 🎨 **Color vs JSON** → Use color for interactive use, JSON for automation/scripts

## 📜 License

This project is licensed under the [MIT License](https://github.com/nicokempe/dnskit/blob/main/LICENSE).
