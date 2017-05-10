package decimal

import (
	"encoding/json"
	"testing"

	"gopkg.in/mgo.v2/bson"
)

type S struct {
	D D `bson:"d" json:"d"`
}

func TestBSONCompatible(t *testing.T) {
	n := "3.14"
	// from Decimal to bson.Decimal128
	{
		d, err := String(n)
		if err != nil {
			t.Fatal(err)
		}
		dBuf, err := bson.Marshal(&S{d})
		if err != nil {
			t.Fatal(err)
		}
		var d128 struct {
			D bson.Decimal128 `bson:"d"`
		}
		if err := bson.Unmarshal(dBuf, &d128); err != nil {
			t.Fatal(err)
		}
		if d128.D.String() != n {
			t.Fatalf("expect\n%x\ngot\n%x\n", n, d128.D.String())
		}
	}
	// from bson.Decimal128 to Decimal
	{
		d128, err := bson.ParseDecimal128(n)
		if err != nil {
			t.Fatal(err)
		}
		dBuf, err := bson.Marshal(&struct {
			D bson.Decimal128 `bson:"d"`
		}{d128})
		if err != nil {
			t.Fatal(err)
		}
		var s S
		if err := bson.Unmarshal(dBuf, &s); err != nil {
			t.Fatal(err)
		}
		if s.D.String() != n {
			t.Fatalf("expect\n%x\ngot\n%x\n", n, s.D.String())
		}
	}

}

func TestJSON(t *testing.T) {
	n := "1.23456789"
	d, err := String(n)
	if err != nil {
		t.Fatal(err)
	}
	if d.String() != n {
		t.Fatalf("expect\n%s\ngot\n%s", n, d.String())
	}
	buf, err := json.Marshal(d)
	if err != nil {
		t.Fatal(err)
	}
	if string(buf) != n {
		t.Fatalf("expect\n%s\ngot\n%s", n, string(buf))
	}
	{
		var d D
		if err := json.Unmarshal(buf, &d); err != nil {
			t.Fatal(err)
		}
		if d.String() != n {
			t.Fatalf("expect\n%s\ngot\n%s", n, d.String())
		}
	}
}

func TestRound(t *testing.T) {
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
