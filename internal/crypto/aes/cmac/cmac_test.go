package cmac

import (
	"bytes"
	"crypto/aes"
	"encoding/hex"
	"strconv"
	"testing"
)

func h2b(h string) []byte {
	b, err := hex.DecodeString(h)
	if err != nil {
		panic(err)
	}
	return b
}

var key = h2b("2b7e151628aed2a6abf7158809cf4f3c")

type testVector struct {
	msg []byte
	mac []byte
}

var testVectors = []testVector{
	{
		msg: nil,
		mac: h2b("bb1d6929e95937287fa37d129b756746"),
	},
	{
		msg: h2b("6bc1bee22e409f96e93d7e117393172a"),
		mac: h2b("070a16b46b4d4144f79bdd9dd04a287c"),
	},
	{
		msg: h2b("6bc1bee22e409f96e93d7e117393172a" +
			"ae2d8a571e03ac9c9eb76fac45af8e51" +
			"30c81c46a35ce411"),
		mac: h2b("dfa66747de9ae63030ca32611497c827"),
	},
	{
		msg: h2b("6bc1bee22e409f96e93d7e117393172a" +
			"ae2d8a571e03ac9c9eb76fac45af8e51" +
			"30c81c46a35ce411e5fbc1191a0a52ef" +
			"f69f2445df4f9b17ad2b417be66c3710"),
		mac: h2b("51f0bebf7e3b9d92fc49741779363cfe"),
	},
}

func TestGenerate(t *testing.T) {
	for _, tv := range testVectors {
		t.Run(strconv.Itoa(len(tv.msg)), func(t *testing.T) {
			mac, err := Generate(key, tv.msg)
			if err != nil {
				t.Error(err)
			}
			assertBytes(t, mac, tv.mac)
		})
	}
}

func TestSubKey(t *testing.T) {
	expectedK1 := h2b("fbeed618357133667c85e08f7236a8de")
	expectedK2 := h2b("f7ddac306ae266ccf90bc11ee46d513b")

	c, err := aes.NewCipher(key)
	if err != nil {
		t.Fatal(err)
	}

	var k1, k2 [bs]byte

	generateSubKey(c, k1[:], k2[:])

	assertBytes(t, k1[:], expectedK1)
	assertBytes(t, k2[:], expectedK2)
}

func assertBytes(t *testing.T, have, want []byte) {
	t.Helper()
	if !bytes.Equal(have, want) {
		t.Log("want:", hex.EncodeToString(want))
		t.Log("have:", hex.EncodeToString(have))
		t.Fail()
	}
}
