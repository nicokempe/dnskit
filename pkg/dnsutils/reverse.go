package dnsutils

import "context"

/*
ReverseLookup performs a reverse DNS lookup.

	@param ipAddress IP address to resolve.
	@param resolverAddr optional custom resolver.
	@returns hostnames or an error.
*/
func ReverseLookup(ipAddress, resolverAddr string) ([]string, error) {
	resolver := getResolver(resolverAddr)
	return resolver.LookupAddr(context.Background(), ipAddress)
}
