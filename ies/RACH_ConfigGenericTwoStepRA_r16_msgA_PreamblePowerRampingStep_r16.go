package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	RACH_ConfigGenericTwoStepRA_r16_msgA_PreamblePowerRampingStep_r16_Enum_dB0 aper.Enumerated = 0
	RACH_ConfigGenericTwoStepRA_r16_msgA_PreamblePowerRampingStep_r16_Enum_dB2 aper.Enumerated = 1
	RACH_ConfigGenericTwoStepRA_r16_msgA_PreamblePowerRampingStep_r16_Enum_dB4 aper.Enumerated = 2
	RACH_ConfigGenericTwoStepRA_r16_msgA_PreamblePowerRampingStep_r16_Enum_dB6 aper.Enumerated = 3
)

type RACH_ConfigGenericTwoStepRA_r16_msgA_PreamblePowerRampingStep_r16 struct {
	Value aper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *RACH_ConfigGenericTwoStepRA_r16_msgA_PreamblePowerRampingStep_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode RACH_ConfigGenericTwoStepRA_r16_msgA_PreamblePowerRampingStep_r16", err)
	}
	return nil
}

func (ie *RACH_ConfigGenericTwoStepRA_r16_msgA_PreamblePowerRampingStep_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode RACH_ConfigGenericTwoStepRA_r16_msgA_PreamblePowerRampingStep_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
