package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SL_EventTriggerConfig_r16_sl_EventId_r16_eventS1_r16 struct {
	S1_Threshold_r16     SL_MeasTriggerQuantity_r16 `madatory`
	Sl_ReportOnLeave_r16 bool                       `madatory`
	Sl_Hysteresis_r16    Hysteresis                 `madatory`
	Sl_TimeToTrigger_r16 TimeToTrigger              `madatory`
}

func (ie *SL_EventTriggerConfig_r16_sl_EventId_r16_eventS1_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.S1_Threshold_r16.Encode(w); err != nil {
		return utils.WrapError("Encode S1_Threshold_r16", err)
	}
	if err = w.WriteBoolean(ie.Sl_ReportOnLeave_r16); err != nil {
		return utils.WrapError("WriteBoolean Sl_ReportOnLeave_r16", err)
	}
	if err = ie.Sl_Hysteresis_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Sl_Hysteresis_r16", err)
	}
	if err = ie.Sl_TimeToTrigger_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Sl_TimeToTrigger_r16", err)
	}
	return nil
}

func (ie *SL_EventTriggerConfig_r16_sl_EventId_r16_eventS1_r16) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.S1_Threshold_r16.Decode(r); err != nil {
		return utils.WrapError("Decode S1_Threshold_r16", err)
	}
	var tmp_bool_Sl_ReportOnLeave_r16 bool
	if tmp_bool_Sl_ReportOnLeave_r16, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean Sl_ReportOnLeave_r16", err)
	}
	ie.Sl_ReportOnLeave_r16 = tmp_bool_Sl_ReportOnLeave_r16
	if err = ie.Sl_Hysteresis_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Sl_Hysteresis_r16", err)
	}
	if err = ie.Sl_TimeToTrigger_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Sl_TimeToTrigger_r16", err)
	}
	return nil
}
