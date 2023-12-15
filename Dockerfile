FROM ubuntu:22.04 AS base

ENV GO_VERSION "1.20.3"
ENV GO_ARCH "amd64"
ENV PATH "$PATH:/usr/local/go/bin:/home/appuser/go/bin"

RUN apt update && apt install -y --no-install-recommends curl jq make ca-certificates
RUN curl -O -L --proto "=https" --tlsv1.2 "https://golang.org/dl/go${GO_VERSION}.linux-${GO_ARCH}.tar.gz" && \
    tar -C /usr/local -xzf go${GO_VERSION}.linux-${GO_ARCH}.tar.gz && \
    rm -v go*.tar.gz
RUN curl -O -L --proto "=https" --tlsv1.2  https://github.com/mikefarah/yq/releases/latest/download/yq_linux_amd64 && \
    mv yq_linux_amd64 /usr/bin/yq && \
    chmod 755 /usr/bin/yq
RUN useradd -m appuser

FROM base AS final

ARG BUILD_DATE
ARG VCS_REF
ARG BUILD_VERSION

# Labels.
LABEL org.label-schema.schema-version="1.0"
LABEL org.label-schema.build-date=$BUILD_DATE
LABEL org.label-schema.name="bitshifted/launchcode"
LABEL org.label-schema.description="Generates launchers for Java applications"
LABEL org.label-schema.url="https://github.com/bitshifted/launchcode"
LABEL org.label-schema.vcs-url="https://github.com/bitshifted/launchcode"
LABEL org.label-schema.vcs-ref=$VCS_REF
LABEL org.label-schema.vendor="Bitshift"
LABEL org.label-schema.version=$BUILD_VERSION

RUN mkdir -p /usr/src/launchcode

COPY ./ /usr/src/launchcode
COPY target/launch-config /usr/bin
COPY docker/build-launchers.sh /usr/bin/build-launchers

RUN rm -rvf /usr/src/launchcode/target
RUN chmod 755 /usr/bin/launch-config && chmod 755 /usr/bin/build-launchers && chown -Rv appuser:appuser /usr/src/launchcode
RUN mkdir /workspace
WORKDIR /workspace

USER appuser

ENTRYPOINT [ "/usr/bin/build-launchers" ]
