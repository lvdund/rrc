package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type Hysteresis struct {
	Value uint64 `lb:0,ub:30,madatory`
}

func (ie *Hysteresis) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &aper.Constraint{Lb: 0, Ub: 30}, false); err != nil {
		return utils.WrapError("Encode Hysteresis", err)
	}
	return nil
}

func (ie *Hysteresis) Decode(r *aper.AperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 30}, false); err != nil {
		return utils.WrapError("Decode Hysteresis", err)
	}
	ie.Value = uint64(v)
	return nil
}
