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
