// SPDX-FileCopyrightText: 2020 Ethel Morgan
//
// SPDX-License-Identifier: MIT

package arp

import (
	"bytes"
	"context"
	"fmt"
	"net"
	"os/exec"
	"regexp"
)

var (
	ipAndMac = regexp.MustCompile(`^[[:digit:].]+[[:space:]]+([[:xdigit:]:]+)$`)
)

func Scan(ctx context.Context, iface *net.Interface) ([]net.HardwareAddr, error) {
	cmd := exec.CommandContext(ctx, "arp-scan", "--localnet", "--quiet", fmt.Sprintf("--interface=%s", iface.Name))

	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("could not run arp-scan: %w", err)
	}

	return parseOutput(output)
}

func parseOutput(output []byte) ([]net.HardwareAddr, error) {
	var macs []net.HardwareAddr
	for _, line := range bytes.Split(output, []byte("\n")) {
		submatches := ipAndMac.FindSubmatch(line)
		if len(submatches) == 0 {
			continue
		}

		mac, err := net.ParseMAC(string(submatches[1]))
		if err != nil {
			return macs, err
		}
		macs = append(macs, mac)
	}
	return macs, nil
}
