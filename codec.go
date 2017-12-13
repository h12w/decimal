package decimal

import (
	"bytes"
	"database/sql/driver"
	"encoding/xml"
	"fmt"
	"strings"

	"github.com/shopspring/decimal"
	"gopkg.in/mgo.v2/bson"
)

func (d *D) SetBSON(raw bson.Raw) error {
	switch raw.Kind {
	case 0x13:
		var d128 bson.Decimal128
		if err := raw.Unmarshal(&d128); err != nil {
			return err
		}
		dec, err := decimal.NewFromString(d128.String())
		if err != nil {
			return err
		}
		d.dec = dec
		return nil
	case 0x01:
		var f float64
		if err := raw.Unmarshal(&f); err != nil {
			return err
		}
		d.dec = decimal.NewFromFloat(f)
		return nil
	}
	return fmt.Errorf("expected data type %x", raw.Kind)
}

func (d D) GetBSON() (interface{}, error) {
	d128, err := bson.ParseDecimal128(d.String())
	if err != nil {
		return nil, err
	}
	// bson.Decimal128 must be wrapped in a struct to be valid
	return d128, nil
}

func (d D) String() string {
	return strings.Trim(d.dec.String(), `"`)
}

func (d D) MarshalText() ([]byte, error) {
	buf, err := d.dec.MarshalJSON()
	if err != nil {
		return nil, err
	}
	return bytes.Trim(buf, `"`), nil
}

func (d *D) UnmarshalText(data []byte) error {
	buf := make([]byte, len(data)+2)
	copy(buf[1:], data)
	buf[0], buf[len(buf)-1] = '"', '"'
	return d.dec.UnmarshalJSON(buf)
}

func (d D) MarshalJSON() ([]byte, error) {
	return d.MarshalText()
}

func (d *D) UnmarshalJSON(data []byte) error {
	return d.UnmarshalText(data)
}

func (d D) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	buf, err := d.dec.MarshalText()
	if err != nil {
		return xml.Attr{}, err
	}
	return xml.Attr{
		Name:  name,
		Value: string(buf),
	}, nil
}

func (d *D) UnmarshalXMLAttr(attr xml.Attr) error {
	return d.dec.UnmarshalText([]byte(attr.Value))
}

func (d D) MarshalBinary() (data []byte, err error) {
	return d.dec.MarshalBinary()
}

func (d *D) UnmarshalBinary(data []byte) error {
	return d.dec.UnmarshalBinary(data)
}

func (d D) Value() (driver.Value, error) {
	return d.dec.Value()
}

func (d *D) Scan(src interface{}) error {
	if src == nil {
		return nil
	}
	return d.dec.Scan(src)
}
