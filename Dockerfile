ARG ARCH="amd64"
ARG OS="linux"
FROM quay.io/prometheus/busybox-${OS}-${ARCH}:glibc
LABEL maintainer="Stany MARCEL <stanypub@gmail.com>"

ARG ARCH="amd64"
ARG OS="linux"
COPY .build/${OS}-${ARCH}/nftables_exporter /bin/nftables_exporter

EXPOSE      9732
USER        nobody
ENTRYPOINT  [ "/bin/nftables_exporter" ]
