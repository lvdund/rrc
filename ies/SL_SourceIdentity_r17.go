package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SL_SourceIdentity_r17 struct {
	Value aper.BitString `lb:24,ub:24,madatory`
}

func (ie *SL_SourceIdentity_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteBitString(ie.Value.Bytes, uint(ie.Value.NumBits), &aper.Constraint{Lb: 24, Ub: 24}, false); err != nil {
		return utils.WrapError("Encode SL_SourceIdentity_r17", err)
	}
	return nil
}

func (ie *SL_SourceIdentity_r17) Decode(r *aper.AperReader) error {
	var err error
	var v []byte
	var n uint
	if v, n, err = r.ReadBitString(&aper.Constraint{Lb: 24, Ub: 24}, false); err != nil {
		return utils.WrapError("Decode SL_SourceIdentity_r17", err)
	}
	ie.Value = aper.BitString{
		Bytes:   v,
		NumBits: uint64(n),
	}
	return nil
}
