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
	//fmt.Println(multiply(a, b))
	fmt.Println(k_multiply(a, b))
}
