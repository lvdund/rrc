package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CLI_EventTriggerConfig_r16 struct {
	EventId_r16        CLI_EventTriggerConfig_r16_eventId_r16      `madatory`
	ReportInterval_r16 ReportInterval                              `madatory,ext`
	ReportAmount_r16   CLI_EventTriggerConfig_r16_reportAmount_r16 `madatory,ext`
	MaxReportCLI_r16   int64                                       `lb:1,ub:maxCLI_Report_r16,madatory,ext`
}

func (ie *CLI_EventTriggerConfig_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.EventId_r16.Encode(w); err != nil {
		return utils.WrapError("Encode EventId_r16", err)
	}
	return nil
}

func (ie *CLI_EventTriggerConfig_r16) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.EventId_r16.Decode(r); err != nil {
		return utils.WrapError("Decode EventId_r16", err)
	}
	return nil
}
