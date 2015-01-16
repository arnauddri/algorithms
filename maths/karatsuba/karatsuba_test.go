// Karatsuba Big Int multiplication
package karatsuba

import (
	"fmt"
	"math/big"
	"testing"
)

func TestMultipy(t *testing.T) {
	a := big.NewInt(10)
	b := big.NewInt(10)
	fmt.Println(classic_multiply(a, b))
}

func classic_multiply(a, b *big.Int) *big.Int {
	a1 := big.NewInt(a.Int64())
	b1 := big.NewInt(b.Int64())

	c := big.NewInt(0)
	return c.Mul(a1, b1)
}
