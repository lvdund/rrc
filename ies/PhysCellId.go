package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type PhysCellId struct {
	Value uint64 `lb:0,ub:1007,madatory`
}

func (ie *PhysCellId) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &aper.Constraint{Lb: 0, Ub: 1007}, false); err != nil {
		return utils.WrapError("Encode PhysCellId", err)
	}
	return nil
}

func (ie *PhysCellId) Decode(r *aper.AperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 1007}, false); err != nil {
		return utils.WrapError("Decode PhysCellId", err)
	}
	ie.Value = uint64(v)
	return nil
}
