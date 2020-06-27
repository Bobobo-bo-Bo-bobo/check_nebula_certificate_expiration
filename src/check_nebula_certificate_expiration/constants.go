package main

const name = "check_nebula_certificate_expiration"
const version = "1.0.0"

const (
	// OK - Nagios state OK
	OK int = iota
	// WARNING - Nagios state WARNING
	WARNING
	// CRITICAL - Nagios state CRITICAL
	CRITICAL
	// UNKNOWN - Nagios state UNKNOWN
	UNKNOWN
)

const defaultCriticalTime = "168h"
const defaultWarningTime = "336h"

const versionText = `%s version %s
Copyright (C) 2020 by Andreas Maus <maus@ypbind.de>
This program comes with ABSOLUTELY NO WARRANTY.

%s is distributed under the Terms of the GNU General
Public License Version 3. (http://www.gnu.org/copyleft/gpl.html)

Build with go version: %s
`

const helpText = `Usage: %s --certificate=<file> [--critical=<time>] [--warning=<time>] [--help] [--version]
    --help                  Show help text

    --certificate=<file>    Certificate file to check

    --critical=<time>       Report critical status if certificate expires in <time> or less
                            Default: %s

    --version               Show version information

    --warning=<time>        Report warning status if certificate expires in <time> or less
                            Default: %s

`
