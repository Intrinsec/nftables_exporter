// Copyright 2020 Intrinsec
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build linux,amd64
// +build !nonftables

package collector

import (
	"runtime"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/google/nftables"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/vishvananda/netns"
)

type nftablesCollector struct {
	current typedDesc
	logger  log.Logger
}

func init() {
	registerCollector("nftables", defaultEnabled, NewNFTablesCollector)
}

// NewNFTablesCollector returns a new Collector exposing IpTables stats.
func NewNFTablesCollector(logger log.Logger) (Collector, error) {

	return &nftablesCollector{
		current: typedDesc{prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "rules_total"),
			"nftables number of rules by family, tables and chains",
			[]string{"family", "table", "chain"}, nil,
		), prometheus.CounterValue},
		logger: logger,
	}, nil
}

func tableFamilyToString(family nftables.TableFamily) string {

	switch family {
	case nftables.TableFamilyINet:
		return "inet"
	case nftables.TableFamilyIPv4:
		return "ipv4"
	case nftables.TableFamilyIPv6:
		return "ipv6"
	case nftables.TableFamilyARP:
		return "arp"
	case nftables.TableFamilyNetdev:
		return "netdev"
	case nftables.TableFamilyBridge:
		return "bridge"
	default:
		return "unknown"
	}
}

func (c *nftablesCollector) Update(ch chan<- prometheus.Metric) (err error) {

	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	ns, err := netns.Get()
	if err != nil {
		return err
	}
	defer ns.Close()

	nft := &nftables.Conn{NetNS: int(ns)}

	tables, err := nft.ListTables()
	if err != nil {
		level.Error(c.logger).Log("err", err)
		return err
	}

	chains, err := nft.ListChains()
	if err != nil {
		level.Error(c.logger).Log("err", err)
		return err
	}
	for _, table := range tables {
		for _, chain := range chains {

			rules, err := nft.GetRule(table, chain)
			if err != nil {
				level.Error(c.logger).Log("err", err)
				return err
			}
			ch <- c.current.mustNewConstMetric(
				float64(len(rules)),
				tableFamilyToString(table.Family),
				table.Name,
				chain.Name,
			)
		}
	}
	return nil
}
