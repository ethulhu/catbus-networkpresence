// SPDX-FileCopyrightText: 2020 Ethel Morgan
//
// SPDX-License-Identifier: MIT

package arp

import (
	"net"
	"reflect"
	"testing"
)

const output = `
Interface: wlp2s0, type: EN10MB, MAC: 18:3d:a2:1a:f1:14, IPv4: 192.168.69.195
Starting arp-scan 1.9.7 with 256 hosts (https://github.com/royhills/arp-scan)
192.168.16.1    ad:d3:8f:73:cf:c6
192.168.16.12   dd:5f:f4:ed:7a:1e
192.168.16.128  dc:93:32:81:e7:cd

15 packets received by filter, 0 packets dropped by kernel
Ending arp-scan 1.9.7: 256 hosts scanned in 1.799 seconds (142.30 hosts/sec). 15 responded
`

func TestParseOutput(t *testing.T) {
	want := []net.HardwareAddr{
		{0xad, 0xd3, 0x8f, 0x73, 0xcf, 0xc6},
		{0xdd, 0x5f, 0xf4, 0xed, 0x7a, 0x1e},
		{0xdc, 0x93, 0x32, 0x81, 0xe7, 0xcd},
	}

	got, err := parseOutput([]byte(output))
	if err != nil {
		t.Fatalf("got error: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
