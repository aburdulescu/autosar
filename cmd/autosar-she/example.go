package main

import (
	"encoding/json"
	"flag"
	"os"

	"github.com/aburdulescu/autosar/she/mup"
)

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
