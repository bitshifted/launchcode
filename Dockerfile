FROM ubuntu:22.04 as base

ENV GO_VERSION "1.20.3"
ENV GO_ARCH "amd64"
ENV PATH "$PATH:/usr/local/go/bin:/root/go/bin"

RUN apt update && apt install -y --no-install-recommends curl jq make ca-certificates
RUN curl -O -L "https://golang.org/dl/go${GO_VERSION}.linux-${GO_ARCH}.tar.gz" && \
    tar -C /usr/local -xzf go${GO_VERSION}.linux-${GO_ARCH}.tar.gz
RUN curl -O -L  https://github.com/mikefarah/yq/releases/latest/download/yq_linux_amd64 && \
    cp yq_linux_amd64 /usr/bin/yq && \
    chmod 755 /usr/bin/yq

FROM base as final

RUN mkdir -p /usr/src/launchcode

COPY ./ /usr/src/launchcode
COPY target/launch-config /usr/bin
COPY docker/build-launchers.sh /usr/bin/build-launchers

RUN chmod 755 /usr/bin/launch-config && chmod 755 /usr/bin/build-launchers
RUN mkdir /workspace
WORKDIR /workspace

ENTRYPOINT [ "/usr/bin/build-launchers" ]
