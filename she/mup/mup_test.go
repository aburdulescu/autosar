package mup

import (
	"bytes"
	"encoding/hex"
	"testing"

	"github.com/aburdulescu/autosar/she"
)

func TestEncode(t *testing.T) {
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

	m1m2m3, _, err := in.Encode()
	if err != nil {
		t.Fatal(err)
	}

	m1 := m1m2m3[:16]
	m2 := m1m2m3[16:48]
	m3 := m1m2m3[48:]

	expectedM1, _ := hex.DecodeString("00000000000000000000000000000141")
	if !bytes.Equal(m1, expectedM1) {
		t.Log("want", hex.EncodeToString(expectedM1))
		t.Log("have", hex.EncodeToString(m1))
		t.Fatal("fail")
	}

	expectedM2, _ := hex.DecodeString("2b111e2d93f486566bcbba1d7f7a9797c94643b050fc5d4d7de14cff682203c3")
	if !bytes.Equal(m2, expectedM2) {
		t.Log("want", hex.EncodeToString(expectedM2))
		t.Log("have", hex.EncodeToString(m2))
		t.Fatal("fail")
	}

	expectedM3, _ := hex.DecodeString("b9d745e5ace7d41860bc63c2b9f5bb46")
	if !bytes.Equal(m3, expectedM3) {
		t.Log("want", hex.EncodeToString(expectedM3))
		t.Log("have", hex.EncodeToString(m3))
		t.Fatal("fail")
	}

}

func TestDecode(t *testing.T) {
	m1m2m3, _ := hex.DecodeString(
		"000000000000000000000000000001412b111e2d93f486566bcbba1d7f7a9797c94643b050fc5d4d7de14cff682203c3b9d745e5ace7d41860bc63c2b9f5bb46",
	)

	in, err := Decode(m1m2m3)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(in)
}
