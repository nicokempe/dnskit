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
		results []string
		err     error
	)

	r := getResolver(resolverAddr)
	ctx := context.Background()

	switch strings.ToUpper(recordType) {
	case "A":
		ips, err := r.LookupIPAddr(ctx, hostname)
		if err != nil {
			return nil, err
		}
		for _, ip := range ips {
			if ipv4 := ip.IP.To4(); ipv4 != nil {
				results = append(results, ipv4.String())
			}
		}
	case "AAAA":
		ips, err := r.LookupIPAddr(ctx, hostname)
		if err != nil {
			return nil, err
		}
		for _, ip := range ips {
			if ipv6 := ip.IP.To16(); ipv6 != nil && ipv6.To4() == nil {
				results = append(results, ipv6.String())
			}
		}
	case "CNAME":
		cname, err := r.LookupCNAME(ctx, hostname)
		if err != nil {
			return nil, err
		}
		results = append(results, cname)
	case "MX":
		mxRecords, err := r.LookupMX(ctx, hostname)
		if err != nil {
			return nil, err
		}
		for _, mx := range mxRecords {
			results = append(results, fmt.Sprintf("%v %v", mx.Host, mx.Pref))
		}
	case "TXT":
		txtRecords, err := r.LookupTXT(ctx, hostname)
		if err != nil {
			return nil, err
		}
		results = append(results, txtRecords...)
	case "NS":
		nsRecords, err := r.LookupNS(ctx, hostname)
		if err != nil {
			return nil, err
		}
		for _, ns := range nsRecords {
			results = append(results, ns.Host)
		}
	case "SRV":
		parts := strings.Split(hostname, ".")
		if len(parts) < 3 || !strings.HasPrefix(parts[0], "_") || !strings.HasPrefix(parts[1], "_") {
			return nil, fmt.Errorf("SRV lookups require hostname in format _service._proto.name")
		}
		service := strings.TrimPrefix(parts[0], "_")
		proto := strings.TrimPrefix(parts[1], "_")
		name := strings.Join(parts[2:], ".")
		_, srvs, err := r.LookupSRV(ctx, service, proto, name)
		if err != nil {
			return nil, err
		}
		for _, srv := range srvs {
			results = append(results, fmt.Sprintf("%s %d %d %d", srv.Target, srv.Port, srv.Priority, srv.Weight))
		}
	default:
		return nil, fmt.Errorf("unsupported record type: %s", recordType)
	}

	return results, err
}
