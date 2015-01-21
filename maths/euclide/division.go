package division

import ()

func divide(a, b int) int {
	if a < b {
		return a
	} else {
		return divide(a-b, b)
	}
}
