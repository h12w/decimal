package decimal

import (
	"bytes"
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

func (d D) MarshalJSON() ([]byte, error) {
	buf, err := d.dec.MarshalJSON()
	if err != nil {
		return nil, err
	}
	return bytes.Trim(buf, `"`), nil
}

func (d *D) UnmarshalJSON(data []byte) error {
	buf := make([]byte, len(data)+2)
	copy(buf[1:], data)
	buf[0], buf[len(buf)-1] = '"', '"'
	return d.dec.UnmarshalJSON(buf)
}
