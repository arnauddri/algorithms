package fast_power

import (
	"errors"
)

// Exponentiation by squaring - O(log n). Overflow wraps modulo 2^32.
func fast_power(n uint32, power int) (uint32, error) {
	if power < 0 {
		return 0, errors.New("Power must be a positive integer or zero")
	}

	result := uint32(1)
	for factor := n; power > 0; power, factor = power>>1, factor*factor {
		if power&1 == 1 {
			result *= factor
		}
	}

	return result, nil
}

// Iterative - O(n)
func slow_power(n uint32, power int) (uint32, error) {
	if power < 0 {
		return 0, errors.New("Power must be a positive integer or zero")
	}

	a := uint32(1)
	for i := 0; i < power; i++ {
		a *= n
	}
	return a, nil
}
