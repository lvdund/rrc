package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	GroupB_ConfiguredTwoStepRA_r16_messagePowerOffsetGroupB_Enum_minusinfinity aper.Enumerated = 0
	GroupB_ConfiguredTwoStepRA_r16_messagePowerOffsetGroupB_Enum_dB0           aper.Enumerated = 1
	GroupB_ConfiguredTwoStepRA_r16_messagePowerOffsetGroupB_Enum_dB5           aper.Enumerated = 2
	GroupB_ConfiguredTwoStepRA_r16_messagePowerOffsetGroupB_Enum_dB8           aper.Enumerated = 3
	GroupB_ConfiguredTwoStepRA_r16_messagePowerOffsetGroupB_Enum_dB10          aper.Enumerated = 4
	GroupB_ConfiguredTwoStepRA_r16_messagePowerOffsetGroupB_Enum_dB12          aper.Enumerated = 5
	GroupB_ConfiguredTwoStepRA_r16_messagePowerOffsetGroupB_Enum_dB15          aper.Enumerated = 6
	GroupB_ConfiguredTwoStepRA_r16_messagePowerOffsetGroupB_Enum_dB18          aper.Enumerated = 7
)

type GroupB_ConfiguredTwoStepRA_r16_messagePowerOffsetGroupB struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *GroupB_ConfiguredTwoStepRA_r16_messagePowerOffsetGroupB) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode GroupB_ConfiguredTwoStepRA_r16_messagePowerOffsetGroupB", err)
	}
	return nil
}

func (ie *GroupB_ConfiguredTwoStepRA_r16_messagePowerOffsetGroupB) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode GroupB_ConfiguredTwoStepRA_r16_messagePowerOffsetGroupB", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
