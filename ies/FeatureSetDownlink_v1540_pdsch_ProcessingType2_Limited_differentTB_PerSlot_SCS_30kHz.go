package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	FeatureSetDownlink_v1540_pdsch_ProcessingType2_Limited_differentTB_PerSlot_SCS_30kHz_Enum_upto1 aper.Enumerated = 0
	FeatureSetDownlink_v1540_pdsch_ProcessingType2_Limited_differentTB_PerSlot_SCS_30kHz_Enum_upto2 aper.Enumerated = 1
	FeatureSetDownlink_v1540_pdsch_ProcessingType2_Limited_differentTB_PerSlot_SCS_30kHz_Enum_upto4 aper.Enumerated = 2
	FeatureSetDownlink_v1540_pdsch_ProcessingType2_Limited_differentTB_PerSlot_SCS_30kHz_Enum_upto7 aper.Enumerated = 3
)

type FeatureSetDownlink_v1540_pdsch_ProcessingType2_Limited_differentTB_PerSlot_SCS_30kHz struct {
	Value aper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *FeatureSetDownlink_v1540_pdsch_ProcessingType2_Limited_differentTB_PerSlot_SCS_30kHz) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode FeatureSetDownlink_v1540_pdsch_ProcessingType2_Limited_differentTB_PerSlot_SCS_30kHz", err)
	}
	return nil
}

func (ie *FeatureSetDownlink_v1540_pdsch_ProcessingType2_Limited_differentTB_PerSlot_SCS_30kHz) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode FeatureSetDownlink_v1540_pdsch_ProcessingType2_Limited_differentTB_PerSlot_SCS_30kHz", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
