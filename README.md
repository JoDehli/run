# `run`
A simple tool to save and run frequently used commands

## Overview
Save your project's commands in `run.yaml`:
```yaml
build: go build -o ./bin/run ./lib
```

Then run a saved command with `run <command>`:
```shell script
$ run build
```

## Installation
As of right now, `run` is only officially available for macOS via [`homebrew`](https://brew.sh/). To install:
```shell script
$ brew install lukecjohnson/packages/run
``` 
For other platforms, `run` can be built and installed from the source.

## Feedback
`run` is currently in early development and any feedback would be greatly appreciated. If you come across a bug or have 
an idea for a feature, please feel free to [open an issue](https://github.com/lukecjohnson/run/issues/new).
