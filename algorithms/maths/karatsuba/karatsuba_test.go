// Karatsuba Big Int multiplication
package karatsuba

import (
	"fmt"
	"math/big"
	"testing"
)

func TestMultipyOdd(t *testing.T) {
	a := big.NewInt(1234567890123456789)

	if k_multiply(a, a).Cmp(mul(a, a)) != 0 {
		fmt.Println(k_multiply(a, a), mul(a, a))
		t.Fatal("Wrong result")
	}
}

func TestMultipyEven(t *testing.T) {
	a := big.NewInt(123456789012345678)

	if k_multiply(a, a).Cmp(mul(a, a)) != 0 {
		fmt.Println(k_multiply(a, a), mul(a, a))
		t.Fatal("Wrong result")
	}
}

func BenchmarkMultiply(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := big.NewInt(1234567890123456789)
		mul(a, a)
	}
}

func BenchmarkKaratsubaMultiply(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := big.NewInt(1234567890123456789)
		k_multiply(a, a)
	}
}
