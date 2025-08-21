package dnsutils

import (
	"context"
	"net"
	"strings"
)

/*
getResolver creates a DNS resolver.

	@param resolverAddress custom resolver in form host[:port]; empty uses system resolver.
	@returns configured *net.Resolver.
*/
func getResolver(resolverAddress string) *net.Resolver {
	if resolverAddress == "" {
		return net.DefaultResolver
	}
	if !strings.Contains(resolverAddress, ":") {
		resolverAddress = resolverAddress + ":53"
	}
	return &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, _ string) (net.Conn, error) {
			dialer := net.Dialer{}
			return dialer.DialContext(ctx, network, resolverAddress)
		},
	}
}
