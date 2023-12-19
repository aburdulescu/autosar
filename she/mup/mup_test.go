package mup

import (
	"bytes"
	"encoding/hex"
	"testing"

	"github.com/aburdulescu/autosar/she"
)

func TestMUPEncode(t *testing.T) {
	WithLogs()

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

	expectedM1, _ := hex.DecodeString("00000000000000000000000000000141")
	if !bytes.Equal(result.M1[:], expectedM1) {
		t.Log("want", hex.EncodeToString(expectedM1))
		t.Log("have", hex.EncodeToString(result.M1[:]))
		t.Fatal("fail")
	}

	expectedM2, _ := hex.DecodeString("2b111e2d93f486566bcbba1d7f7a9797c94643b050fc5d4d7de14cff682203c3")
	if !bytes.Equal(result.M2[:], expectedM2) {
		t.Log("want", hex.EncodeToString(expectedM2))
		t.Log("have", hex.EncodeToString(result.M2[:]))
		t.Fatal("fail")
	}

	expectedM3, _ := hex.DecodeString("b9d745e5ace7d41860bc63c2b9f5bb46")
	if !bytes.Equal(result.M3[:], expectedM3) {
		t.Log("want", hex.EncodeToString(expectedM3))
		t.Log("have", hex.EncodeToString(result.M3[:]))
		t.Fatal("fail")
	}

}
