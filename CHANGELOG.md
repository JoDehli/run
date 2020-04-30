# Changelog

## `1.0.0-beta.4` (2020-04-30)

### Features
- Shorthand flags
    - `-v`, `--version`
    - `-l`, `--list`

### Improvements
- Implemented `pflag` package
- Utilizing `PrintDefaults` instead of custom usage handling

## `1.0.0-beta.3` (2020-04-30)

### Breaking changes
- Switched to JSON for commands file
    - `run.yaml` ⟶ `run.json`

### Code refactoring
- Simplifying project structure - moving everything to `main` package

## `1.0.0-beta.2` (2020-04-26)

### Fixes
- Ensure `run.yaml` exists before attempting to read
- Remove commit hash from version number generated from git tag
- Empty `run` command outputs flag usage

### Improvements
- "run" added to `--version` output
- Cleaner formatting for `--help` and `--list` output

## `1.0.0-beta.1` (2020-04-25)
Initial release 🎉