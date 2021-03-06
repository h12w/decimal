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
	if Float(1.2).LessThan(Float(1.2)) {
		t.Fatal()
	}
	if Float(1.3).LessThan(Float(1.2)) {
		t.Fatal()
	}

	if !Float(1.2).LessEqual(Float(1.3)) {
		t.Fatal()
	}
	if !Float(1.2).LessEqual(Float(1.2)) {
		t.Fatal()
	}
	if Float(1.3).LessEqual(Float(1.2)) {
		t.Fatal()
	}

	if !Float(1.3).GreaterThan(Float(1.2)) {
		t.Fatal()
	}
	if Float(1.2).GreaterThan(Float(1.2)) {
		t.Fatal()
	}
	if !Float(1.3).GreaterThan(Float(1.2)) {
		t.Fatal()
	}

	if !Float(1.3).GreaterEqual(Float(1.2)) {
		t.Fatal()
	}
	if !Float(1.2).GreaterEqual(Float(1.2)) {
		t.Fatal()
	}
	if !Float(1.3).GreaterEqual(Float(1.2)) {
		t.Fatal()
	}
}

func TestRound(t *testing.T) {
	{
		n := "12.3456"
		d, err := String(n)
		if err != nil {
			t.Fatal(err)
		}
		r := d.Round(0)
		if r.String() != "12" {
			t.Fatalf("expect 12 got %s", r.String())
		}
		r = d.Round(1)
		if r.String() != "12.3" {
			t.Fatalf("expect 12.3 got %s", r.String())
		}
		r = d.Round(2)
		if r.String() != "12.35" {
			t.Fatalf("expect 12.35 got %s", r.String())
		}
		r = d.Round(3)
		if r.String() != "12.346" {
			t.Fatalf("expect 12.346 got %s", r.String())
		}
	}
	{
		n := "0.3456"
		d, err := String(n)
		if err != nil {
			t.Fatal(err)
		}
		r := d.Round(0)
		if r.String() != "0" {
			t.Fatalf("expect 0 got %s", r.String())
		}
		r = d.Round(1)
		if r.String() != "0.3" {
			t.Fatalf("expect 0.3 got %s", r.String())
		}
	}
	{
		n := "0.03456"
		d, err := String(n)
		if err != nil {
			t.Fatal(err)
		}
		r := d.Round(1)
		if r.String() != "0" {
			t.Fatalf("expect 0 got %s", r.String())
		}
		r = d.Round(2)
		if r.String() != "0.03" {
			t.Fatalf("expect 0.03 got %s", r.String())
		}
	}
}

func TestZero(t *testing.T) {
	var d D
	if !d.IsZero() {
		t.Fatal("expect zero")
	}
	if !d.LessThan(Float(0.1)) {
		t.Fatal("expect less")
	}
	if !Float(-0.1).LessThan(d) {
		t.Fatal("expect less")
	}
}
