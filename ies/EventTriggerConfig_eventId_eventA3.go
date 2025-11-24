package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type EventTriggerConfig_eventId_eventA3 struct {
	A3_Offset          MeasTriggerQuantityOffset `madatory`
	ReportOnLeave      bool                      `madatory`
	Hysteresis         Hysteresis                `madatory`
	TimeToTrigger      TimeToTrigger             `madatory`
	UseAllowedCellList bool                      `madatory`
}

func (ie *EventTriggerConfig_eventId_eventA3) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.A3_Offset.Encode(w); err != nil {
		return utils.WrapError("Encode A3_Offset", err)
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
	if err = w.WriteBoolean(ie.UseAllowedCellList); err != nil {
		return utils.WrapError("WriteBoolean UseAllowedCellList", err)
	}
	return nil
}

func (ie *EventTriggerConfig_eventId_eventA3) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.A3_Offset.Decode(r); err != nil {
		return utils.WrapError("Decode A3_Offset", err)
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
	var tmp_bool_UseAllowedCellList bool
	if tmp_bool_UseAllowedCellList, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean UseAllowedCellList", err)
	}
	ie.UseAllowedCellList = tmp_bool_UseAllowedCellList
	return nil
}
