package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SRS_SpatialRelationInfoPos_r16_servingRS_r16_referenceSignal_r16_Choice_nothing uint64 = iota
	SRS_SpatialRelationInfoPos_r16_servingRS_r16_referenceSignal_r16_Choice_Ssb_IndexServing_r16
	SRS_SpatialRelationInfoPos_r16_servingRS_r16_referenceSignal_r16_Choice_Csi_RS_IndexServing_r16
	SRS_SpatialRelationInfoPos_r16_servingRS_r16_referenceSignal_r16_Choice_Srs_SpatialRelation_r16
)

type SRS_SpatialRelationInfoPos_r16_servingRS_r16_referenceSignal_r16 struct {
	Choice                  uint64
	Ssb_IndexServing_r16    *SSB_Index
	Csi_RS_IndexServing_r16 *NZP_CSI_RS_ResourceId
	Srs_SpatialRelation_r16 *SRS_SpatialRelationInfoPos_r16_servingRS_r16_referenceSignal_r16_srs_SpatialRelation_r16
}

func (ie *SRS_SpatialRelationInfoPos_r16_servingRS_r16_referenceSignal_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 3, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SRS_SpatialRelationInfoPos_r16_servingRS_r16_referenceSignal_r16_Choice_Ssb_IndexServing_r16:
		if err = ie.Ssb_IndexServing_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode Ssb_IndexServing_r16", err)
		}
	case SRS_SpatialRelationInfoPos_r16_servingRS_r16_referenceSignal_r16_Choice_Csi_RS_IndexServing_r16:
		if err = ie.Csi_RS_IndexServing_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode Csi_RS_IndexServing_r16", err)
		}
	case SRS_SpatialRelationInfoPos_r16_servingRS_r16_referenceSignal_r16_Choice_Srs_SpatialRelation_r16:
		if err = ie.Srs_SpatialRelation_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode Srs_SpatialRelation_r16", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *SRS_SpatialRelationInfoPos_r16_servingRS_r16_referenceSignal_r16) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(3, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SRS_SpatialRelationInfoPos_r16_servingRS_r16_referenceSignal_r16_Choice_Ssb_IndexServing_r16:
		ie.Ssb_IndexServing_r16 = new(SSB_Index)
		if err = ie.Ssb_IndexServing_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Ssb_IndexServing_r16", err)
		}
	case SRS_SpatialRelationInfoPos_r16_servingRS_r16_referenceSignal_r16_Choice_Csi_RS_IndexServing_r16:
		ie.Csi_RS_IndexServing_r16 = new(NZP_CSI_RS_ResourceId)
		if err = ie.Csi_RS_IndexServing_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Csi_RS_IndexServing_r16", err)
		}
	case SRS_SpatialRelationInfoPos_r16_servingRS_r16_referenceSignal_r16_Choice_Srs_SpatialRelation_r16:
		ie.Srs_SpatialRelation_r16 = new(SRS_SpatialRelationInfoPos_r16_servingRS_r16_referenceSignal_r16_srs_SpatialRelation_r16)
		if err = ie.Srs_SpatialRelation_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Srs_SpatialRelation_r16", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
