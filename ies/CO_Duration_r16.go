package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CO_Duration_r16 struct {
	Value uint64 `lb:0,ub:1120,madatory`
}

func (ie *CO_Duration_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &aper.Constraint{Lb: 0, Ub: 1120}, false); err != nil {
		return utils.WrapError("Encode CO_Duration_r16", err)
	}
	return nil
}

func (ie *CO_Duration_r16) Decode(r *aper.AperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 1120}, false); err != nil {
		return utils.WrapError("Decode CO_Duration_r16", err)
	}
	ie.Value = uint64(v)
	return nil
}
