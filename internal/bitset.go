package internal

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"math/big"
)

func Bitset(v uint64, size int) protocol.Bitset { // this is from @dasciam, :>
	bitset := protocol.NewBitset(size)
	b := big.NewInt(int64(v))

	for s := range min(size, 64) {
		if b.Bit(s) != 0 {
			bitset.Set(s)
		}
	}
	return bitset
}
