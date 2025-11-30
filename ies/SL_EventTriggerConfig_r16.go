package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SL_EventTriggerConfig_r16 struct {
	Sl_EventId_r16        SL_EventTriggerConfig_r16_sl_EventId_r16      `madatory`
	Sl_ReportInterval_r16 ReportInterval                                `madatory,ext`
	Sl_ReportAmount_r16   SL_EventTriggerConfig_r16_sl_ReportAmount_r16 `madatory,ext`
	Sl_ReportQuantity_r16 SL_MeasReportQuantity_r16                     `madatory,ext`
	Sl_RS_Type_r16        SL_RS_Type_r16                                `madatory,ext`
}

func (ie *SL_EventTriggerConfig_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.Sl_EventId_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Sl_EventId_r16", err)
	}
	return nil
}

func (ie *SL_EventTriggerConfig_r16) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.Sl_EventId_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Sl_EventId_r16", err)
	}
	return nil
}
