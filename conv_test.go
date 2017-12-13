package decimal

import "testing"

func TestConv(t *testing.T) {
	if Float(3.14).Float64() != 3.14 {
		t.Fatal()
	}
}

func TestInt(t *testing.T) {
	if Float(3.14).Int() != 3 {
		t.Fatal()
	}
	if Float(3.99).Int() != 3 {
		t.Fatal()
	}
}

func TestElemAddr(t *testing.T) {
	d := Int(3)
	if d.Elem().Int() != 3 {
		t.Fatal()
	}
}

func TestAddr(t *testing.T) {
	d := Int(3)
	if d.Addr().Int() != (&d).Int() {
		t.Fatal()
	}
}
