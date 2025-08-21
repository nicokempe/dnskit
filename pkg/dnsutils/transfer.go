package dnsutils

import (
	"fmt"
	"github.com/miekg/dns"
)

// ZoneTransfer attempts a DNS zone transfer (AXFR) for the specified domain
// using the supplied nameserver address.
func ZoneTransfer(domain, nameserverAddr string) ([]string, error) {
	if nameserverAddr == "" {
		return nil, fmt.Errorf("no nameserver provided")
	}

	var zoneRecords []string
	// Ensure the nameserver address includes a port; default to 53.
	if nameserverAddr[:1] == "[" || nameserverAddr[len(nameserverAddr)-1:] == "]" {
		if len(nameserverAddr) >= 2 && nameserverAddr[len(nameserverAddr)-2:] != ":53" {
			nameserverAddr = fmt.Sprintf("%s:53", nameserverAddr)
		}
	} else if len(nameserverAddr) > 0 && nameserverAddr[len(nameserverAddr)-1] != ']' && !containsPort(nameserverAddr) {
		nameserverAddr = fmt.Sprintf("%s:53", nameserverAddr)
	}

	zoneTransferRequest := new(dns.Msg)
	zoneTransferRequest.SetAxfr(domain + ".")

	dnsTransferClient := new(dns.Transfer)

	transferMessages, err := dnsTransferClient.In(zoneTransferRequest, nameserverAddr)

	if err != nil {
		return nil, err
	}

	for envelope := range transferMessages {
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
