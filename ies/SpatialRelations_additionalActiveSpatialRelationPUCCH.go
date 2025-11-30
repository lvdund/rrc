package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SpatialRelations_additionalActiveSpatialRelationPUCCH_Enum_supported aper.Enumerated = 0
)

type SpatialRelations_additionalActiveSpatialRelationPUCCH struct {
	Value aper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *SpatialRelations_additionalActiveSpatialRelationPUCCH) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode SpatialRelations_additionalActiveSpatialRelationPUCCH", err)
	}
	return nil
}

func (ie *SpatialRelations_additionalActiveSpatialRelationPUCCH) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode SpatialRelations_additionalActiveSpatialRelationPUCCH", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
