package dnsutils

import "context"

// ReverseLookup performs a reverse DNS lookup for the given IP address using
// the optional custom resolver.
func ReverseLookup(ipAddress, resolverAddr string) ([]string, error) {
	resolver := getResolver(resolverAddr)
	return resolver.LookupAddr(context.Background(), ipAddress)
}
