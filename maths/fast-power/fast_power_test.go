package fast_power

import (
	"math"
	"testing"
)

func TestFastPower(t *testing.T) {
	a := uint32(2)

	for i := 0; i < 100000; i++ {
		n, ok := fast_power(a, i)

		if ok != nil || n != uint32(math.Pow(float64(a), float64(i))) {
			t.Error()
		}
	}
}
