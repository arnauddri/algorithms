// Karatsuba Big Int multiplication
package karatsuba

import (
	"math"
	"math/big"
)

func k_multiply(a, b *big.Int) *big.Int {
	if a.Cmp(big.NewInt(10)) < 1 || b.Cmp(big.NewInt(10)) < 1 {
		return mul(a, b)
	}

	m := _pivot(a, b)

	leftA, rightA := _split(a, uint(m))
	leftB, rightB := _split(b, uint(m))

	z0 := k_multiply(leftA, leftB)
	z1 := k_multiply(rightA, rightB)
	z2 := k_multiply(add(leftA, rightA), add(leftB, rightB))
	z2 = sub(z2, add(z0, z1))

	temp0 := mul(z0, big.NewInt(int64(math.Pow(10.0, 2.0*float64(m)))))
	temp1 := mul(z2, big.NewInt(int64(math.Pow(10.0, float64(m)))))
	temp2 := add(temp0, temp1)

	return add(temp2, z1)
}

func _split(a *big.Int, m uint) (left, right *big.Int) {
	denominator := big.NewInt(int64(math.Pow(10.0, float64(m))))

	left = big.NewInt(0).Div(a, denominator)
	right = sub(a, big.NewInt(0).Mul(left, denominator))

	return
}

func _pivot(a, b *big.Int) int {
	len_a := len(a.String())
	len_b := len(b.String())

	if len_a > len_b {
		return len_a / 2
	} else {
		return len_b / 2
	}
}

func add(a, b *big.Int) *big.Int {
	return big.NewInt(0).Add(a, b)
}

func mul(a, b *big.Int) *big.Int {
	return big.NewInt(0).Mul(a, b)
}

func sub(a, b *big.Int) *big.Int {
	return big.NewInt(0).Sub(a, b)
}
