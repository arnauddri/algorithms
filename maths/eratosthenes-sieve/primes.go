package primes

import (
	"math"
)

func getAllPrimesTo(n int) (primes []int) {
	checked := make([]bool, n)
	sqrt_n := int(math.Sqrt(float64(n)))

	for i := 2; i <= sqrt_n; i++ {
		if !checked[i] {
			for j := i * i; j < n; j += i {
				checked[j] = true
			}
		}
	}

	for i := 1; i < n; i++ {
		if !checked[i] {
			primes = append(primes, i)
		}
	}

	return
}
