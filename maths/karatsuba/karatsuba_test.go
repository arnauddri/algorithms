// Karatsuba Big Int multiplication
package karatsuba

import (
	"fmt"
	"math/big"
	"testing"
)

func TestMultipy(t *testing.T) {
	a := big.NewInt(1234)
	b := big.NewInt(1234)

	if k_multiply(a, b).Cmp(mul(a, b)) != 0 {
		fmt.Println(k_multiply(a, b), mul(a, b))
		t.Fatal("Wrong result")
	}
}
