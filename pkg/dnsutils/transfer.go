package dnsutils

import (
	"fmt"
	"github.com/miekg/dns"
)

// ZoneTransfer attempts to perform a DNS zone transfer against the provided nameserver.
func ZoneTransfer(domain, nameserverAddr string) ([]string, error) {
	if nameserverAddr == "" {
		return nil, fmt.Errorf("no nameserver provided")
	}

	var zoneRecords []string
	// Ensure the nameserver address includes a port; default to 53.
	if nameserverAddr[:1] == "[" || nameserverAddr[len(nameserverAddr)-1:] == "]" {
		// IPv6 in brackets
		if len(nameserverAddr) >= 2 && nameserverAddr[len(nameserverAddr)-2:] != ":53" {
			nameserverAddr = fmt.Sprintf("%s:53", nameserverAddr)
		}
	} else if len(nameserverAddr) > 0 && nameserverAddr[len(nameserverAddr)-1] != ']' && !containsPort(nameserverAddr) {
		nameserverAddr = fmt.Sprintf("%s:53", nameserverAddr)
	}

	// Prepare AXFR message
	zoneTransferMsg := new(dns.Msg)
	zoneTransferMsg.SetAxfr(domain + ".")

	transferClient := new(dns.Transfer)

	// Perform the transfer
	transferResults, err := transferClient.In(zoneTransferMsg, nameserverAddr)
	if err != nil {
		return nil, err
	}

	for envelope := range transferResults {
		if envelope.Error != nil {
			return nil, envelope.Error
		}
		for _, resourceRecord := range envelope.RR {
			zoneRecords = append(zoneRecords, resourceRecord.String())
		}
	}
	return zoneRecords, nil
}

// containsPort checks whether the host string already specifies a port.
func containsPort(hostAddress string) bool {
	return len(hostAddress) > 0 && (hostAddress[0] == '[' || (hostAddress[len(hostAddress)-3:] == ":53"))
}
