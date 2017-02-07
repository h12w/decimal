package decimal

import "testing"

func TestConv(t *testing.T) {
	if Float(3.14).Float64() != 3.14 {
		t.Fatal()
	}
}
