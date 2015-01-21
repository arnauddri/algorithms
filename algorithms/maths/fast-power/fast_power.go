package fast_power

import (
	"errors"
	"math"
)

// Recursive - O(log n)
func fast_power(n uint32, power int) (uint32, error) {
	if power < 0 && math.Floor(float64(power)) == float64(power) {
		return uint32(math.NaN()), errors.New("Power must be a positive integer or zero")
	}

	if power == 0 {
		return 1, nil
	}

	var factor uint32
	var result uint32

	mul := func(v uint32) {
		if result == 0 {
			result = v
		} else {
			result *= v
		}
	}

	for factor = n; power > 0; power, factor = power>>1, factor*factor {
		if power&1 == 1 {
			mul(factor)
		}
	}

	return result, nil
}

// Iterative - O(n)
func slow_power(n uint32, power int) (uint32, error) {
	if power < 0 && math.Floor(float64(power)) == float64(power) {
		return uint32(math.NaN()), errors.New("Power must be a positive integer or zero")
	}

	a := uint32(1)
	for i := 0; i < power; i++ {
		a *= n
	}
	return a, nil
}
