FROM golang:1.19.5-buster

RUN mkdir -p /usr/src/launchcode
COPY ./ /usr/src/launchcode
COPY docker/copy-src.sh /usr/local/bin/copy-src
COPY docker/build-launchers.sh /usr/local/bin/build-launchers
