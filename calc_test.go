package decimal

import "testing"

func TestArith(t *testing.T) {
	if !Float(1.2).Add(Float(3.4)).Equal(Float(4.6)) {
		t.Fatal()
	}
	if !Float(1.2).Sub(Float(3.4)).Equal(Float(-2.2)) {
		t.Fatal()
	}
	if !Float(1.2).Mul(Float(3.4)).Equal(Float(4.08)) {
		t.Fatal()
	}
	if !Float(1.2).Div(Float(2.4)).Equal(Float(0.5)) {
		t.Fatal()
	}
	if !Float(1.2).LessThan(Float(1.3)) {
		t.Fatal()
	}
}
