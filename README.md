<!--
SPDX-FileCopyrightText: 2020 Ethel Morgan

SPDX-License-Identifier: CC0-1.0
-->

# Catbus Network-Presence

A daemon to detect devices on a network and update [Catbus](https://ethulhu.co.uk/catbus) accordingly.

## Config

```json
{
  "mqttBroker": "tcp://broker.local:1883",
  "devices": {
    "TV": {
      "mac": "aa:bb:cc:dd:ee:ff",
      "topic": "home/living-room/tv/power"
    }
  }
}
```

## Methods

### ARP scan

The daemon (and `cmd/arp-scan` tool) wraps [`arp-scan`](https://linux.die.net/man/1/arp-scan) to detect devices that exist on the network, periodically polling.

### Ping

TODO: ping devices that showed up on ARP scanning, to poll more frequently, to detect shutdowns faster.

### DHCP sniff

TODO: implement a small DHCP server to pick up connects faster.
