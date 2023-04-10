[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=bitshifted_launchcode&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=bitshifted_launchcode) 
[![Publish Docker Image](https://github.com/bitshifted/launchcode/actions/workflows/publish-docker-image.yml/badge.svg?branch=master)](https://github.com/bitshifted/launchcode/actions/workflows/publish-docker-image.yml)

# Launchcode

Launchcode is a native application launcher for Java based applications. It provides integration with target 
operating system and gives more native look and feel to Java-based GUI applications.

Launchcode provides binaries for the following operating systems and CPU architectures:
* Windows on Intel/AMD 64-bit CPUs
* Mac on Intel and ARM 64-bit CPUs
* Linux on Intel/AMD and ARM 64-bit CPUs

For Windows launchers, it is possible to embed icon into executable.

# License

Backstage is released under Mozilla Public License 2.0. See [LICENSE](./LICENSE) file for details.

# Building and running

Launchode is written in Go. Requirements for the build:

* Go 1.20 or higher
* make

To build the binary for all supported platforms, run `make all`
 command.