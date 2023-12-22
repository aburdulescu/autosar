package cmac

import (
	"crypto/aes"
	"errors"
)

const bs = 16

func Sum(key, msg []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(msg)%bs != 0 {
		return nil, errors.New("unaligned input not supported for now")
	}

	aligned := len(msg)%bs == 0

	n := len(msg) / bs

	if n == 0 {
		n = 1
		aligned = false
	}

	var k1 [bs]byte
	var k2 [bs]byte

	var mLast [bs]byte
	if aligned {
		copy(mLast[:], msg[(len(msg)-1)*16:])
		xor(mLast[:], k1[:])
	} else {
		copy(mLast[:], msg[(len(msg)-1)*16:])
		xor(mLast[:], k2[:])
	}

	var x [bs]byte

	for i := 0; i < n-1; i++ {
		mi := msg[i*16 : (i+1)*bs]
		xor(x[:], mi)
		c.Encrypt(x[:], x[:])
	}

	xor(mLast[:], x[:])
	c.Encrypt(mLast[:], mLast[:])

	return mLast[:], nil
}

func xor(dst, src []byte) {
	for i := range dst {
		dst[i] = dst[i] ^ src[i]
	}
}
