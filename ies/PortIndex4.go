package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type PortIndex4 struct {
	Value uint64 `lb:0,ub:3,madatory`
}

func (ie *PortIndex4) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode PortIndex4", err)
	}
	return nil
}

func (ie *PortIndex4) Decode(r *aper.AperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode PortIndex4", err)
	}
	ie.Value = uint64(v)
	return nil
}
