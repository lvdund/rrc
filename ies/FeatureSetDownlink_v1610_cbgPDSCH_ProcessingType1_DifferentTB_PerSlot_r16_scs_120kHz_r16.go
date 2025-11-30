package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	FeatureSetDownlink_v1610_cbgPDSCH_ProcessingType1_DifferentTB_PerSlot_r16_scs_120kHz_r16_Enum_one   aper.Enumerated = 0
	FeatureSetDownlink_v1610_cbgPDSCH_ProcessingType1_DifferentTB_PerSlot_r16_scs_120kHz_r16_Enum_upto2 aper.Enumerated = 1
	FeatureSetDownlink_v1610_cbgPDSCH_ProcessingType1_DifferentTB_PerSlot_r16_scs_120kHz_r16_Enum_upto4 aper.Enumerated = 2
	FeatureSetDownlink_v1610_cbgPDSCH_ProcessingType1_DifferentTB_PerSlot_r16_scs_120kHz_r16_Enum_upto7 aper.Enumerated = 3
)

type FeatureSetDownlink_v1610_cbgPDSCH_ProcessingType1_DifferentTB_PerSlot_r16_scs_120kHz_r16 struct {
	Value aper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *FeatureSetDownlink_v1610_cbgPDSCH_ProcessingType1_DifferentTB_PerSlot_r16_scs_120kHz_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode FeatureSetDownlink_v1610_cbgPDSCH_ProcessingType1_DifferentTB_PerSlot_r16_scs_120kHz_r16", err)
	}
	return nil
}

func (ie *FeatureSetDownlink_v1610_cbgPDSCH_ProcessingType1_DifferentTB_PerSlot_r16_scs_120kHz_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode FeatureSetDownlink_v1610_cbgPDSCH_ProcessingType1_DifferentTB_PerSlot_r16_scs_120kHz_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
