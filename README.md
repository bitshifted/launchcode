[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=bitshifted_launchcode&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=bitshifted_launchcode) 
[![Publish Docker Image](https://github.com/bitshifted/launchcode/actions/workflows/publish-docker-image.yml/badge.svg?branch=master)](https://github.com/bitshifted/launchcode/actions/workflows/publish-docker-image.yml) 
[![License: MPL 2.0](https://img.shields.io/badge/License-MPL_2.0-brightgreen.svg)](https://opensource.org/licenses/MPL-2.0)

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

# Usage

Launchcode is packaged as Docker image and can be run on any platform where DOcker is supported. To generate launchers for all platforms, run the following command:

```bash
docker run -v ${PWD}:/workspace  ghcr.io/bitshifted/launchcode:<version> config-file.yaml
```

This command mounts the current directory into `/workspace` directory in container. Argument `config-file.yaml` is a configuration file for installers generation. It is assumed that configuration file is in a present directory.

Generated installers will be placed in `dist` directory inside the current directory. Contents of this directory will look like this:

```bash
dist
  |- launcher-linux-x64
  |- launcher-linux-aarch64
  |- launcher-mac-x64
  |- launcher-mac-aarch64
  |- launcher-win-x64.exe
```

## Configuration file reference

This section outlines the available options for configuration. Configuration file is in YAML format. Example file is shown bellow:

```yaml
common:
  jre-dir: runtime
  restart-code: 20
  jvm-options: -Xms=20m  -Xmx200m
  jvm-system-properties: |
    -Dsome.property=1 -Dfoo.bar=baz \
    -Danother.property=x
  splash-screen: splash-screen.png
  classpath: classpath/dir/*
  module-path: some-directory
  add-modules: |
    javafx.core javafx.graphics \
    javafx.fxml
  module: mymodule
  main-class: com.test.MainClass
linux:
  splash-screen: splash-linux.png
  jvm-options: -Xms=10m  -Xmx=300m
mac:
  jre-dir: Contents/runtime/jre
  module-path: Contents/runtime/modules
windows:
  icon: test-icon.ico
  jvm-system-properties: |
    -Dwindows.foo=1 -Dwindows.bar=baz
```

There are 4 top level sections, which basically have the same options.
* `common` - options for all platforms, unless overriden for specific platforms
* `linux` - overrides common options on Linux
* `mac` - overrides common options on Mac OS
* `windows` - overrides common options on WIndows

Configuration options:
* `jre-dir` - location of JRE directory. It is assumes that path is relative to application root directory (on Mac OS, relative to bundle root)
* `restart-code` - return code from Java application that will cause a restart. This is useful if you need to eg. restart after update
* `jvm-options` - options to pass to Java VM
* `jvm-system-properties` - system properties to pass to JVM
* `splash-screen` - splash screen option to pass to JVM (ie. `--splash:image.png`)
* `classpath` - application classpath
* `module-path` - application module path
* `add-modules` - list of modules to add to application
* `module` - module to launch. This options requires `main-class` to be specified, so it can ba passed in form `mymodule/MyMainClass`
* `main-class` - application main class
* `jar` - application JAR file
* `arguments` - list of arguments to pass to application
* `icon` - icon to embed in the launcher (applicable only on Windows)
