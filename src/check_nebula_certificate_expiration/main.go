package main

import (
	"durafmt"

	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	var certFile = flag.String("certificate", "", "Certificate file to check")
	var critStr = flag.String("critical", defaultCriticalTime, "Report critical status if certificate expires in <time> or less")
	var warnStr = flag.String("warning", defaultWarningTime, "Report warning status if certificate expires in <time> or less")
	var showVers = flag.Bool("version", false, "Show version information")
	var help = flag.Bool("help", false, "Show help text")
	var now = time.Now()
	var nowSec = now.Unix()
	var critSec float64
	var warnSec float64
	var notafterSec int64

	flag.Usage = showUsage

	flag.Parse()

	if len(flag.Args()) > 0 {
		fmt.Fprintf(os.Stderr, "Error: Trailing arguments\n")
		showUsage()
		os.Exit(UNKNOWN)
	}

	if *help {
		showUsage()
		os.Exit(OK)
	}

	if *showVers {
		showVersion()
		os.Exit(OK)
	}

	if *certFile == "" {
		fmt.Fprintf(os.Stderr, "Error: Certificate file is missing\n")
		showUsage()
		os.Exit(UNKNOWN)
	}

	crit, err := durafmt.ParseString(*critStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		os.Exit(UNKNOWN)
	}

	warn, err := durafmt.ParseString(*warnStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		os.Exit(UNKNOWN)
	}

	warnSec = warn.Duration().Seconds()
	critSec = crit.Duration().Seconds()

	// basic sanity checks
	if critSec > warnSec {
		fmt.Fprintf(os.Stderr, "Error: Warning time must be greater or equal than critical time\n")
		os.Exit(UNKNOWN)
	}

	if critSec <= 0 || warnSec <= 0 {
		fmt.Fprintf(os.Stderr, "Error: Warning / critical thresholds must be positive values\n")
		os.Exit(UNKNOWN)
	}

	raw, err := readCertificateFile(*certFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		os.Exit(UNKNOWN)
	}

	nebcert, err := parseNebulaCertificate(raw)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		os.Exit(UNKNOWN)
	}

	if now.Before(nebcert.Details.NotBefore) {
		fmt.Printf("WARNING - Certificate is not valid yet. Validity starts at %s (in %s)\n", nebcert.Details.NotBefore.String(), durafmt.Parse(now.Sub(nebcert.Details.NotBefore)).String())
		os.Exit(WARNING)
	}

	if now.After(nebcert.Details.NotAfter) {
		fmt.Printf("CRITICAL - Certificate has been expired %s ago on %s\n", durafmt.Parse(nebcert.Details.NotAfter.Sub(now)).String(), nebcert.Details.NotAfter.String())
		os.Exit(CRITICAL)
	}

	notafterSec = nebcert.Details.NotAfter.Unix()
	if float64(notafterSec-nowSec) <= critSec {
		fmt.Printf("CRITICAL - Certificate will expire in %s on %s\n", durafmt.Parse(nebcert.Details.NotAfter.Sub(now)).String(), nebcert.Details.NotAfter.String())
		os.Exit(CRITICAL)
	} else if float64(notafterSec-nowSec) <= warnSec {
		fmt.Printf("WARNING - Certificate will expire in %s on %s\n", durafmt.Parse(nebcert.Details.NotAfter.Sub(now)).String(), nebcert.Details.NotAfter.String())
		os.Exit(WARNING)
	}

	fmt.Printf("OK - Certificate will expire in %s on %s\n", durafmt.Parse(nebcert.Details.NotAfter.Sub(now)).String(), nebcert.Details.NotAfter.String())
	os.Exit(OK)
}
