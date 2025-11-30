package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type FeatureSetUplink_v1610_pusch_RepetitionTypeB_r16 struct {
	MaxNumberPUSCH_Tx_r16 FeatureSetUplink_v1610_pusch_RepetitionTypeB_r16_maxNumberPUSCH_Tx_r16 `madatory`
	HoppingScheme_r16     FeatureSetUplink_v1610_pusch_RepetitionTypeB_r16_hoppingScheme_r16     `madatory`
}

func (ie *FeatureSetUplink_v1610_pusch_RepetitionTypeB_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.MaxNumberPUSCH_Tx_r16.Encode(w); err != nil {
		return utils.WrapError("Encode MaxNumberPUSCH_Tx_r16", err)
	}
	if err = ie.HoppingScheme_r16.Encode(w); err != nil {
		return utils.WrapError("Encode HoppingScheme_r16", err)
	}
	return nil
}

func (ie *FeatureSetUplink_v1610_pusch_RepetitionTypeB_r16) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.MaxNumberPUSCH_Tx_r16.Decode(r); err != nil {
		return utils.WrapError("Decode MaxNumberPUSCH_Tx_r16", err)
	}
	if err = ie.HoppingScheme_r16.Decode(r); err != nil {
		return utils.WrapError("Decode HoppingScheme_r16", err)
	}
	return nil
}
