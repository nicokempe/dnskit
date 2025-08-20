package dnsutils

import (
	"context"
	"net"
	"strings"
)

// getResolver returns a *net.Resolver. If resolverAddress is empty the system
// resolver is used. When provided, the address should be in the form
// host[:port]. If the port is omitted, 53 is assumed.
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
