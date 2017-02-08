package decimal

import (
	"github.com/shopspring/decimal"
)

type D struct {
	dec decimal.Decimal
}

func String(v string) (D, error) {
	dec, err := decimal.NewFromString(v)
	if err != nil {
		return D{}, err
	}
	return D{dec}, err
}

func Float(v float64) D {
	return D{decimal.NewFromFloat(v)}
}

func Int(v int) D {
	return D{decimal.New(int64(v), 0)}
}
