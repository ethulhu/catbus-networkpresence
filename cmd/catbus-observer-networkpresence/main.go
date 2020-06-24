// SPDX-FileCopyrightText: 2020 Ethel Morgan
//
// SPDX-License-Identifier: MIT

// Binary catbus-observer-networkpresence detects devices on the network for Catbus.
package main

import (
	"context"
	"net"
	"time"

	"go.eth.moe/catbus"
	"go.eth.moe/catbus-networkpresence/arp"
	"go.eth.moe/catbus-networkpresence/config"
	"go.eth.moe/flag"
	"go.eth.moe/logger"
)

var (
	configPath = flag.Custom("config-path", "", "path to config file", flag.RequiredString)
	iface      = flag.Custom("interface", "", "interface to scan on", func(raw string) (interface{}, error) {
		return net.InterfaceByName(raw)
	})

	scanPeriod = flag.Duration("scan-period", 30*time.Second, "how frequently to scan ARP")
)

const retain = catbus.Retain

func main() {
	flag.Parse()

	configPath := (*configPath).(string)
	iface := (*iface).(*net.Interface)

	log, ctx := logger.FromContext(context.Background())

	config, err := config.ParseFile(configPath)
	if err != nil {
		log.AddField("config-path", configPath)
		log.WithError(err).Fatal("could not parse config file")
	}

	catbus := catbus.NewClient(config.BrokerURI, catbus.ClientOptions{})
	go func() {
		log, _ := log.Fork(ctx)
		log.AddField("broker-uri", config.BrokerURI)
		log.Info("connecting to Catbus")
		if err := catbus.Connect(); err != nil {
			log.WithError(err).Fatal("could not connect to Catbus")
		}
	}()

	for range time.Tick(*scanPeriod) {
		macs, err := arp.Scan(ctx, iface)
		if err != nil {
			log.WithError(err).Error("could not scan ARP")
			continue
		}
		log.Debug("scanned MACs")

		macPresent := map[string]bool{}
		for _, mac := range macs {
			macPresent[mac.String()] = true
		}

		for topic, mac := range config.MACsByTopic {
			power := "off"
			if macPresent[mac.String()] {
				power = "on"
			}

			if err := catbus.Publish(topic, retain, power); err != nil {
				log := log.WithError(err)
				log.AddField("topic", topic)
				log.AddField("payload", power)
				log.Error("could not publish MAC")
				continue
			}
		}
		log.Debug("published MACs")
	}
}
