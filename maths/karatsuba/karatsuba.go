// Karatsuba Big Int multiplication
package karatsuba

import (
	"fmt"
	"math/big"
)

func multiply(a, b *big.Int) *big.Int {
	fmt.Println(a)
	return a
}

func _split(a *big.Int, m uint) []*big.Int {
	big_a := big.NewInt(a.Int64())

	left := big.NewInt(int64(a.Uint64())).Rsh(big_a, m)

	return []*big.Int{left}
}

func _pivot(a, b *big.Int) int {
	bitLen_a := a.BitLen()
	bitLen_b := b.BitLen()

	if bitLen_a < bitLen_b {
		return bitLen_a / 2
	} else {
		return bitLen_b / 2
	}
}
