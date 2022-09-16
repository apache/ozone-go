package utils

import (
	"encoding/binary"
)

// Uint64toBytes TODO
func Uint64toBytes(b []byte, v uint64) {
	binary.LittleEndian.PutUint64(b, v)
}

// Uint32toBytes TODO
func Uint32toBytes(b []byte, v uint32) {
	for i := uint(0); i < 4; i++ {
		b[3-i] = byte(v >> (i * 8))
	}
}

// Uint16toBytes TODO
func Uint16toBytes(b []byte, v uint16) {
	b[0] = byte(v >> 8)
	b[1] = byte(v)
}

// Uint8toBytes TODO
func Uint8toBytes(b []byte, v uint8) {
	b[0] = byte(v)
}

// UintToBitSet TODO
func UintToBitSet(n uint) []byte {
	bitSet := make([]byte, 0)
	for n > 0 {
		if n&1 != 0 {
			bitSet = append(bitSet, 1)
		} else {
			bitSet = append(bitSet, 0)
		}
		n = n >> 1
	}
	return bitSet
}
