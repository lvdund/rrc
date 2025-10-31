package utils

import (
	"rrc/aper"
)

const (
	ChoiceNothing = 0
	ChoiceRelease = 1
	ChoiceSetup   = 2
)

type BOOLEAN struct {
	// Value aper.Boolean
	// c     aper.Constraint
	// ext   bool
}

func NewBOOLEAN(v bool, c aper.Constraint, ext bool) BOOLEAN {
	return BOOLEAN{
		// Value: aper.Boolean(v),
		// c:     c,
		// ext:   ext,
	}
}

func (t *BOOLEAN) Encode(w *aper.AperWriter) (err error) {
	// err = w.WriteBoolean(bool(t.Value), t.c, t.ext)
	return
}
func (t *BOOLEAN) Decode(r *aper.AperReader) (err error) {
	// v, err := r.ReadBoolean(t.c, t.ext)
	// t.Value = aper.Boolean(v)
	return
}

type ENUMERATED struct {
	Value aper.Enumerated
	c     aper.Constraint
	ext   bool
}

func NewENUMERATED(v int64, c aper.Constraint, ext bool) ENUMERATED {
	return ENUMERATED{
		Value: aper.Enumerated(v),
		c:     c,
		ext:   ext,
	}
}
func (t *ENUMERATED) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(t.Value), t.c, t.ext)
	return
}
func (t *ENUMERATED) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(t.c, t.ext)
	t.Value = aper.Enumerated(v)
	return
}

type BITSTRING struct {
	Value aper.BitString
	c     aper.Constraint
	ext   bool
}

func NewBITSTRING(v aper.BitString, c aper.Constraint, ext bool) BITSTRING {
	return BITSTRING{
		Value: aper.BitString{
			Bytes:   v.Bytes,
			NumBits: v.NumBits,
		},
		c:   c,
		ext: ext,
	}
}
func (t *BITSTRING) Encode(w *aper.AperWriter) (err error) {
	if t.c.Lb == t.c.Ub {
		t.Value.NumBits = uint64(t.c.Lb)
	} else if len(t.Value.Bytes)*8 < int(t.c.Lb) {
		t.Value.NumBits = uint64(t.c.Lb)
	}
	err = w.WriteBitString(t.Value.Bytes, uint(t.Value.NumBits), &t.c, t.ext)
	return
}
func (t *BITSTRING) Decode(r *aper.AperReader) (err error) {
	var v []byte
	var n uint
	if v, n, err = r.ReadBitString(&t.c, t.ext); err != nil {
		return
	}
	t.Value.Bytes = v
	t.Value.NumBits = uint64(n)
	return
}

type OCTETSTRING struct {
	Value aper.OctetString
	c     aper.Constraint
	ext   bool
}

func NewOCTETSTRING(v []byte, c aper.Constraint, ext bool) OCTETSTRING {
	return OCTETSTRING{
		Value: v,
		c:     c,
		ext:   ext,
	}
}
func (t *OCTETSTRING) Encode(w *aper.AperWriter) (err error) {
	if t.c.Lb == t.c.Ub && t.c.Lb == 0 {
		err = w.WriteOctetString(t.Value, nil, t.ext)
	} else {
		err = w.WriteOctetString(t.Value, &t.c, t.ext)
	}
	return
}
func (t *OCTETSTRING) Decode(r *aper.AperReader) (err error) {
	var v aper.OctetString
	if t.c.Lb == t.c.Ub && t.c.Lb == 0 {
		if v, err = r.ReadOctetString(nil, t.ext); err != nil {
			return
		}
	} else {
		if v, err = r.ReadOctetString(&t.c, t.ext); err != nil {
			return
		}
	}

	t.Value = v
	return
}

type INTEGER struct {
	Value aper.Integer
	c     aper.Constraint
	ext   bool
}

func NewINTEGER(v int64, c aper.Constraint, ext bool) INTEGER {
	return INTEGER{
		Value: aper.Integer(v),
		c:     c,
		ext:   ext,
	}
}
func (t *INTEGER) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteInteger(int64(t.Value), &t.c, t.ext)
	return
}
func (t *INTEGER) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadInteger(&t.c, t.ext)
	t.Value = aper.Integer(v)
	return
}

type Sequence[T aper.IE] struct {
	Value []T
	c     aper.Constraint
	ext   bool
}

func NewSequence[T aper.IE](items []T, c aper.Constraint, ext bool) Sequence[T] {
	return Sequence[T]{
		Value: items,
		c:     c,
		ext:   ext,
	}
}

func (s *Sequence[T]) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[T](s.Value, w, &s.c, s.ext); err != nil {
		return
	}
	return
}
func (s *Sequence[T]) Decode(r *aper.AperReader, fn func() T) (err error) {
	var newItems []T
	newItems, err = aper.ReadSequenceOfEx(fn, r, &s.c, s.ext)
	if err != nil {
		return
	}
	s.Value = []T{}
	s.Value = append(s.Value, newItems...)
	return
}

// temparory
type TAC struct {
	Value []byte
}

func (ie *TAC) Encode(w *aper.AperWriter) (err error) {
	return
}
func (ie *TAC) Decode(r *aper.AperReader) (err error) {
	return
}

type PLMNIdentity struct {
	Value []byte
}

func (ie *PLMNIdentity) Encode(w *aper.AperWriter) (err error) {
	return
}
func (ie *PLMNIdentity) Decode(r *aper.AperReader) (err error) {
	return
}

type EmergencyAreaID struct {
	Value []byte
}

func (ie *EmergencyAreaID) Encode(w *aper.AperWriter) (err error) {
	return
}
func (ie *EmergencyAreaID) Decode(r *aper.AperReader) (err error) {
	return
}

type TransportLayerAddress struct {
	Value []byte
}

func (ie *TransportLayerAddress) Encode(w *aper.AperWriter) (err error) {
	return
}
func (ie *TransportLayerAddress) Decode(r *aper.AperReader) (err error) {
	return
}