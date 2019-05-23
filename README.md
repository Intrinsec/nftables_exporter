# iptables Exporter 

[![Go Report Card](https://goreportcard.com/badge/github.com/ynsta/iptables_exporter)](https://goreportcard.com/report/github.com/ynsta/iptables_exporter)

Prometheus stand alone exporter for iptables metrics.

## Building and running

### Build

    make

### Capabilities

iptables_exporter requires some additional capabilities in order to collect its metrics

    setcap CAP_NET_RAW,CAP_NET_ADMIN+ep iptables_exporter

### Running

    ./iptables_exporter <flags>

## Using Docker

You can deploy this exporter using the [ynsta/iptables-exporter](https://registry.hub.docker.com/u/ynsta/iptables-exporter/) Docker image.

For example:

```bash
FIXME
```
