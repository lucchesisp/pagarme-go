package client

import "testing"

func TestSum(t *testing.T) {
	if sum(1, 2) != 3 {
		t.Error("Sum did not return 3")
	}
}
