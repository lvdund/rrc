package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CondTriggerConfig_r16_condEventId_condEventA5 struct {
	A5_Threshold1 MeasTriggerQuantity `madatory`
	A5_Threshold2 MeasTriggerQuantity `madatory`
	Hysteresis    Hysteresis          `madatory`
	TimeToTrigger TimeToTrigger       `madatory`
}

func (ie *CondTriggerConfig_r16_condEventId_condEventA5) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.A5_Threshold1.Encode(w); err != nil {
		return utils.WrapError("Encode A5_Threshold1", err)
	}
	if err = ie.A5_Threshold2.Encode(w); err != nil {
		return utils.WrapError("Encode A5_Threshold2", err)
	}
	if err = ie.Hysteresis.Encode(w); err != nil {
		return utils.WrapError("Encode Hysteresis", err)
	}
	if err = ie.TimeToTrigger.Encode(w); err != nil {
		return utils.WrapError("Encode TimeToTrigger", err)
	}
	return nil
}

func (ie *CondTriggerConfig_r16_condEventId_condEventA5) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.A5_Threshold1.Decode(r); err != nil {
		return utils.WrapError("Decode A5_Threshold1", err)
	}
	if err = ie.A5_Threshold2.Decode(r); err != nil {
		return utils.WrapError("Decode A5_Threshold2", err)
	}
	if err = ie.Hysteresis.Decode(r); err != nil {
		return utils.WrapError("Decode Hysteresis", err)
	}
	if err = ie.TimeToTrigger.Decode(r); err != nil {
		return utils.WrapError("Decode TimeToTrigger", err)
	}
	return nil
}
