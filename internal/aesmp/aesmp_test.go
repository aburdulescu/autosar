package aesmp

import (
	"bytes"
	"encoding/hex"
	"testing"
)

func TestCompressSHE(t *testing.T) {
	// test vectors from SHE doc
	key := decodeHex(t, []byte("000102030405060708090a0b0c0d0e0f"))
	c1 := decodeHex(t, []byte("010153484500800000000000000000b0"))
	c2 := decodeHex(t, []byte("010253484500800000000000000000b0"))
	expectedK1 := decodeHex(t, []byte("118a46447a770d87828a69c222e2d17e"))
	expectedK2 := decodeHex(t, []byte("2ebb2a3da62dbd64b18ba6493e9fbe22"))
	k1, err := Compress(key, c1)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(expectedK1, k1) {
		t.Fatal("k1 nok")
	}
	k2, err := Compress(key, c2)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(expectedK2, k2) {
		t.Fatal("k2 nok")
	}
}

func TestCompressAES256(t *testing.T) {
	key := decodeHex(t, []byte("000102030405060708090a0b0c0d0e0f000102030405060708090a0b0c0d0e0f"))
	c := decodeHex(t, []byte("010001455854008000000000000000B8"))
	out, err := Compress(key, c)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(hex.EncodeToString(out))
}

func decodeHex(t *testing.T, in []byte) []byte {
	t.Helper()
	r, err := hex.DecodeString(string(in))
	if err != nil {
		t.Fatalf("decodeHex %s: %v", in, err)
	}
	return r
}
