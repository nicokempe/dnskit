// pkg/dnsutils/lookup.go
package dnsutils

import (
	"fmt"
	"net"
	"strings"
)

// Lookup performs a basic DNS lookup for the specified record type
func Lookup(hostname, recordType string) ([]string, error) {
	var results []string
	var err error

	switch strings.ToUpper(recordType) {
	case "A":
		ips, err := net.LookupIP(hostname)
		if err != nil {
			return nil, err
		}
		for _, ip := range ips {
			if ipv4 := ip.To4(); ipv4 != nil {
				results = append(results, ipv4.String())
			}
		}
	case "AAAA":
		ips, err := net.LookupIP(hostname)
		if err != nil {
			return nil, err
		}
		for _, ip := range ips {
			if ipv6 := ip.To16(); ipv6 != nil && ipv6.To4() == nil {
				results = append(results, ipv6.String())
			}
		}
	case "CNAME":
		cname, err := net.LookupCNAME(hostname)
		if err != nil {
			return nil, err
		}
		results = append(results, cname)
	case "MX":
		mxRecords, err := net.LookupMX(hostname)
		if err != nil {
			return nil, err
		}
		for _, mx := range mxRecords {
			results = append(results, fmt.Sprintf("%v %v", mx.Host, mx.Pref))
		}
	case "TXT":
		txtRecords, err := net.LookupTXT(hostname)
		if err != nil {
			return nil, err
		}
		results = append(results, txtRecords...)
	default:
		return nil, fmt.Errorf("unsupported record type: %s", recordType)
	}

	return results, err
}
