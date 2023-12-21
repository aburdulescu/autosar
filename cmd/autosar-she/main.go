package main

import (
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/aburdulescu/autosar/she/mup"
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

func cmdExample(args ...string) error {
	fset := flag.NewFlagSet("example", flag.ContinueOnError)
	if err := fset.Parse(args); err != nil {
		return err
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")

	input := mup.Input{
		UID:     "000000000000000000000000000000",
		ID:      0,
		AuthID:  0,
		AuthKey: "00000000000000000000000000000000",
		NewKey:  "00000000000000000000000000000000",
	}

	return enc.Encode(input)
}

func cmdEncode(args ...string) error {
	fset := flag.NewFlagSet("encode", flag.ContinueOnError)

	if err := fset.Parse(args); err != nil {
		return err
	}

	inFile := os.Stdin
	if fset.NArg() > 0 {
		f, err := os.Open(fset.Arg(0))
		if err != nil {
			return err
		}
		defer f.Close()
		inFile = f
	}

	var input mup.Input
	if err := json.NewDecoder(inFile).Decode(&input); err != nil {
		return err
	}

	result, err := input.Encode()
	if err != nil {
		return err
	}

	fmt.Println(hex.EncodeToString(result))

	return nil
}
