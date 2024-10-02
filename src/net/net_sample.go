package net

import (
	dns "github.com/dictxwang/benburkert-dns"
	"log"
	"net"
	"time"
)

func initCustomDns() {
	zone := &dns.Zone{
		Origin: "example.org.",
		TTL:    5 * time.Minute,
		RRs: dns.RRSet{
			"foo": {
				dns.TypeA: []dns.Record{
					&dns.A{A: net.ParseIP("1.2.3.4")},
				},
			},
		},
	}
	mux := new(dns.ResolveMux)
	mux.Handle(dns.TypeANY, zone.Origin, zone)

	net.DefaultResolver = &net.Resolver{
		PreferGo: true,
		Dial: (&dns.Client{
			Resolver: mux,
		}).Dial,
	}
}

func NewMain() {
	log.Println(net.LookupHost("foo.example.org")) // it's working. output: [1.2.3.4] <nil>
	log.Println(net.LookupHost("www.example.org")) // error no such host. I don't known how to fallback to DNS query
}
