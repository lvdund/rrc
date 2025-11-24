package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CondTriggerConfig_r16_condEventId_condEventA4_r17 struct {
	A4_Threshold_r17  MeasTriggerQuantity `madatory`
	Hysteresis_r17    Hysteresis          `madatory`
	TimeToTrigger_r17 TimeToTrigger       `madatory`
}

func (ie *CondTriggerConfig_r16_condEventId_condEventA4_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.A4_Threshold_r17.Encode(w); err != nil {
		return utils.WrapError("Encode A4_Threshold_r17", err)
	}
	if err = ie.Hysteresis_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Hysteresis_r17", err)
	}
	if err = ie.TimeToTrigger_r17.Encode(w); err != nil {
		return utils.WrapError("Encode TimeToTrigger_r17", err)
	}
	return nil
}

func (ie *CondTriggerConfig_r16_condEventId_condEventA4_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.A4_Threshold_r17.Decode(r); err != nil {
		return utils.WrapError("Decode A4_Threshold_r17", err)
	}
	if err = ie.Hysteresis_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Hysteresis_r17", err)
	}
	if err = ie.TimeToTrigger_r17.Decode(r); err != nil {
		return utils.WrapError("Decode TimeToTrigger_r17", err)
	}
	return nil
}
