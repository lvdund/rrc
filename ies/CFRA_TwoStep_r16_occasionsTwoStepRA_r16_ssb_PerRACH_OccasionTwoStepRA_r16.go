package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	CFRA_TwoStep_r16_occasionsTwoStepRA_r16_ssb_PerRACH_OccasionTwoStepRA_r16_Enum_oneEighth aper.Enumerated = 0
	CFRA_TwoStep_r16_occasionsTwoStepRA_r16_ssb_PerRACH_OccasionTwoStepRA_r16_Enum_oneFourth aper.Enumerated = 1
	CFRA_TwoStep_r16_occasionsTwoStepRA_r16_ssb_PerRACH_OccasionTwoStepRA_r16_Enum_oneHalf   aper.Enumerated = 2
	CFRA_TwoStep_r16_occasionsTwoStepRA_r16_ssb_PerRACH_OccasionTwoStepRA_r16_Enum_one       aper.Enumerated = 3
	CFRA_TwoStep_r16_occasionsTwoStepRA_r16_ssb_PerRACH_OccasionTwoStepRA_r16_Enum_two       aper.Enumerated = 4
	CFRA_TwoStep_r16_occasionsTwoStepRA_r16_ssb_PerRACH_OccasionTwoStepRA_r16_Enum_four      aper.Enumerated = 5
	CFRA_TwoStep_r16_occasionsTwoStepRA_r16_ssb_PerRACH_OccasionTwoStepRA_r16_Enum_eight     aper.Enumerated = 6
	CFRA_TwoStep_r16_occasionsTwoStepRA_r16_ssb_PerRACH_OccasionTwoStepRA_r16_Enum_sixteen   aper.Enumerated = 7
)

type CFRA_TwoStep_r16_occasionsTwoStepRA_r16_ssb_PerRACH_OccasionTwoStepRA_r16 struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *CFRA_TwoStep_r16_occasionsTwoStepRA_r16_ssb_PerRACH_OccasionTwoStepRA_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode CFRA_TwoStep_r16_occasionsTwoStepRA_r16_ssb_PerRACH_OccasionTwoStepRA_r16", err)
	}
	return nil
}

func (ie *CFRA_TwoStep_r16_occasionsTwoStepRA_r16_ssb_PerRACH_OccasionTwoStepRA_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode CFRA_TwoStep_r16_occasionsTwoStepRA_r16_ssb_PerRACH_OccasionTwoStepRA_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
