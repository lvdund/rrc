package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type ControlResourceSetId_r16 struct {
	Value uint64 `lb:0,ub:maxNrofControlResourceSets_1_r16,madatory`
}

func (ie *ControlResourceSetId_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &aper.Constraint{Lb: 0, Ub: maxNrofControlResourceSets_1_r16}, false); err != nil {
		return utils.WrapError("Encode ControlResourceSetId_r16", err)
	}
	return nil
}

func (ie *ControlResourceSetId_r16) Decode(r *aper.AperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: maxNrofControlResourceSets_1_r16}, false); err != nil {
		return utils.WrapError("Decode ControlResourceSetId_r16", err)
	}
	ie.Value = uint64(v)
	return nil
}
