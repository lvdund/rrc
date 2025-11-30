package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SRS_SpatialRelationInfoPos_r16_servingRS_r16_referenceSignal_r16_srs_SpatialRelation_r16_resourceSelection_r16_Choice_nothing uint64 = iota
	SRS_SpatialRelationInfoPos_r16_servingRS_r16_referenceSignal_r16_srs_SpatialRelation_r16_resourceSelection_r16_Choice_Srs_ResourceId_r16
	SRS_SpatialRelationInfoPos_r16_servingRS_r16_referenceSignal_r16_srs_SpatialRelation_r16_resourceSelection_r16_Choice_Srs_PosResourceId_r16
)

type SRS_SpatialRelationInfoPos_r16_servingRS_r16_referenceSignal_r16_srs_SpatialRelation_r16_resourceSelection_r16 struct {
	Choice                uint64
	Srs_ResourceId_r16    *SRS_ResourceId
	Srs_PosResourceId_r16 *SRS_PosResourceId_r16
}

func (ie *SRS_SpatialRelationInfoPos_r16_servingRS_r16_referenceSignal_r16_srs_SpatialRelation_r16_resourceSelection_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SRS_SpatialRelationInfoPos_r16_servingRS_r16_referenceSignal_r16_srs_SpatialRelation_r16_resourceSelection_r16_Choice_Srs_ResourceId_r16:
		if err = ie.Srs_ResourceId_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode Srs_ResourceId_r16", err)
		}
	case SRS_SpatialRelationInfoPos_r16_servingRS_r16_referenceSignal_r16_srs_SpatialRelation_r16_resourceSelection_r16_Choice_Srs_PosResourceId_r16:
		if err = ie.Srs_PosResourceId_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode Srs_PosResourceId_r16", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *SRS_SpatialRelationInfoPos_r16_servingRS_r16_referenceSignal_r16_srs_SpatialRelation_r16_resourceSelection_r16) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SRS_SpatialRelationInfoPos_r16_servingRS_r16_referenceSignal_r16_srs_SpatialRelation_r16_resourceSelection_r16_Choice_Srs_ResourceId_r16:
		ie.Srs_ResourceId_r16 = new(SRS_ResourceId)
		if err = ie.Srs_ResourceId_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Srs_ResourceId_r16", err)
		}
	case SRS_SpatialRelationInfoPos_r16_servingRS_r16_referenceSignal_r16_srs_SpatialRelation_r16_resourceSelection_r16_Choice_Srs_PosResourceId_r16:
		ie.Srs_PosResourceId_r16 = new(SRS_PosResourceId_r16)
		if err = ie.Srs_PosResourceId_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Srs_PosResourceId_r16", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
