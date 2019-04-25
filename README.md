
[![Godoc Reference](https://godoc.org/github.com/jwhett/gdimmer?status.svg)](http://godoc.org/github.com/jwhett/gdimmer)
[![Build Status](https://travis-ci.org/jwhett/gdimmer.svg?branch=master)](https://travis-ci.org/jwhett/gdimmer)
[![Go Report Card](https://goreportcard.com/badge/github.com/jwhett/gdimmer)](https://goreportcard.com/report/github.com/jwhett/gdimmer)


`gdimmer` - Backlight dimmer written in Go


## Description

`gdimmer` is a substitute for the built-in screen dimming utilities where these
utilities fail to handle custom or unusual configurations.


## Environment

You may choose to set the `BLPROVIDER` environment variable instead of using `-p`  or `-provider`.


## Options

`-h`, `-help` - Display help...

`-p`, `-provider` - Set your backlight provider. Providers are listed as directories
in `/sys/class/backlight/`. Can be full path or basename. This is `gmux_backlight` by default.

`-u`, `-up` - Increase the brightness. Default step is 10%.

`-d`, `-down`- Decrease the brightness. Default step is 10%.

`-m`, `-max` - Maximum brightness.

`-s`, `-set`- Explicitly set the brightness. Must be used with `-f` or `-force`.
