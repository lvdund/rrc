package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CFRA_TwoStep_r16_occasionsTwoStepRA_r16 struct {
	Rach_ConfigGenericTwoStepRA_r16   RACH_ConfigGenericTwoStepRA_r16                                           `madatory`
	Ssb_PerRACH_OccasionTwoStepRA_r16 CFRA_TwoStep_r16_occasionsTwoStepRA_r16_ssb_PerRACH_OccasionTwoStepRA_r16 `madatory`
}

func (ie *CFRA_TwoStep_r16_occasionsTwoStepRA_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.Rach_ConfigGenericTwoStepRA_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Rach_ConfigGenericTwoStepRA_r16", err)
	}
	if err = ie.Ssb_PerRACH_OccasionTwoStepRA_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Ssb_PerRACH_OccasionTwoStepRA_r16", err)
	}
	return nil
}

func (ie *CFRA_TwoStep_r16_occasionsTwoStepRA_r16) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.Rach_ConfigGenericTwoStepRA_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Rach_ConfigGenericTwoStepRA_r16", err)
	}
	if err = ie.Ssb_PerRACH_OccasionTwoStepRA_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Ssb_PerRACH_OccasionTwoStepRA_r16", err)
	}
	return nil
}
