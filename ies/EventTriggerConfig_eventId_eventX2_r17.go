package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type EventTriggerConfig_eventId_eventX2_r17 struct {
	X2_Threshold_Relay_r17 SL_MeasTriggerQuantity_r16 `madatory`
	ReportOnLeave_r17      bool                       `madatory`
	Hysteresis_r17         Hysteresis                 `madatory`
	TimeToTrigger_r17      TimeToTrigger              `madatory`
}

func (ie *EventTriggerConfig_eventId_eventX2_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.X2_Threshold_Relay_r17.Encode(w); err != nil {
		return utils.WrapError("Encode X2_Threshold_Relay_r17", err)
	}
	if err = w.WriteBoolean(ie.ReportOnLeave_r17); err != nil {
		return utils.WrapError("WriteBoolean ReportOnLeave_r17", err)
	}
	if err = ie.Hysteresis_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Hysteresis_r17", err)
	}
	if err = ie.TimeToTrigger_r17.Encode(w); err != nil {
		return utils.WrapError("Encode TimeToTrigger_r17", err)
	}
	return nil
}

func (ie *EventTriggerConfig_eventId_eventX2_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.X2_Threshold_Relay_r17.Decode(r); err != nil {
		return utils.WrapError("Decode X2_Threshold_Relay_r17", err)
	}
	var tmp_bool_ReportOnLeave_r17 bool
	if tmp_bool_ReportOnLeave_r17, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean ReportOnLeave_r17", err)
	}
	ie.ReportOnLeave_r17 = tmp_bool_ReportOnLeave_r17
	if err = ie.Hysteresis_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Hysteresis_r17", err)
	}
	if err = ie.TimeToTrigger_r17.Decode(r); err != nil {
		return utils.WrapError("Decode TimeToTrigger_r17", err)
	}
	return nil
}
