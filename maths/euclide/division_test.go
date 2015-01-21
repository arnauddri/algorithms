package division

import (
	"testing"
)

func TestRSelect(t *testing.T) {
	if divide(10, 2) != 0 && divide(10, 3) != 1 {
		t.Error()
	}
}
