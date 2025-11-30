package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SRS_SpatialRelationInfoPos_r16_servingRS_r16_referenceSignal_r16_srs_SpatialRelation_r16 struct {
	ResourceSelection_r16 SRS_SpatialRelationInfoPos_r16_servingRS_r16_referenceSignal_r16_srs_SpatialRelation_r16_resourceSelection_r16 `madatory`
	UplinkBWP_r16         BWP_Id                                                                                                         `madatory`
}

func (ie *SRS_SpatialRelationInfoPos_r16_servingRS_r16_referenceSignal_r16_srs_SpatialRelation_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.ResourceSelection_r16.Encode(w); err != nil {
		return utils.WrapError("Encode ResourceSelection_r16", err)
	}
	if err = ie.UplinkBWP_r16.Encode(w); err != nil {
		return utils.WrapError("Encode UplinkBWP_r16", err)
	}
	return nil
}

func (ie *SRS_SpatialRelationInfoPos_r16_servingRS_r16_referenceSignal_r16_srs_SpatialRelation_r16) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.ResourceSelection_r16.Decode(r); err != nil {
		return utils.WrapError("Decode ResourceSelection_r16", err)
	}
	if err = ie.UplinkBWP_r16.Decode(r); err != nil {
		return utils.WrapError("Decode UplinkBWP_r16", err)
	}
	return nil
}
