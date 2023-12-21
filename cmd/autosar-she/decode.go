package main

import (
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/aburdulescu/autosar/she/mup"
)

func cmdDecode(args ...string) error {
	fset := flag.NewFlagSet("decode", flag.ContinueOnError)

	keyHex := fset.String("key", "", "Secret key as hex")

	if err := fset.Parse(args); err != nil {
		return err
	}

	input := fset.Arg(0)
	if fset.NArg() == 0 {
		data, err := io.ReadAll(os.Stdin)
		if err != nil {
			return fmt.Errorf("failed to read stdin: %w", err)
		}
		input = strings.Trim(string(data), " \t\r\n")
	}

	key, err := hex.DecodeString(*keyHex)
	if err != nil {
		return fmt.Errorf("failed to decode key: %w", err)
	}
	if len(key) != 16 {
		return fmt.Errorf("invalid key size: %d", len(key))
	}

	m1m2m3, err := hex.DecodeString(input)
	if err != nil {
		return fmt.Errorf("failed to decode input: %w", err)
	}

	result, err := mup.Decode(m1m2m3, key)
	if err != nil {
		return err
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")

	return enc.Encode(result)
}
