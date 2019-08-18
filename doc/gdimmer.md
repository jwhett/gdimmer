% GDIMMER(1)
% Josh Whetton
% August 2018


# NAME

**gdimmer** - Backlight dimmer written in Go


# SYNOPSIS

**gdimmer** [**-p** *PROVIDER*|**-provider**=*PROVIDER*] [**-u**|**-up**] | [**-d**|**-down**] | [**-m**|**-max**] | [**-f**|**-force**] [**-s** *VALUE* | **-set** *VALUE*]


# DESCRIPTION

**gdimmer** is a substitute for the built-in screen dimming utilities where these
utilities fail to handle custom or unusual configurations.


# ENVIRONMENT

You may choose to set the **BLPROVIDER** environment variable instead of using *-p* or *-provider*.


# OPTIONS

**-h**, **-help**
:	Display help...

**-p**, **-provider**
:	Set your backlight provider. Providers are listed as directories in /sys/class/backlight/. Can be full path or basename. This is the first provider found in /sys/class/backlight by default.

**-u**, **-up**
:	Increase the brightness. Default step is 10%.

**-d**, **-down**
:	Decrease the brightness. Default step is 10%.

**-m**, **-max**
:	Maximum brightness.

**-s**, **-set**
:	Explicitly set the brightness. Must be used with **-f** or **-force**.


# EXAMPLES

**gdimmer** **-u**
:	Increase the brightness.

**gdimmer** **-d**
:	Decrease the brightness.

**gdimmer** **-m**
:	Max the brightness.

**gdimmer** **-f** **-s** *0*
:	Set the brightness to zero.

**gdimmer** **-p** intel_backlight **-f** **-s** *0*
:	Set the brightness to zero.

# LICENSE

This is free and unencumbered software released into the public domain.

For more information, please refer to <http://unlicense.org>

# SEE ALSO

*godoc.org/github.com/jwhett/gdimmer*
