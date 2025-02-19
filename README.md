A modern, multi-featured command-line tool for DNS analysis, penetration testing, and system administration. DNSKit offers colorized or JSON-formatted output, concurrency for fast subdomain enumeration, and the ability to specify custom DNS resolvers for advanced use cases.

## Table of Contents
1. [Features](#features)
2. [Installation](#installation)
3. [Usage](#usage)
4. [Subcommands](#subcommands)
5. [Advanced Usage](#advanced-usage)
6. [License](#license)

## Features
- **DNS Lookup** for various record types (`A`, `AAAA`, `CNAME`, `MX`, `TXT`, `NS`, `SRV`, etc.).
- **Subdomain Enumeration** with concurrency to speed up scanning.
- **Zone Transfer** checks (`AXFR`) to detect misconfiguration vulnerabilities.
- **Reverse DNS** lookups for IPv4 and IPv6.
- **Custom DNS Resolver** support (e.g., `8.8.8.8`).
- **JSON Output** or **colorized** console output.
- **Cross-Platform** (macOS, Linux, and Windows).

## Installation

1. **Clone** the repository (private in this example):
   ```bash
   git clone git@github.com:nicokempe/dnskit.git
   cd dnskit
   ```

2. **Build** the binary:
   ```bash
   go build -o dnskit main.go
   ```
   This will produce an executable named dnskit (or dnskit.exe on Windows).

3. (Optional) Install it to your `$GOPATH/bin`:
   ```bash
   go install
   ```

## Usage

After building, run:
```bash
./dnskit --help
```
You will see a list of available subcommands and global flags.
**Global Flags**:
* `--json`
    Output results in JSON format instead of colorized/text.
* `--resolver <IP>`
    Use a custom DNS resolver (e.g., 8.8.8.8) for queries.

## Subcommands

1.  **lookup**\
    Performs DNS lookups of a given record type.\
    **Usage**: `dnskit lookup <hostname> --type A|AAAA|MX|TXT|NS|CNAME|SRV`

2.  **enum**\
    Enumerates subdomains using a wordlist with concurrent DNS checks.\
    **Usage**: `dnskit enum <domain> --wordlist subdomains.txt --concurrency 10`

3.  **transfer**\
    Attempts a DNS zone transfer (AXFR) against a nameserver.\
    **Usage**: `dnskit transfer <domain> --nameserver <ns.host>`

4.  **reverse**\
    Performs reverse DNS lookups on a given IP address (supports IPv4/IPv6).\
    **Usage**: `dnskit reverse <ip>`

## Advanced Usage

-   **Concurrency**: You can specify the number of concurrent workers for subdomain enumeration with `--concurrency`. Increase it for larger wordlists or to speed up scans.
-   **DNS Resolver**: By default, DNSKit uses the OS resolver. If you're pentesting or need more control, use `--resolver` to direct all queries through a specific server (e.g., `8.8.8.8:53`).
-   **Color vs. JSON**: Colorized output helps with interactive sessions. For automation (e.g., scripting or storing results), use `--json`.

---

## License

This project is distributed under the **MIT License**. See the [LICENSE file](https://github.com/nicokempe/dnskit/blob/main/LICENSE) for more details.
