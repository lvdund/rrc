package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type PUCCH_ResourceSetId struct {
	Value uint64 `lb:0,ub:maxNrofPUCCH_ResourceSets_1,madatory`
}

func (ie *PUCCH_ResourceSetId) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &aper.Constraint{Lb: 0, Ub: maxNrofPUCCH_ResourceSets_1}, false); err != nil {
		return utils.WrapError("Encode PUCCH_ResourceSetId", err)
	}
	return nil
}

func (ie *PUCCH_ResourceSetId) Decode(r *aper.AperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: maxNrofPUCCH_ResourceSets_1}, false); err != nil {
		return utils.WrapError("Decode PUCCH_ResourceSetId", err)
	}
	ie.Value = uint64(v)
	return nil
}
