package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CellIdentity struct {
	Value aper.BitString `lb:36,ub:36,madatory`
}

func (ie *CellIdentity) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteBitString(ie.Value.Bytes, uint(ie.Value.NumBits), &aper.Constraint{Lb: 36, Ub: 36}, false); err != nil {
		return utils.WrapError("Encode CellIdentity", err)
	}
	return nil
}

func (ie *CellIdentity) Decode(r *aper.AperReader) error {
	var err error
	var v []byte
	var n uint
	if v, n, err = r.ReadBitString(&aper.Constraint{Lb: 36, Ub: 36}, false); err != nil {
		return utils.WrapError("Decode CellIdentity", err)
	}
	ie.Value = aper.BitString{
		Bytes:   v,
		NumBits: uint64(n),
	}
	return nil
}
