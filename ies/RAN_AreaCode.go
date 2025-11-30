package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type RAN_AreaCode struct {
	Value uint64 `lb:0,ub:255,madatory`
}

func (ie *RAN_AreaCode) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &aper.Constraint{Lb: 0, Ub: 255}, false); err != nil {
		return utils.WrapError("Encode RAN_AreaCode", err)
	}
	return nil
}

func (ie *RAN_AreaCode) Decode(r *aper.AperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 255}, false); err != nil {
		return utils.WrapError("Decode RAN_AreaCode", err)
	}
	ie.Value = uint64(v)
	return nil
}
