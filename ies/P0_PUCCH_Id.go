package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type P0_PUCCH_Id struct {
	Value uint64 `lb:1,ub:8,madatory`
}

func (ie *P0_PUCCH_Id) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &aper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
		return utils.WrapError("Encode P0_PUCCH_Id", err)
	}
	return nil
}

func (ie *P0_PUCCH_Id) Decode(r *aper.AperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
		return utils.WrapError("Decode P0_PUCCH_Id", err)
	}
	ie.Value = uint64(v)
	return nil
}
