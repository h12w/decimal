//go:generate msgp
package decimal

type (
	// W is wire type for D
	W struct {
		B []byte
	}
)

func (d D) Wire() W {
	bytes, _ := d.MarshalBinary()
	return W{B: bytes}
}

func (w W) Unmarshal() (d D, err error) {
	err = d.UnmarshalBinary(w.B)
	return d, err
}
