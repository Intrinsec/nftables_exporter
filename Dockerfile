ARG ARCH="amd64"
ARG OS="linux"
FROM quay.io/prometheus/busybox-${OS}-${ARCH}:glibc
LABEL maintainer="Stany MARCEL <stanypub@gmail.com>"

ARG ARCH="amd64"
ARG OS="linux"
COPY .build/${OS}-${ARCH}/iptables_exporter /bin/iptables_exporter

EXPOSE      9732
USER        nobody
ENTRYPOINT  [ "/bin/iptables_exporter" ]
