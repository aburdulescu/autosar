package mupfmt

import (
	"encoding/hex"
	"fmt"
)

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

type Input struct {
	// SHE module Unique ID as hex string
	UID string

	// ID of the secret key
	AuthID KeyID

	// secret key as hex string
	AuthKey string

	// ID of the slot to be updated
	ID KeyID

	// the new value of the key as hex string
	NewKey string

	// new value of the counter
	Counter uint32

	// key flags
	Flags ProtectionFlags
}

type EncodeResult struct {
	M1 [16]byte
	M2 [32]byte
	M3 [16]byte
}

func (in Input) Encode() (*EncodeResult, error) {
	uid, err := hex.DecodeString(in.UID)
	if err != nil {
		return nil, err
	}
	if len(uid) != 15 {
		return nil, fmt.Errorf("UID expected length is 15 bytes, have %d bytes", len(uid))
	}

	authKey, err := hex.DecodeString(in.AuthKey)
	if err != nil {
		return nil, err
	}
	if len(authKey) != 16 {
		return nil, fmt.Errorf("AuthKey expected length is 16 bytes, have %d bytes", len(authKey))
	}

	newKey, err := hex.DecodeString(in.NewKey)
	if err != nil {
		return nil, err
	}
	if len(newKey) != 16 {
		return nil, fmt.Errorf("NewKey expected length is 16 bytes, have %d bytes", len(newKey))
	}

	if err := in.ID.IsCompatible(in.AuthID); err != nil {
		return nil, err
	}

	var result EncodeResult
	return &result, nil
}

//go:generate stringer -type=KeyID
type KeyID uint8

const (
	SECRET_KEY KeyID = iota
	MASTER_ECU_KEY
	BOOT_MAC_KEY
	BOOT_MAC
	KEY_1
	KEY_2
	KEY_3
	KEY_4
	KEY_5
	KEY_6
	KEY_7
	KEY_8
	KEY_9
	KEY_10
	RAM_KEY
)

func (id KeyID) IsValid() bool {
	return id < KeyID(len(_KeyID_index)-1)
}

func (id KeyID) IsCompatible(other KeyID) error {
	if !id.IsValid() {
		return fmt.Errorf("%s is not valid", id)
	}
	if !other.IsValid() {
		return fmt.Errorf("%s is not valid", other)
	}

	switch id {
	case
		MASTER_ECU_KEY,
		BOOT_MAC_KEY,
		BOOT_MAC:

		if other != id {
			return fmt.Errorf("%s can be updated only with itself", id)
		}

	case
		KEY_1,
		KEY_2,
		KEY_3,
		KEY_4,
		KEY_5,
		KEY_6,
		KEY_7,
		KEY_8,
		KEY_9,
		KEY_10:

		if other != id && other != MASTER_ECU_KEY {
			return fmt.Errorf("%s can be updated only with itself or %s", id, MASTER_ECU_KEY)
		}

	case RAM_KEY:

		if other != SECRET_KEY &&
			other != KEY_1 &&
			other != KEY_2 &&
			other != KEY_3 &&
			other != KEY_4 &&
			other != KEY_5 &&
			other != KEY_6 &&
			other != KEY_7 &&
			other != KEY_8 &&
			other != KEY_9 &&
			other != KEY_10 {
			return fmt.Errorf("%s can be updated only with %s or one of KEY_n", id, SECRET_KEY)
		}

	default:
		return fmt.Errorf("%s is not valid or cannot be used", id)

	}

	return nil
}
