# `run`
Save and run frequently used shell commands

## Usage
Save your project's commands in `run.yaml`:
```yaml
test: go test .
build: go build -ldflags "-s -w" -o bin/example
package: tar -czf dist/example.tar.gz bin LICENSE
```

Then execute a saved command with `run [command]`:
```
$ run build
```

Multiple commands can be chained together with `run [command] ...`:
```
$ run test build package
```

A `default` command can be specified in `run.yaml` that will be executed with the base `run` command:
```yaml
default: run test build package
```
```
$ run
```

## Installation
Currently, `run` is only officially distributed for macOS via [Homebrew](https://brew.sh/). To install:
```
$ brew install lukecjohnson/packages/run
```

For other platforms, prebuilt binaries can be downloaded directly from the [releases page](https://github.com/lukecjohnson/run/releases).
