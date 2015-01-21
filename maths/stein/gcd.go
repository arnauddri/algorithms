package gcd

import ()

func recurse(a, b int) int {
	if a == b {
		return a
	}
	if a == 0 || b == 0 {
		return a + b
	}

	if a%2 == 0 {
		if b%2 == 1 {
			return recurse(a>>1, b)
		} else {
			return recurse(a>>1, b>>1) << 1
		}
	}

	if b%2 == 0 {
		return recurse(a, b>>1)
	}

	if a > b {
		return recurse((a-b)>>1, b)
	}

	return recurse((b-a)>>1, a)
}
