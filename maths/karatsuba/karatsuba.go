// Karatsuba Big Int multiplication
package karatsuba

import (
	"fmt"
	"math/big"
)

func k_multiply(a, b *big.Int) *big.Int {
	fmt.Println(a)
	fmt.Println(a.BitLen())
	m := _pivot(a, b)
	splitA := _split(a, uint(m))
	leftA, rightA := splitA[0], splitA[1]

	splitB := _split(b, uint(m))
	leftB, rightB := splitB[0], splitB[1]

	fmt.Println(leftB, rightB, leftA, rightA)
	return a
}

func _split(a *big.Int, m uint) []*big.Int {
	big_a := big.NewInt(a.Int64())
	left := big.NewInt(int64(a.Uint64())).Rsh(big_a, m)

	tempRight := big.NewInt(left.Int64())
	tempRight = tempRight.Rsh(tempRight, m)

	_right := big.NewInt(a.Int64())
	right := _right.Sub(_right, tempRight)

	return []*big.Int{left, right}
}

func _pivot(a, b *big.Int) int {
	bitLen_a := a.BitLen()
	bitLen_b := b.BitLen()

	if bitLen_a < bitLen_b {
		return bitLen_a >> 1
	} else {
		return bitLen_b >> 1
	}
}

func multiply(a, b *big.Int) *big.Int {
	a1 := big.NewInt(a.Int64())
	b1 := big.NewInt(b.Int64())

	c := big.NewInt(0)
	return c.Mul(a1, b1)
}
