package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CSI_IM_ResourceSetId struct {
	Value uint64 `lb:0,ub:maxNrofCSI_IM_ResourceSets_1,madatory`
}

func (ie *CSI_IM_ResourceSetId) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &aper.Constraint{Lb: 0, Ub: maxNrofCSI_IM_ResourceSets_1}, false); err != nil {
		return utils.WrapError("Encode CSI_IM_ResourceSetId", err)
	}
	return nil
}

func (ie *CSI_IM_ResourceSetId) Decode(r *aper.AperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: maxNrofCSI_IM_ResourceSets_1}, false); err != nil {
		return utils.WrapError("Decode CSI_IM_ResourceSetId", err)
	}
	ie.Value = uint64(v)
	return nil
}
