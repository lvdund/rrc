package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type FeatureSetDownlink_v1540_pdsch_ProcessingType2_Limited struct {
	DifferentTB_PerSlot_SCS_30kHz FeatureSetDownlink_v1540_pdsch_ProcessingType2_Limited_differentTB_PerSlot_SCS_30kHz `madatory`
}

func (ie *FeatureSetDownlink_v1540_pdsch_ProcessingType2_Limited) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.DifferentTB_PerSlot_SCS_30kHz.Encode(w); err != nil {
		return utils.WrapError("Encode DifferentTB_PerSlot_SCS_30kHz", err)
	}
	return nil
}

func (ie *FeatureSetDownlink_v1540_pdsch_ProcessingType2_Limited) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.DifferentTB_PerSlot_SCS_30kHz.Decode(r); err != nil {
		return utils.WrapError("Decode DifferentTB_PerSlot_SCS_30kHz", err)
	}
	return nil
}
