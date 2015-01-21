package division

import (
	"fmt"
	"testing"
)

func TestDivide(t *testing.T) {
	if divide(10, 2) != 0 && divide(10, 3) != 1 {
		t.Error()
	}
}

func TestGetCoeff(t *testing.T) {
	a, b := 11, 6
	bezout_a, bezout_b := getCoeff(11, 6)

	if bezout_a*a+bezout_b*b != 1 {
		fmt.Println(a, b, a*11+b*6)
		t.Error()
	}

	a, b = 12, 6
	bezout_a, bezout_b = getCoeff(12, 6)

	if bezout_a*a+bezout_b*b != 6 {
		fmt.Println(a, b, a*11+b*6)
		t.Error()
	}
}

func BenchmarkGetCoeff(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getCoeff(131313131, 121212121)
	}
}
