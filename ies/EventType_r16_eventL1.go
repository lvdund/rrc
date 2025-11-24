package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type EventType_r16_eventL1 struct {
	L1_Threshold  MeasTriggerQuantity `madatory`
	Hysteresis    Hysteresis          `madatory`
	TimeToTrigger TimeToTrigger       `madatory`
}

func (ie *EventType_r16_eventL1) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.L1_Threshold.Encode(w); err != nil {
		return utils.WrapError("Encode L1_Threshold", err)
	}
	if err = ie.Hysteresis.Encode(w); err != nil {
		return utils.WrapError("Encode Hysteresis", err)
	}
	if err = ie.TimeToTrigger.Encode(w); err != nil {
		return utils.WrapError("Encode TimeToTrigger", err)
	}
	return nil
}

func (ie *EventType_r16_eventL1) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.L1_Threshold.Decode(r); err != nil {
		return utils.WrapError("Decode L1_Threshold", err)
	}
	if err = ie.Hysteresis.Decode(r); err != nil {
		return utils.WrapError("Decode Hysteresis", err)
	}
	if err = ie.TimeToTrigger.Decode(r); err != nil {
		return utils.WrapError("Decode TimeToTrigger", err)
	}
	return nil
}
