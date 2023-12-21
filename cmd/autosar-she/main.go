package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	if err := mainErr(os.Args[1:]...); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func mainErr(args ...string) error {
	fset := flag.NewFlagSet("main", flag.ContinueOnError)
	if err := fset.Parse(args); err != nil {
		return err
	}

	if fset.NArg() < 1 {
		fset.Usage()
		return fmt.Errorf("need command")
	}

	cmd := fset.Arg(0)
	args = fset.Args()[1:]

	switch cmd {
	case "encode":
		return cmdEncode(args...)
	case "example":
		return cmdExample(args...)
	default:
		return fmt.Errorf("unknown command '%s'", cmd)
	}
}
