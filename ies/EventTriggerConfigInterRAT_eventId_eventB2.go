package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type EventTriggerConfigInterRAT_eventId_eventB2 struct {
	B2_Threshold1      MeasTriggerQuantity      `madatory`
	B2_Threshold2EUTRA MeasTriggerQuantityEUTRA `madatory`
	ReportOnLeave      bool                     `madatory`
	Hysteresis         Hysteresis               `madatory`
	TimeToTrigger      TimeToTrigger            `madatory`
}

func (ie *EventTriggerConfigInterRAT_eventId_eventB2) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.B2_Threshold1.Encode(w); err != nil {
		return utils.WrapError("Encode B2_Threshold1", err)
	}
	if err = ie.B2_Threshold2EUTRA.Encode(w); err != nil {
		return utils.WrapError("Encode B2_Threshold2EUTRA", err)
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

func (ie *EventTriggerConfigInterRAT_eventId_eventB2) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.B2_Threshold1.Decode(r); err != nil {
		return utils.WrapError("Decode B2_Threshold1", err)
	}
	if err = ie.B2_Threshold2EUTRA.Decode(r); err != nil {
		return utils.WrapError("Decode B2_Threshold2EUTRA", err)
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
