package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SpatialRelationsSRS_Pos_r16_spatialRelation_SRS_PosBasedOnSSB_Neigh_r16_Enum_supported aper.Enumerated = 0
)

type SpatialRelationsSRS_Pos_r16_spatialRelation_SRS_PosBasedOnSSB_Neigh_r16 struct {
	Value aper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *SpatialRelationsSRS_Pos_r16_spatialRelation_SRS_PosBasedOnSSB_Neigh_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode SpatialRelationsSRS_Pos_r16_spatialRelation_SRS_PosBasedOnSSB_Neigh_r16", err)
	}
	return nil
}

func (ie *SpatialRelationsSRS_Pos_r16_spatialRelation_SRS_PosBasedOnSSB_Neigh_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode SpatialRelationsSRS_Pos_r16_spatialRelation_SRS_PosBasedOnSSB_Neigh_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
