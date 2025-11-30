package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SRS_SpatialRelationInfoPos_r16_Choice_nothing uint64 = iota
	SRS_SpatialRelationInfoPos_r16_Choice_ServingRS_r16
	SRS_SpatialRelationInfoPos_r16_Choice_Ssb_Ncell_r16
	SRS_SpatialRelationInfoPos_r16_Choice_Dl_PRS_r16
)

type SRS_SpatialRelationInfoPos_r16 struct {
	Choice        uint64
	ServingRS_r16 *SRS_SpatialRelationInfoPos_r16_servingRS_r16
	Ssb_Ncell_r16 *SSB_InfoNcell_r16
	Dl_PRS_r16    *DL_PRS_Info_r16
}

func (ie *SRS_SpatialRelationInfoPos_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 3, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SRS_SpatialRelationInfoPos_r16_Choice_ServingRS_r16:
		if err = ie.ServingRS_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode ServingRS_r16", err)
		}
	case SRS_SpatialRelationInfoPos_r16_Choice_Ssb_Ncell_r16:
		if err = ie.Ssb_Ncell_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode Ssb_Ncell_r16", err)
		}
	case SRS_SpatialRelationInfoPos_r16_Choice_Dl_PRS_r16:
		if err = ie.Dl_PRS_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode Dl_PRS_r16", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *SRS_SpatialRelationInfoPos_r16) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(3, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SRS_SpatialRelationInfoPos_r16_Choice_ServingRS_r16:
		ie.ServingRS_r16 = new(SRS_SpatialRelationInfoPos_r16_servingRS_r16)
		if err = ie.ServingRS_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ServingRS_r16", err)
		}
	case SRS_SpatialRelationInfoPos_r16_Choice_Ssb_Ncell_r16:
		ie.Ssb_Ncell_r16 = new(SSB_InfoNcell_r16)
		if err = ie.Ssb_Ncell_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Ssb_Ncell_r16", err)
		}
	case SRS_SpatialRelationInfoPos_r16_Choice_Dl_PRS_r16:
		ie.Dl_PRS_r16 = new(DL_PRS_Info_r16)
		if err = ie.Dl_PRS_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Dl_PRS_r16", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
