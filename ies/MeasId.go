package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MeasId struct {
	Value uint64 `lb:1,ub:maxNrofMeasId,madatory`
}

func (ie *MeasId) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &aper.Constraint{Lb: 1, Ub: maxNrofMeasId}, false); err != nil {
		return utils.WrapError("Encode MeasId", err)
	}
	return nil
}

func (ie *MeasId) Decode(r *aper.AperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: maxNrofMeasId}, false); err != nil {
		return utils.WrapError("Decode MeasId", err)
	}
	ie.Value = uint64(v)
	return nil
}
