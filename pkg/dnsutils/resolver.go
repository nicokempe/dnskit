package dnsutils

import (
	"context"
	"net"
	"strings"
)

// getResolver returns a *net.Resolver. If addr is empty the system resolver is
// used. When provided, addr should be in the form host[:port]. If the port is
// omitted, 53 is assumed.
func getResolver(addr string) *net.Resolver {
	if addr == "" {
		return net.DefaultResolver
	}
	if !strings.Contains(addr, ":") {
		addr = addr + ":53"
	}
	return &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, _ string) (net.Conn, error) {
			d := net.Dialer{}
			return d.DialContext(ctx, network, addr)
		},
	}
}
