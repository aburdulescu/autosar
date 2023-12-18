package mup

import (
	"encoding/hex"
	"fmt"

	"github.com/aburdulescu/autosar/she"
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
