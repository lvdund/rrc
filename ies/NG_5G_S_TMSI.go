package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type NG_5G_S_TMSI struct {
	Value aper.BitString `lb:48,ub:48,madatory`
}

func (ie *NG_5G_S_TMSI) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteBitString(ie.Value.Bytes, uint(ie.Value.NumBits), &aper.Constraint{Lb: 48, Ub: 48}, false); err != nil {
		return utils.WrapError("Encode NG_5G_S_TMSI", err)
	}
	return nil
}

func (ie *NG_5G_S_TMSI) Decode(r *aper.AperReader) error {
	var err error
	var v []byte
	var n uint
	if v, n, err = r.ReadBitString(&aper.Constraint{Lb: 48, Ub: 48}, false); err != nil {
		return utils.WrapError("Decode NG_5G_S_TMSI", err)
	}
	ie.Value = aper.BitString{
		Bytes:   v,
		NumBits: uint64(n),
	}
	return nil
}
