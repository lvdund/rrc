package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type PUCCH_SpatialRelationInfoId struct {
	Value uint64 `lb:1,ub:maxNrofSpatialRelationInfos,madatory`
}

func (ie *PUCCH_SpatialRelationInfoId) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &aper.Constraint{Lb: 1, Ub: maxNrofSpatialRelationInfos}, false); err != nil {
		return utils.WrapError("Encode PUCCH_SpatialRelationInfoId", err)
	}
	return nil
}

func (ie *PUCCH_SpatialRelationInfoId) Decode(r *aper.AperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: maxNrofSpatialRelationInfos}, false); err != nil {
		return utils.WrapError("Decode PUCCH_SpatialRelationInfoId", err)
	}
	ie.Value = uint64(v)
	return nil
}
