package dnsutils

import (
	"fmt"
	"github.com/miekg/dns"
)

// ZoneTransfer attempts to perform a DNS zone transfer
func ZoneTransfer(domain, nameserver string) ([]string, error) {
	if nameserver == "" {
		return nil, fmt.Errorf("no nameserver provided")
	}

	var results []string
	// Making sure, that the nameserver has the port appended if not provided
	if nameserver[:1] == "[" || nameserver[len(nameserver)-1:] == "]" {
		// IPv6 in brackets
		if len(nameserver) >= 2 && nameserver[len(nameserver)-2:] != ":53" {
			nameserver = fmt.Sprintf("%s:53", nameserver)
		}
	} else if len(nameserver) > 0 && nameserver[len(nameserver)-1] != ']' && !containsPort(nameserver) {
		nameserver = fmt.Sprintf("%s:53", nameserver)
	}

	// Preparing AXFR message
	m := new(dns.Msg)
	m.SetAxfr(domain + ".")
	t := new(dns.Transfer)

	// Performing the transfer
	channel, err := t.In(m, nameserver)
	if err != nil {
		return nil, err
	}

	for envelope := range channel {
		if envelope.Error != nil {
			return nil, envelope.Error
		}
		for _, rr := range envelope.RR {
			results = append(results, rr.String())
		}
	}
	return results, nil
}

func containsPort(host string) bool {
	return len(host) > 0 && (host[0] == '[' || (host[len(host)-3:] == ":53"))
}
