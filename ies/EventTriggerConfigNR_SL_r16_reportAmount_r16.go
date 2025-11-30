package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	EventTriggerConfigNR_SL_r16_reportAmount_r16_Enum_r1       aper.Enumerated = 0
	EventTriggerConfigNR_SL_r16_reportAmount_r16_Enum_r2       aper.Enumerated = 1
	EventTriggerConfigNR_SL_r16_reportAmount_r16_Enum_r4       aper.Enumerated = 2
	EventTriggerConfigNR_SL_r16_reportAmount_r16_Enum_r8       aper.Enumerated = 3
	EventTriggerConfigNR_SL_r16_reportAmount_r16_Enum_r16      aper.Enumerated = 4
	EventTriggerConfigNR_SL_r16_reportAmount_r16_Enum_r32      aper.Enumerated = 5
	EventTriggerConfigNR_SL_r16_reportAmount_r16_Enum_r64      aper.Enumerated = 6
	EventTriggerConfigNR_SL_r16_reportAmount_r16_Enum_infinity aper.Enumerated = 7
)

type EventTriggerConfigNR_SL_r16_reportAmount_r16 struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *EventTriggerConfigNR_SL_r16_reportAmount_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode EventTriggerConfigNR_SL_r16_reportAmount_r16", err)
	}
	return nil
}

func (ie *EventTriggerConfigNR_SL_r16_reportAmount_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode EventTriggerConfigNR_SL_r16_reportAmount_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
