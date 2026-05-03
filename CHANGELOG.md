# Changelog

### Added

- Added the initial BugCraft Go CLI application skeleton.
- Added a Cobra-powered root command for the `bugcraft` executable.
- Added a `version` command that prints the build version, commit, and build date.
- Added Go module metadata for `github.com/dcxforge/bugcraft` and Cobra dependencies.
- Added a `Makefile` with targets for building, installing, testing, vetting, formatting, tidying, and cleaning the project.
- Expanded the README from a short pitch into a fuller game design overview covering Cronfarm, Stacksmith, Regex Rogue, and Bug Hunt.
- Added ignore rules for build artifacts, distribution output, VS Code settings, and macOS metadata.
- Added a `start` command that initializes the BugCraft home directory and prints the opening welcome.
- Added app path helpers with `BUGCRAFT_HOME` override support for config, save, and core pack locations.
- Added the default save model, JSON save load/write support, and startup creation of a default `save.json`.
- Added tests for start command initialization, app directory behavior, default save preservation, model defaults, and save round trips.
