# nftables Exporter 

[![Go Report Card](https://goreportcard.com/badge/github.com/intrinsec/nftables_exporter)](https://goreportcard.com/report/github.com/intrinsec/nftables_exporter)

Prometheus stand alone exporter for nftables metrics.

## Building and running

### Build

    make

### Capabilities

`nftables_exporter` does not depends on `libnftnl` or `nft` binary but requires some 
additional capabilities in order to collect its metrics

    setcap CAP_NET_ADMIN,CAP_SYS_ADMIN+ep nftables_exporter+ep nftables_exporter

### Running

    ./nftables_exporter <flags>
