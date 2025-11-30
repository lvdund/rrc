package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MeasObjectId struct {
	Value uint64 `lb:1,ub:maxNrofObjectId,madatory`
}

func (ie *MeasObjectId) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &aper.Constraint{Lb: 1, Ub: maxNrofObjectId}, false); err != nil {
		return utils.WrapError("Encode MeasObjectId", err)
	}
	return nil
}

func (ie *MeasObjectId) Decode(r *aper.AperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: maxNrofObjectId}, false); err != nil {
		return utils.WrapError("Decode MeasObjectId", err)
	}
	ie.Value = uint64(v)
	return nil
}
