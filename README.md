# mcsm - Minecraft Server Manager

_Consolidate Minecraft server operations_


:warning: *This project is too early in development to be of any practical use to
anyone* :warning:

## tl;dr

## Install

*Out of date -- shouldn't really install locally on dev machines, see `just
--list` for actual tasks*

The [just](https://github.com/casey/just) CLI is used to tasks, to install run
the following:

```bash
❯ just install
go install -i -ldflags="-X main.version=loksonarius-dev -s -w"

❯ mcsm --version
loksonarius-dev
```

## Server Definition Spec

*Will be defined and documented in ref wiki pages*
*Will be needing a decent suite of default server definitions to integration
test*
