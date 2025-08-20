package dnsutils

import (
	"context"
)

// ReverseLookup performs a reverse DNS lookup for the given IP address using
// the optional custom resolver.
func ReverseLookup(ip, resolverAddr string) ([]string, error) {
	r := getResolver(resolverAddr)
	return r.LookupAddr(context.Background(), ip)
}
