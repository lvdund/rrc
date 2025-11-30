package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SpatialRelations_maxNumberConfiguredSpatialRelations_Enum_n4  aper.Enumerated = 0
	SpatialRelations_maxNumberConfiguredSpatialRelations_Enum_n8  aper.Enumerated = 1
	SpatialRelations_maxNumberConfiguredSpatialRelations_Enum_n16 aper.Enumerated = 2
	SpatialRelations_maxNumberConfiguredSpatialRelations_Enum_n32 aper.Enumerated = 3
	SpatialRelations_maxNumberConfiguredSpatialRelations_Enum_n64 aper.Enumerated = 4
	SpatialRelations_maxNumberConfiguredSpatialRelations_Enum_n96 aper.Enumerated = 5
)

type SpatialRelations_maxNumberConfiguredSpatialRelations struct {
	Value aper.Enumerated `lb:0,ub:5,madatory`
}

func (ie *SpatialRelations_maxNumberConfiguredSpatialRelations) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Encode SpatialRelations_maxNumberConfiguredSpatialRelations", err)
	}
	return nil
}

func (ie *SpatialRelations_maxNumberConfiguredSpatialRelations) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Decode SpatialRelations_maxNumberConfiguredSpatialRelations", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
