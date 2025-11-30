package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type EventTriggerConfigInterRAT_eventId_eventB2_UTRA_FDD_r16 struct {
	B2_Threshold1_r16         MeasTriggerQuantity             `madatory`
	B2_Threshold2UTRA_FDD_r16 MeasTriggerQuantityUTRA_FDD_r16 `madatory`
	ReportOnLeave_r16         bool                            `madatory`
	Hysteresis_r16            Hysteresis                      `madatory`
	TimeToTrigger_r16         TimeToTrigger                   `madatory`
}

func (ie *EventTriggerConfigInterRAT_eventId_eventB2_UTRA_FDD_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.B2_Threshold1_r16.Encode(w); err != nil {
		return utils.WrapError("Encode B2_Threshold1_r16", err)
	}
	if err = ie.B2_Threshold2UTRA_FDD_r16.Encode(w); err != nil {
		return utils.WrapError("Encode B2_Threshold2UTRA_FDD_r16", err)
	}
	if err = w.WriteBoolean(ie.ReportOnLeave_r16); err != nil {
		return utils.WrapError("WriteBoolean ReportOnLeave_r16", err)
	}
	if err = ie.Hysteresis_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Hysteresis_r16", err)
	}
	if err = ie.TimeToTrigger_r16.Encode(w); err != nil {
		return utils.WrapError("Encode TimeToTrigger_r16", err)
	}
	return nil
}

func (ie *EventTriggerConfigInterRAT_eventId_eventB2_UTRA_FDD_r16) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.B2_Threshold1_r16.Decode(r); err != nil {
		return utils.WrapError("Decode B2_Threshold1_r16", err)
	}
	if err = ie.B2_Threshold2UTRA_FDD_r16.Decode(r); err != nil {
		return utils.WrapError("Decode B2_Threshold2UTRA_FDD_r16", err)
	}
	var tmp_bool_ReportOnLeave_r16 bool
	if tmp_bool_ReportOnLeave_r16, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean ReportOnLeave_r16", err)
	}
	ie.ReportOnLeave_r16 = tmp_bool_ReportOnLeave_r16
	if err = ie.Hysteresis_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Hysteresis_r16", err)
	}
	if err = ie.TimeToTrigger_r16.Decode(r); err != nil {
		return utils.WrapError("Decode TimeToTrigger_r16", err)
	}
	return nil
}
