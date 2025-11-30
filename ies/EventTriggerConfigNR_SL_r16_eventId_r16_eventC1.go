package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type EventTriggerConfigNR_SL_r16_eventId_r16_eventC1 struct {
	C1_Threshold_r16  SL_CBR_r16    `madatory`
	Hysteresis_r16    Hysteresis    `madatory`
	TimeToTrigger_r16 TimeToTrigger `madatory`
}

func (ie *EventTriggerConfigNR_SL_r16_eventId_r16_eventC1) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.C1_Threshold_r16.Encode(w); err != nil {
		return utils.WrapError("Encode C1_Threshold_r16", err)
	}
	if err = ie.Hysteresis_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Hysteresis_r16", err)
	}
	if err = ie.TimeToTrigger_r16.Encode(w); err != nil {
		return utils.WrapError("Encode TimeToTrigger_r16", err)
	}
	return nil
}

func (ie *EventTriggerConfigNR_SL_r16_eventId_r16_eventC1) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.C1_Threshold_r16.Decode(r); err != nil {
		return utils.WrapError("Decode C1_Threshold_r16", err)
	}
	if err = ie.Hysteresis_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Hysteresis_r16", err)
	}
	if err = ie.TimeToTrigger_r16.Decode(r); err != nil {
		return utils.WrapError("Decode TimeToTrigger_r16", err)
	}
	return nil
}
