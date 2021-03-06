package decimal

import "github.com/shopspring/decimal"

func (d D) Equal(o D) bool        { return d.dec.Equal(o.dec) }
func (d D) LessThan(o D) bool     { return d.dec.Cmp(o.dec) < 0 }
func (d D) LessEqual(o D) bool    { return d.dec.Cmp(o.dec) <= 0 }
func (d D) GreaterThan(o D) bool  { return d.dec.Cmp(o.dec) > 0 }
func (d D) GreaterEqual(o D) bool { return d.dec.Cmp(o.dec) >= 0 }

func (d D) Add(o D) D { return D{d.dec.Add(o.dec)} }
func (d D) Sub(o D) D { return D{d.dec.Sub(o.dec)} }
func (d D) Mul(o D) D { return D{d.dec.Mul(o.dec)} }
func (d D) Div(o D) D { return D{d.dec.Div(o.dec)} }

func (d D) Round(places int32) D { return D{d.dec.Round(places)} }

func (d D) IsZero() bool { return d.dec.Equal(decimal.Zero) }
