# `dot`
A simple tool to save and run frequently used commands

## Overview
Save your project's commands in `dot.yaml`:
```yaml
build: go build -o ./bin/dot ./lib
```

Then run a saved command with `dot <command>`:
```shell script
$ dot build
```

## Installation
As of right now, `dot` is only officially available for macOS via [`homebrew`](https://brew.sh/). To install:
```shell script
$ brew install lukecjohnson/packages/dot
``` 
For other platforms, `dot` can be built and installed from the source.

## Feedback
`dot` is currently in early development and any feedback would be greatly appreciated. If you come across a bug or have 
an idea for a feature, please feel free to [open an issue](https://github.com/lukecjohnson/dot/issues/new).
