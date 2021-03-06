package decimal

import (
	"encoding/json"
	"encoding/xml"
	"strconv"
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
			t.Fatalf("expect\n%s\ngot\n%x\n", n, d128.D.String())
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
			t.Fatalf("expect\n%s\ngot\n%x\n", n, s.D.String())
		}
	}
	// from float64 to Decimal
	{
		f64, err := strconv.ParseFloat(n, 64)
		if err != nil {
			t.Fatal(err)
		}
		dBuf, err := bson.Marshal(&struct {
			F float64 `bson:"d"`
		}{f64})
		if err != nil {
			t.Fatal(err)
		}
		var s S
		if err := bson.Unmarshal(dBuf, &s); err != nil {
			t.Fatal(err)
		}
		if s.D.String() != n {
			t.Fatalf("expect\n%s\ngot\n%x\n", n, s.D.String())
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

func TestXMLAttr(t *testing.T) {
	n := "1.23456789"
	d, err := String(n)
	if err != nil {
		t.Fatal(err)
	}
	type testS struct {
		XMLName xml.Name `xml:"test"`
		D       D        `xml:"d,attr"`
	}
	s := testS{D: d}
	buf, err := xml.Marshal(s)
	if err != nil {
		t.Fatal(err)
	}
	if string(buf) != `<test d="1.23456789"></test>` {
		t.Fatal("got " + string(buf))
	}
	{
		var s testS
		if err := xml.Unmarshal(buf, &s); err != nil {
			t.Fatal(err)
		}
		if s.D.String() != n {
			t.Fatalf("expect %s got %s", n, s.D.String())
		}
	}
}

func TestXMLCharData(t *testing.T) {
	n := "1.23456789"
	d, err := String(n)
	if err != nil {
		t.Fatal(err)
	}
	type testS struct {
		XMLName xml.Name `xml:"test"`
		D       D        `xml:",chardata"`
	}
	s := testS{D: d}
	buf, err := xml.Marshal(s)
	if err != nil {
		t.Fatal(err)
	}
	if string(buf) != `<test>1.23456789</test>` {
		t.Fatal("got " + string(buf))
	}
	{
		var s testS
		if err := xml.Unmarshal(buf, &s); err != nil {
			t.Fatal(err)
		}
		if s.D.String() != n {
			t.Fatalf("expect %s got %s", n, s.D.String())
		}
	}
}
