package newton_sqrt

import (
	"testing"
)

func TestNewtonSqrt(t *testing.T) {
	if newton_sqrt(333, 1e-7, 1e7) != 18.248287591617554 ||
		newton_sqrt(61009, 1e-7, 1e7) != 247 ||
		newton_sqrt(9, 1e-7, 1e7) != 3 {
		t.Error()
	}
}
