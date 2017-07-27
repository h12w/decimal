package decimal

func (d D) Float64() float64 {
	f, _ := d.dec.Float64()
	return f
}

func (d D) Int() int {
	return int(d.dec.IntPart())
}

func (d D) Addr() *D {
	return &d
}

func (d *D) Value() D {
	if d == nil {
		return Int(0)
	}
	return *d
}
