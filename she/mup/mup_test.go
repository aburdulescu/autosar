package mup

import (
	"bytes"
	"encoding/hex"
	"fmt"
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
		"00000000000000000000000000000141" +
			"2b111e2d93f486566bcbba1d7f7a9797c94643b050fc5d4d7de14cff682203c3" +
			"b9d745e5ace7d41860bc63c2b9f5bb46",
	)

	authKey, _ := hex.DecodeString("000102030405060708090a0b0c0d0e0f")

	in, err := Decode(m1m2m3, authKey)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(in)

	expectedIn := Input{
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

	if err := in.equals(expectedIn); err != nil {
		t.Fatal(err)
	}
}

func TestEncodeCounterAndFlags(t *testing.T) {
	t.Run("CounterOverMax", func(t *testing.T) {
		_, err := encodeCounterAndFlags(counterMax+1, 0)
		if err == nil {
			t.Fatalf("expected error")
		}
	})

	t.Run("All1", func(t *testing.T) {
		b, err := encodeCounterAndFlags(0x0fffffff, 0x1f)
		if err != nil {
			t.Fatal(err)
		}
		expected := []byte{0xff, 0xff, 0xff, 0xff, 0x80}
		if !bytes.Equal(b[:], expected) {
			t.Log("want:", hex.EncodeToString(expected))
			t.Log("have:", hex.EncodeToString(b[:]))
			t.Fatal("fail")
		}
	})

	t.Run("All0", func(t *testing.T) {
		b, err := encodeCounterAndFlags(0, 0)
		if err != nil {
			t.Fatal(err)
		}
		expected := []byte{0, 0, 0, 0, 0}
		if !bytes.Equal(b[:], expected) {
			t.Log("want:", hex.EncodeToString(expected))
			t.Log("have:", hex.EncodeToString(b[:]))
			t.Fatal("fail")
		}
	})

}

func (in Input) equals(other Input) error {
	if in.UID != other.UID {
		return fmt.Errorf("UID: %q != %q", in.UID, other.UID)
	}
	if in.AuthID != other.AuthID {
		return fmt.Errorf("AuthID: %q != %q", in.AuthID, other.AuthID)
	}
	if in.ID != other.ID {
		return fmt.Errorf("ID: %q != %q", in.ID, other.ID)
	}
	if in.AuthKey != other.AuthKey {
		return fmt.Errorf("AuthKey: %q != %q", in.AuthKey, other.AuthKey)
	}
	if in.NewKey != other.NewKey {
		return fmt.Errorf("NewKey: %q != %q", in.NewKey, other.NewKey)
	}
	if in.Counter != other.Counter {
		return fmt.Errorf("Counter: %d != %d", in.Counter, other.Counter)
	}
	if err := in.Flags.equals(other.Flags); err != nil {
		return fmt.Errorf("Flags: %w", err)
	}
	return nil
}

func (f ProtectionFlags) equals(other ProtectionFlags) error {
	if f.Write != other.Write {
		return fmt.Errorf("Write: %v != %v", f.Write, other.Write)
	}
	if f.Boot != other.Boot {
		return fmt.Errorf("Boot: %v != %v", f.Boot, other.Boot)
	}
	if f.Debugger != other.Debugger {
		return fmt.Errorf("Debugger: %v != %v", f.Debugger, other.Debugger)
	}
	if f.KeyUsage != other.KeyUsage {
		return fmt.Errorf("KeyUsage: %v != %v", f.KeyUsage, other.KeyUsage)
	}
	if f.Wildcard != other.Wildcard {
		return fmt.Errorf("Wildcard: %v != %v", f.Wildcard, other.Wildcard)
	}
	return nil
}
