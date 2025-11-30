package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	PUCCH_SpatialRelationInfo_closedLoopIndex_Enum_i0 aper.Enumerated = 0
	PUCCH_SpatialRelationInfo_closedLoopIndex_Enum_i1 aper.Enumerated = 1
)

type PUCCH_SpatialRelationInfo_closedLoopIndex struct {
	Value aper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *PUCCH_SpatialRelationInfo_closedLoopIndex) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode PUCCH_SpatialRelationInfo_closedLoopIndex", err)
	}
	return nil
}

func (ie *PUCCH_SpatialRelationInfo_closedLoopIndex) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode PUCCH_SpatialRelationInfo_closedLoopIndex", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
