package mupfmt

type ProtectionFlags struct {
	Write    bool
	Boot     bool
	Debugger bool
	Wildcard bool
}

type Input struct {
	// SHE module Unique ID as hex string
	UID string

	// ID of the secret key
	AuthKeyID uint32

	// secret key as hex string
	AuthKey string

	// ID of the slot to be updated
	ID uint32

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

func (i Input) Encode() (*EncodeResult, error) {
	var result EncodeResult
	return &result, nil
}
