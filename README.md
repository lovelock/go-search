# go-search
search package by name from pkg.go.dev with cli

## Installation

`go install github.com/lovelock/go-search`

## Usage

`go-search -q viper`

```
❯ go-search -q viper                                                                      ─╯
visiting https://pkg.go.dev/search?q=viper&m=package&limit=5
viper github.com/spf13/viper

=============================

server github.com/cosmos/cosmos-sdk/server
The commands from the SDK are defined with `cobra` and configured with the `viper` package.
=============================

viper github.com/FZambia/viper-lite

=============================

viper github.com/mattermost/viper

=============================

viper github.com/sunmi-OS/gocore/viper

=============================
```
