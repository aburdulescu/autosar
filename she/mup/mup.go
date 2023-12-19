package mup

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/binary"
	"encoding/hex"
	"fmt"

	"github.com/aburdulescu/autosar/internal/aesmp"
	"github.com/aburdulescu/autosar/she"

	"github.com/aead/cmac"
)

type Input struct {
	// SHE module Unique ID as hex string
	UID string

	// ID of the secret key
	AuthID she.KeyID

	// secret key as hex string
	AuthKey string

	// ID of the slot to be updated
	ID she.KeyID

	// the new value of the key as hex string
	NewKey string

	// new value of the counter
	Counter uint32

	// key flags
	Flags ProtectionFlags
}

// memory update protocol encode result
type EncodeResult struct {
	M1 [16]byte
	M2 [32]byte
	M3 [16]byte
}

// Encode the memory update protocol input
func (in Input) Encode() (*EncodeResult, error) {
	if err := in.ID.IsCompatible(in.AuthID); err != nil {
		return nil, err
	}

	authKey, err := hex.DecodeString(in.AuthKey)
	if err != nil {
		return nil, err
	}
	if len(authKey) != 16 {
		return nil, fmt.Errorf("AuthKey expected length is 16 bytes, have %d bytes", len(authKey))
	}

	encConst := sheKeyUpdateEncConstBase.Encode()
	macConst := sheKeyUpdateMacConstBase.Encode()

	// if withLogs {
	// 	log.Println("ENC_C:", hex.EncodeToString(encConst))
	// 	log.Println("MAC_C:", hex.EncodeToString(macConst))
	// }

	authkey, err := hex.DecodeString(in.AuthKey)
	if err != nil {
		return nil, fmt.Errorf("authkey: %w", err)
	}

	k1, err := aesmp.Compress(authkey, encConst)
	if err != nil {
		return nil, err
	}

	k2, err := aesmp.Compress(authkey, macConst)
	if err != nil {
		return nil, err
	}

	// if withLogs {
	// 	log.Println("K1:", hex.EncodeToString(k1))
	// 	log.Println("K2:", hex.EncodeToString(k2))
	// }

	m1, err := in.encodeM1()
	if err != nil {
		return nil, err
	}

	m2, err := in.encodeM2(k1)
	if err != nil {
		return nil, err
	}

	m3, err := in.encodeM3(k2, m1, m2)
	if err != nil {
		return nil, err
	}

	var result EncodeResult
	copy(result.M1[:], m1)
	copy(result.M2[:], m2)
	copy(result.M3[:], m3)

	return &result, nil
}

func (in Input) encodeM1() ([]byte, error) {
	uid, err := hex.DecodeString(in.UID)
	if err != nil {
		return nil, err
	}

	if len(uid) != 15 {
		return nil, fmt.Errorf("UID expected length is 15 bytes, have %d bytes", len(uid))
	}

	// if withLogs {
	// 	log.Println("KID:", in.KeyId)
	// 	log.Println("AuthID:", in.AuthKeyId)
	// }

	var r []byte
	r = append(r, uid...)
	r = append(r, uint8(in.ID)<<4|uint8(in.AuthID))

	return r, nil
}

func (in Input) encodeM2(k1 []byte) ([]byte, error) {
	newKey, err := hex.DecodeString(in.NewKey)
	if err != nil {
		return nil, err
	}
	if len(newKey) != 16 {
		return nil, fmt.Errorf("NewKey expected length is 16 bytes, have %d bytes", len(newKey))
	}

	flags := in.Flags.encode()
	counterAndFlags := encodeCounterAndFlags(in.Counter, flags)
	// if withLogs {
	// 	log.Printf("Cid: %032b %v\n", in.Counter, in.Counter)
	// 	log.Printf("Fid: %08b %v\n", flags, in.Flags)
	// 	log.Printf("Cid|Fid: %064b\n", counterAndFlags)
	// }

	data := make([]byte, 32)

	binary.BigEndian.PutUint64(data[0:8], counterAndFlags)

	copy(data[16:], newKey)

	iv := make([]byte, 16) // all zeros

	return cbcEncrypt(k1, iv, data)
}

