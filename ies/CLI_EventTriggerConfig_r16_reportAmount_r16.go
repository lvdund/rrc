package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	CLI_EventTriggerConfig_r16_reportAmount_r16_Enum_r1       aper.Enumerated = 0
	CLI_EventTriggerConfig_r16_reportAmount_r16_Enum_r2       aper.Enumerated = 1
	CLI_EventTriggerConfig_r16_reportAmount_r16_Enum_r4       aper.Enumerated = 2
	CLI_EventTriggerConfig_r16_reportAmount_r16_Enum_r8       aper.Enumerated = 3
	CLI_EventTriggerConfig_r16_reportAmount_r16_Enum_r16      aper.Enumerated = 4
	CLI_EventTriggerConfig_r16_reportAmount_r16_Enum_r32      aper.Enumerated = 5
	CLI_EventTriggerConfig_r16_reportAmount_r16_Enum_r64      aper.Enumerated = 6
	CLI_EventTriggerConfig_r16_reportAmount_r16_Enum_infinity aper.Enumerated = 7
)

type CLI_EventTriggerConfig_r16_reportAmount_r16 struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *CLI_EventTriggerConfig_r16_reportAmount_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode CLI_EventTriggerConfig_r16_reportAmount_r16", err)
	}
	return nil
}

func (ie *CLI_EventTriggerConfig_r16_reportAmount_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode CLI_EventTriggerConfig_r16_reportAmount_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
