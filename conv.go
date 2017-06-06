package decimal

func (d D) Float64() float64 {
	f, _ := d.dec.Float64()
	return f
}

func (d D) Int() int {
	return int(d.dec.IntPart())
}
