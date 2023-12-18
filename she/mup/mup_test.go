package mup

import (
	"encoding/hex"
	"testing"

	"github.com/aburdulescu/autosar/she"
)

func TestMUPEncode(t *testing.T) {
	in := Input{
		UID:     "000000000000000000000000000001",
		AuthID:  she.MASTER_ECU_KEY,
		ID:      she.KEY_1,
		AuthKey: "000102030405060708090a0b0c0d0e0f",
		NewKey:  "0f0e0d0c0b0a09080706050403020100",
		Counter: 1,
		Flags: ProtectionFlags{
			KeyUsage: false,
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
