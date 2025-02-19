package dnsutils

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

// Enumerate performs a simple subdomain enumeration using a wordlist
func Enumerate(domain, wordlistPath string) ([]string, error) {
	var subdomains []string

	// If no wordlist is provided, optionally adefault is used
	if wordlistPath == "" {
		return nil, fmt.Errorf("no wordlist provided")
	}

	file, err := os.Open(wordlistPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sub := scanner.Text()
		candidate := fmt.Sprintf("%s.%s", sub, domain)
		ips, _ := net.LookupIP(candidate)
		if len(ips) > 0 {
			subdomains = append(subdomains, candidate)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return subdomains, nil
}
