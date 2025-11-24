package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type EventTriggerConfigInterRAT_eventId_eventB1 struct {
	B1_ThresholdEUTRA MeasTriggerQuantityEUTRA `madatory`
	ReportOnLeave     bool                     `madatory`
	Hysteresis        Hysteresis               `madatory`
	TimeToTrigger     TimeToTrigger            `madatory`
}

func (ie *EventTriggerConfigInterRAT_eventId_eventB1) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.B1_ThresholdEUTRA.Encode(w); err != nil {
		return utils.WrapError("Encode B1_ThresholdEUTRA", err)
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

func (ie *EventTriggerConfigInterRAT_eventId_eventB1) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.B1_ThresholdEUTRA.Decode(r); err != nil {
		return utils.WrapError("Decode B1_ThresholdEUTRA", err)
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
