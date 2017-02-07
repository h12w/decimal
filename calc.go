package decimal

func (d D) Equal(o D) bool    { return d.dec.Equal(o.dec) }
func (d D) LessThan(o D) bool { return d.dec.Cmp(o.dec) < 0 }

func (d D) Add(o D) D { return D{d.dec.Add(o.dec)} }
func (d D) Sub(o D) D { return D{d.dec.Sub(o.dec)} }
func (d D) Mul(o D) D { return D{d.dec.Mul(o.dec)} }
func (d D) Div(o D) D { return D{d.dec.Div(o.dec)} }
