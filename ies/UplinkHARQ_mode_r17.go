package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type UplinkHARQ_mode_r17 struct {
	Value aper.BitString `lb:32,ub:32,madatory`
}

func (ie *UplinkHARQ_mode_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteBitString(ie.Value.Bytes, uint(ie.Value.NumBits), &aper.Constraint{Lb: 32, Ub: 32}, false); err != nil {
		return utils.WrapError("Encode UplinkHARQ_mode_r17", err)
	}
	return nil
}

func (ie *UplinkHARQ_mode_r17) Decode(r *aper.AperReader) error {
	var err error
	var v []byte
	var n uint
	if v, n, err = r.ReadBitString(&aper.Constraint{Lb: 32, Ub: 32}, false); err != nil {
		return utils.WrapError("Decode UplinkHARQ_mode_r17", err)
	}
	ie.Value = aper.BitString{
		Bytes:   v,
		NumBits: uint64(n),
	}
	return nil
}
