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

func iter(a, b int) int {
	if a == b {
		return a
	}
	if a == 0 || b == 0 {
		return a + b
	}

	shift := 0
	for shift = 0; ((a | b) & 1) == 0; shift++ {
		a >>= 1
		b >>= 1
	}

	for (a & 1) == 0 {
		a >>= 1
	}

	for b != 0 {
		for b&1 == 0 {
			b >>= 1
		}
		if a > b {
			b, a = a, b
		}
		b -= a
	}

	return a << uint32(shift)
}
