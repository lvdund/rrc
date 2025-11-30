package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type EventTriggerConfigNR_SL_r16 struct {
	EventId_r16        EventTriggerConfigNR_SL_r16_eventId_r16      `madatory`
	ReportInterval_r16 ReportInterval                               `madatory,ext`
	ReportAmount_r16   EventTriggerConfigNR_SL_r16_reportAmount_r16 `madatory,ext`
	ReportQuantity_r16 MeasReportQuantity_r16                       `madatory,ext`
}

func (ie *EventTriggerConfigNR_SL_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.EventId_r16.Encode(w); err != nil {
		return utils.WrapError("Encode EventId_r16", err)
	}
	return nil
}

func (ie *EventTriggerConfigNR_SL_r16) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.EventId_r16.Decode(r); err != nil {
		return utils.WrapError("Decode EventId_r16", err)
	}
	return nil
}
