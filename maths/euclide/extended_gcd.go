package division

import ()

func divide(a, b int) int {
	if a < b {
		return a
	} else {
		return divide(a-b, b)
	}
}

func getCoeff(a, b int) (int, int) {
	_a := a
	_b := b

	if a < 0 {
		a *= -1
	}
	if b < 0 {
		b *= -1
	}

	x0, y0 := 0, 1
	x1, y1 := 1, 0

	for b != 0 {
		quotient := a / b
		a, b = b, divide(a, b)
		x1, x0 = x0-quotient*x1, x1
		y1, y0 = y0-quotient*y1, y1
	}

	if _a < 0 {
		y0 *= -1
	}
	if _b < 0 {
		x0 *= -1
	}

	return y0, x0
}
