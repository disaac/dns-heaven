package dnsheaven

import (
	"github.com/miekg/dns"
)

// Resolver interface to handle resolutions
type Resolver interface {
	Resolve(net string, req *dns.Msg) (*dns.Msg, error)
}
