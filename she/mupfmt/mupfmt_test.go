package mupfmt

import (
	"encoding/hex"
	"testing"
)

func TestEncode(t *testing.T) {
	in := Input{
		UID:       "00",
		AuthKeyID: 1,
		AuthKey:   "00",
		ID:        4,
		NewKey:    "00",
		Counter:   1,
		Flags: ProtectionFlags{
			Write:    true,
			Boot:     true,
			Debugger: true,
			Wildcard: true,
		},
	}

	result, err := in.Encode()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(hex.EncodeToString(result.M1[:]))
	t.Log(hex.EncodeToString(result.M2[:]))
	t.Log(hex.EncodeToString(result.M3[:]))
}
