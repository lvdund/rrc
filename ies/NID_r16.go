package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type NID_r16 struct {
	Value aper.BitString `lb:44,ub:44,madatory`
}

func (ie *NID_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteBitString(ie.Value.Bytes, uint(ie.Value.NumBits), &aper.Constraint{Lb: 44, Ub: 44}, false); err != nil {
		return utils.WrapError("Encode NID_r16", err)
	}
	return nil
}

func (ie *NID_r16) Decode(r *aper.AperReader) error {
	var err error
	var v []byte
	var n uint
	if v, n, err = r.ReadBitString(&aper.Constraint{Lb: 44, Ub: 44}, false); err != nil {
		return utils.WrapError("Decode NID_r16", err)
	}
	ie.Value = aper.BitString{
		Bytes:   v,
		NumBits: uint64(n),
	}
	return nil
}
