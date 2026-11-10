package fast_power

import (
	"testing"
)

func TestFastPower(t *testing.T) {
	for i := 0; i < 32; i++ {
		n, err := fast_power(2, i)
		if err != nil || n != uint32(1)<<uint(i) {
			t.Errorf("2^%d: got %d", i, n)
		}
	}

	if n, err := fast_power(3, 5); err != nil || n != 243 {
		t.Errorf("3^5: got %d", n)
	}

	if _, err := fast_power(2, -1); err == nil {
		t.Error("negative power must return an error")
	}
}

func TestSlowPower(t *testing.T) {
	for i := 0; i < 32; i++ {
		n, err := slow_power(2, i)
		if err != nil || n != uint32(1)<<uint(i) {
			t.Errorf("2^%d: got %d", i, n)
		}
	}

	if _, err := slow_power(2, -1); err == nil {
		t.Error("negative power must return an error")
	}
}

// Both wrap modulo 2^32, so they must agree far beyond the overflow point.
func TestFastMatchesSlow(t *testing.T) {
	for _, base := range []uint32{2, 3, 7, 123456789} {
		for i := 0; i < 1000; i++ {
			fast, _ := fast_power(base, i)
			slow, _ := slow_power(base, i)
			if fast != slow {
				t.Errorf("%d^%d: fast %d, slow %d", base, i, fast, slow)
			}
		}
	}
}

func BenchmarkFastPower(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fast_power(2, 100000)
	}
}

func BenchmarkSlowPower(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slow_power(2, 100000)
	}
}
