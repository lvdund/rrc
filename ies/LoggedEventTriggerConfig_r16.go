package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type LoggedEventTriggerConfig_r16 struct {
	EventType_r16       EventType_r16       `madatory`
	LoggingInterval_r16 LoggingInterval_r16 `madatory`
}

func (ie *LoggedEventTriggerConfig_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.EventType_r16.Encode(w); err != nil {
		return utils.WrapError("Encode EventType_r16", err)
	}
	if err = ie.LoggingInterval_r16.Encode(w); err != nil {
		return utils.WrapError("Encode LoggingInterval_r16", err)
	}
	return nil
}

func (ie *LoggedEventTriggerConfig_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.EventType_r16.Decode(r); err != nil {
		return utils.WrapError("Decode EventType_r16", err)
	}
	if err = ie.LoggingInterval_r16.Decode(r); err != nil {
		return utils.WrapError("Decode LoggingInterval_r16", err)
	}
	return nil
}
