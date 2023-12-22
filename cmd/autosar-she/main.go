package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/debug"
)

func main() {
	if err := mainErr(os.Args[1:]...); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func mainErr(args ...string) error {
	fset := flag.NewFlagSet("main", flag.ContinueOnError)

	printVersion := fset.Bool("version", false, "Print version information")

	if err := fset.Parse(args); err != nil {
		return err
	}

	if *printVersion {
		version()
		return nil
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
	case "decode":
		return cmdDecode(args...)
	case "example":
		return cmdExample(args...)
	default:
		return fmt.Errorf("unknown command '%s'", cmd)
	}
}

func version() {
	bi, _ := debug.ReadBuildInfo()
	g := func(k string) string {
		for _, v := range bi.Settings {
			if v.Key == k {
				return v.Value
			}
		}
		return ""
	}
	fmt.Println("go     ", bi.GoVersion)
	fmt.Println("commit ", g("vcs.revision"))
	fmt.Println("time   ", g("vcs.time"))
	fmt.Println("dirty  ", g("vcs.modified"))
}
