package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type ShortMAC_I struct {
	Value aper.BitString `lb:16,ub:16,madatory`
}

func (ie *ShortMAC_I) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteBitString(ie.Value.Bytes, uint(ie.Value.NumBits), &aper.Constraint{Lb: 16, Ub: 16}, false); err != nil {
		return utils.WrapError("Encode ShortMAC_I", err)
	}
	return nil
}

func (ie *ShortMAC_I) Decode(r *aper.AperReader) error {
	var err error
	var v []byte
	var n uint
	if v, n, err = r.ReadBitString(&aper.Constraint{Lb: 16, Ub: 16}, false); err != nil {
		return utils.WrapError("Decode ShortMAC_I", err)
	}
	ie.Value = aper.BitString{
		Bytes:   v,
		NumBits: uint64(n),
	}
	return nil
}
