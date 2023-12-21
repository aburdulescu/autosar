package main

import (
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/aburdulescu/autosar/she/mup"
)

func cmdEncode(args ...string) error {
	fset := flag.NewFlagSet("encode", flag.ContinueOnError)

	oneLine := fset.Bool("l", false, "Print everything on one line")

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

	if *oneLine {
		fmt.Println(hex.EncodeToString(result[:]))
	} else {
		m1, m2, m3, m4, m5 := mup.SliceMs(result)
		fmt.Println(hex.EncodeToString(m1))
		fmt.Println(hex.EncodeToString(m2))
		fmt.Println(hex.EncodeToString(m3))
		fmt.Println(hex.EncodeToString(m4))
		fmt.Println(hex.EncodeToString(m5))
	}

	return nil
}
