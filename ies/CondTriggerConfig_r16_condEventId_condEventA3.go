package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CondTriggerConfig_r16_condEventId_condEventA3 struct {
	A3_Offset     MeasTriggerQuantityOffset `madatory`
	Hysteresis    Hysteresis                `madatory`
	TimeToTrigger TimeToTrigger             `madatory`
}

func (ie *CondTriggerConfig_r16_condEventId_condEventA3) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.A3_Offset.Encode(w); err != nil {
		return utils.WrapError("Encode A3_Offset", err)
	}
	if err = ie.Hysteresis.Encode(w); err != nil {
		return utils.WrapError("Encode Hysteresis", err)
	}
	if err = ie.TimeToTrigger.Encode(w); err != nil {
		return utils.WrapError("Encode TimeToTrigger", err)
	}
	return nil
}

func (ie *CondTriggerConfig_r16_condEventId_condEventA3) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.A3_Offset.Decode(r); err != nil {
		return utils.WrapError("Decode A3_Offset", err)
	}
	if err = ie.Hysteresis.Decode(r); err != nil {
		return utils.WrapError("Decode Hysteresis", err)
	}
	if err = ie.TimeToTrigger.Decode(r); err != nil {
		return utils.WrapError("Decode TimeToTrigger", err)
	}
	return nil
}
