// SPDX-FileCopyrightText: 2020 Ethel Morgan
//
// SPDX-License-Identifier: MIT

package config

import (
	"encoding/json"
	"io/ioutil"
	"net"
)

type (
	Config struct {
		BrokerURI string

		MACsByTopic map[string]net.HardwareAddr
	}

	config struct {
		MQTTBroker string `json:"mqttBroker"`
		Devices    map[string]struct {
			MAC   mac    `json:"mac"`
			Topic string `json:"topic"`
		} `json:"devices"`
	}

	mac struct {
		net.HardwareAddr
	}
)

func ParseFile(path string) (*Config, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	raw := config{}
	if err := json.Unmarshal(bytes, &raw); err != nil {
		return nil, err
	}

	return configFromConfig(raw), nil
}

func configFromConfig(raw config) *Config {
	c := &Config{
		BrokerURI:   raw.MQTTBroker,
		MACsByTopic: map[string]net.HardwareAddr{},
	}

	for _, v := range raw.Devices {
		c.MACsByTopic[v.Topic] = v.MAC.HardwareAddr
	}

	return c
}

func (m mac) MarshalText() ([]byte, error) {
	return []byte(m.String()), nil
}
func (m *mac) UnmarshalText(raw []byte) error {
	mm, err := net.ParseMAC(string(raw))
	if err != nil {
		return err
	}
	m.HardwareAddr = mm
	return nil
}
