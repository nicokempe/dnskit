// pkg/dnsutils/lookup.go
package dnsutils

import (
	"context"
	"fmt"
	"strings"
)

// Lookup performs a DNS lookup for the specified record type using the optional
// custom resolver.
func Lookup(hostname, recordType, resolverAddr string) ([]string, error) {
	var (
		records []string
		err     error
	)

	resolver := getResolver(resolverAddr)
	lookupCtx := context.Background()

	switch strings.ToUpper(recordType) {
	case "A":
		ipAddresses, err := resolver.LookupIPAddr(lookupCtx, hostname)
		if err != nil {
			return nil, err
		}
		for _, ipAddress := range ipAddresses {
			if ipv4 := ipAddress.IP.To4(); ipv4 != nil {
				records = append(records, ipv4.String())
			}
		}
	case "AAAA":
		ipAddresses, err := resolver.LookupIPAddr(lookupCtx, hostname)
		if err != nil {
			return nil, err
		}
		for _, ipAddress := range ipAddresses {
			if ipv6 := ipAddress.IP.To16(); ipv6 != nil && ipv6.To4() == nil {
				records = append(records, ipv6.String())
			}
		}
	case "CNAME":
		canonicalName, err := resolver.LookupCNAME(lookupCtx, hostname)
		if err != nil {
			return nil, err
		}
		records = append(records, cname)
	case "MX":
		mxRecords, err := resolver.LookupMX(lookupCtx, hostname)
		if err != nil {
			return nil, err
		}
		for _, mxRecord := range mxRecords {
			records = append(records, fmt.Sprintf("%v %v", mxRecord.Host, mxRecord.Pref))
		}
	case "TXT":
		txtRecords, err := resolver.LookupTXT(lookupCtx, hostname)
		if err != nil {
			return nil, err
		}
		records = append(records, txtRecords...)
	case "NS":
		nameserverRecords, err := resolver.LookupNS(lookupCtx, hostname)
		if err != nil {
			return nil, err
		}
		for _, nameserverRecord := range nameserverRecords {
			records = append(records, nameserverRecord.Host)
		}
	case "SRV":
		hostnameParts := strings.Split(hostname, ".")
		if len(hostnameParts) < 3 || !strings.HasPrefix(hostnameParts[0], "_") || !strings.HasPrefix(hostnameParts[1], "_") {
			return nil, fmt.Errorf("SRV lookups require hostname in format _service._proto.name")
		}
		serviceName := strings.TrimPrefix(hostnameParts[0], "_")
		protocol := strings.TrimPrefix(hostnameParts[1], "_")
		baseName := strings.Join(hostnameParts[2:], ".")
		_, srvRecords, err := resolver.LookupSRV(lookupCtx, serviceName, protocol, baseName)
		if err != nil {
			return nil, err
		}
		for _, srvRecord := range srvRecords {
			records = append(records, fmt.Sprintf("%s %d %d %d", srvRecord.Target, srvRecord.Port, srvRecord.Priority, srvRecord.Weight))
		}
	default:
		return nil, fmt.Errorf("unsupported record type: %s", recordType)
	}

	return records, err
}
