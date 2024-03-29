package main

import (
	"flag"
	"fmt"
	"io"

	"github.com/weppos/publicsuffix-go/publicsuffix"
)

// Exit codes are int values that represent an exit code for a particular error.
const (
	ExitCodeOK    int = 0
	ExitCodeError int = 1 + iota
)

// CLI is the command line object
type CLI struct {
	// outStream and errStream are the stdout and stderr
	// to write message from the CLI.
	outStream, errStream io.Writer
}

// Run invokes the CLI with the given arguments.
func (cli *CLI) Run(args []string) int {
	var (
		version bool
	)

	// Define option flag parse
	flags := flag.NewFlagSet(Name, flag.ContinueOnError)
	flags.SetOutput(cli.errStream)

	flags.BoolVar(&version, "version", false, "Print version information and quit.")

	// Parse commandline flag
	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeError
	}

	// Show version
	if version {
		fmt.Fprintf(cli.errStream, "%s version %s\n", Name, Version)
		return ExitCodeOK
	}

	parsedArgs := flags.Args()
	if len(parsedArgs) != 1 {
		fmt.Fprintln(cli.errStream, "Invalid argument: Specify domain.")
		return ExitCodeError
	}

	argDomain := parsedArgs[0]

	nakedDomain, err := publicsuffix.Domain(argDomain)
	if err != nil {
		return ExitCodeError
	}

	fmt.Println(nakedDomain)

	return ExitCodeOK
}
