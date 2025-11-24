package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type EventTriggerConfig_eventId_eventA2 struct {
	A2_Threshold  MeasTriggerQuantity `madatory`
	ReportOnLeave bool                `madatory`
	Hysteresis    Hysteresis          `madatory`
	TimeToTrigger TimeToTrigger       `madatory`
}

func (ie *EventTriggerConfig_eventId_eventA2) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.A2_Threshold.Encode(w); err != nil {
		return utils.WrapError("Encode A2_Threshold", err)
	}
	if err = w.WriteBoolean(ie.ReportOnLeave); err != nil {
		return utils.WrapError("WriteBoolean ReportOnLeave", err)
	}
	if err = ie.Hysteresis.Encode(w); err != nil {
		return utils.WrapError("Encode Hysteresis", err)
	}
	if err = ie.TimeToTrigger.Encode(w); err != nil {
		return utils.WrapError("Encode TimeToTrigger", err)
	}
	return nil
}

func (ie *EventTriggerConfig_eventId_eventA2) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.A2_Threshold.Decode(r); err != nil {
		return utils.WrapError("Decode A2_Threshold", err)
	}
	var tmp_bool_ReportOnLeave bool
	if tmp_bool_ReportOnLeave, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean ReportOnLeave", err)
	}
	ie.ReportOnLeave = tmp_bool_ReportOnLeave
	if err = ie.Hysteresis.Decode(r); err != nil {
		return utils.WrapError("Decode Hysteresis", err)
	}
	if err = ie.TimeToTrigger.Decode(r); err != nil {
		return utils.WrapError("Decode TimeToTrigger", err)
	}
	return nil
}
