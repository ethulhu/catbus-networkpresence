// SPDX-FileCopyrightText: 2020 Ethel Morgan
//
// SPDX-License-Identifier: MIT

// Binary arp-scan scans ARP.
package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"go.eth.moe/catbus-networkpresence/arp"
	"go.eth.moe/flag"
)

var (
	iface = flag.Custom("interface", "", "interface to scan on", func(raw string) (interface{}, error) {
		return net.InterfaceByName(raw)
	})
)

func main() {
	flag.Parse()

	ctx := context.Background()

	iface := (*iface).(*net.Interface)

	hwaddrs, err := arp.Scan(ctx, iface)
	if err != nil {
		log.Fatalf("could not scan ARP: %v", err)
	}

	for _, addr := range hwaddrs {
		fmt.Println(addr)
	}
}