func (in Input) encodeM3(k2, m1, m2 []byte) ([]byte, error) {
	var m1m2 []byte
	m1m2 = append(m1m2, m1...)
	m1m2 = append(m1m2, m2...)
	return cmacGenerate(k2, m1m2)
}

type KeyUpdateConst [4]uint32

func (c KeyUpdateConst) Encode() []byte {
	r := make([]byte, 16)
	binary.BigEndian.PutUint32(r[0:4], c[0])
	binary.BigEndian.PutUint32(r[4:8], c[1])
	binary.BigEndian.PutUint32(r[8:12], c[2])
	binary.BigEndian.PutUint32(r[12:16], c[3])
	return r
}

var (
	sheKeyUpdateEncConstBase = KeyUpdateConst{0x01015348, 0x45008000, 0x00000000, 0x000000b0}
	sheKeyUpdateMacConstBase = KeyUpdateConst{0x01025348, 0x45008000, 0x00000000, 0x000000b0}
)

func cbcEncrypt(key, iv, plaintext []byte) ([]byte, error) {
	return cbc(key, iv, plaintext, true)
}

func cbcDecrypt(key, iv, plaintext []byte) ([]byte, error) {
	return cbc(key, iv, plaintext, false)
}

func cbc(key, iv, in []byte, direction bool) ([]byte, error) {
	if len(in)%aes.BlockSize != 0 {
		return nil, fmt.Errorf("data is not a multiple of the block size")
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	out := make([]byte, len(in))
	var mode cipher.BlockMode
	if direction {
		mode = cipher.NewCBCEncrypter(block, iv)
	} else {
		mode = cipher.NewCBCDecrypter(block, iv)
	}
	mode.CryptBlocks(out, in)
	return out, nil
}

func cmacGenerate(key, msg []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	result, err := cmac.Sum(msg, c, c.BlockSize())
	if err != nil {
		return nil, err
	}
	return result, nil
}

func encodeCounterAndFlags(counter uint32, flags uint8) uint64 {
	var v uint64

	// add counter(28bits), i.e. bits 63-36
	v = uint64(counter << 4)

	// add first 4 flags, i.e. bits 35-32
	v |= uint64((flags >> 1) & 0x0f)

	// shift first 32 bits
	v <<= 32

	// add last flag, i.e. bit 31
	v |= uint64(flags&0x01) << 31

	return v
}

type ProtectionFlags struct {
	// Enable write-protection
	Write bool
	// Enable boot-protection
	Boot bool
	// Enable debugger-protection
	Debugger bool
	// key usage: false=encryption/decryption, true=MAC generation/verification
	KeyUsage bool
	// Enable wildcard-protection
	Wildcard bool
}

func (f ProtectionFlags) String() string {
	return fmt.Sprintf(
		"WRITE_PROTECTION=%v, BOOT_PROTECTION=%v, DEBUGGER_PROTECTION=%v, KEY_USAGE=%v, WILDCARD=%v}",
		f.Write,
		f.Boot,
		f.Debugger,
		f.KeyUsage,
		f.Wildcard,
	)
}

func (f *ProtectionFlags) decode(v uint8) {
	f.Write = getBitVal(v, 4) == 1
	f.Boot = getBitVal(v, 3) == 1
	f.Debugger = getBitVal(v, 2) == 1
	f.KeyUsage = getBitVal(v, 1) == 1
	f.Wildcard = getBitVal(v, 0) == 1
}

func (f ProtectionFlags) encode() uint8 {
	var v uint8
	if f.Write {
		v |= putBitVal(1, 4)
	}
	if f.Boot {
		v |= putBitVal(1, 3)
	}
	if f.Debugger {
		v |= putBitVal(1, 2)
	}
	if f.KeyUsage {
		v |= putBitVal(1, 1)
	}
	if f.Wildcard {
		v |= putBitVal(1, 0)
	}
	return v
}

// get bit value of bit number 'i' from 'v'
func getBitVal(v uint8, i int) uint8 {
	mask := uint8(1) << i
	byteVal := (v & mask) >> i
	bitVal := byteVal & 0x01
	return bitVal
}

// put bit value 'v' to bit number 'i'
func putBitVal(v uint8, i int) uint8 {
	bitVal := v & 0x01
	return bitVal << i
}
