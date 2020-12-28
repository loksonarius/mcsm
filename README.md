# mcsm - Minecraft Server Manager

_Consolidate Minecraft server operations_


:warning: *This project is too early in development to be of any practical use to
anyone* :warning:

## tl;dr

## Install

The [just](https://github.com/casey/just) CLI is used to tasks, to install run
the following:

```bash
❯ just install
go install -i -ldflags="-X main.version=loksonarius-dev -s -w"

❯ mcsm --version
loksonarius-dev
```

## Server Definition Spec

```yaml
install:
  type: (forge, vanilla, spigot, bedrock)
  version: <X.Y.Z>
  mods: [ source: (<URL>|<FS PATH>) ]
  plugins: [ source: (<URL>|<FS PATH>) ]

run:
  initial_memory: <Mem>
  max_memory: <Mem>
  debug_gc: <Bool>

config:
  vanilla:
  spigot:
  bukkit:
  bedrock:
  forge:
```
