package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FeatureSetUplink_v1610_ul_IntraUE_Mux_r16 struct {
	Pusch_PreparationLowPriority_r16  FeatureSetUplink_v1610_ul_IntraUE_Mux_r16_pusch_PreparationLowPriority_r16  `madatory`
	Pusch_PreparationHighPriority_r16 FeatureSetUplink_v1610_ul_IntraUE_Mux_r16_pusch_PreparationHighPriority_r16 `madatory`
}

func (ie *FeatureSetUplink_v1610_ul_IntraUE_Mux_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.Pusch_PreparationLowPriority_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Pusch_PreparationLowPriority_r16", err)
	}
	if err = ie.Pusch_PreparationHighPriority_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Pusch_PreparationHighPriority_r16", err)
	}
	return nil
}

func (ie *FeatureSetUplink_v1610_ul_IntraUE_Mux_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.Pusch_PreparationLowPriority_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Pusch_PreparationLowPriority_r16", err)
	}
	if err = ie.Pusch_PreparationHighPriority_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Pusch_PreparationHighPriority_r16", err)
	}
	return nil
}
