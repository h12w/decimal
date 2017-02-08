package decimal

import "testing"

func TestNew(t *testing.T) {
	if d, err := String("1.23"); err != nil || d.String() != "1.23" {
		t.Fatal()
	}
	if Int(42).String() != "42" {
		t.Fatal()
	}
	if Float(3.14).String() != "3.14" {
		t.Fatal()
	}
}

func TestZero(t *testing.T) {
	if !Int(0).IsZero() {
		t.Fatal(Int(0).dec)
	}
	if !Float(0).IsZero() {
		t.Fatal(Int(0).dec)
	}
	if n, _ := String("0"); !n.IsZero() {
		t.Fatal(Int(0).dec)
	}
}
