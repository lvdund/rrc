package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SL_PeriodicalReportConfig_r16 struct {
	Sl_ReportInterval_r16 ReportInterval                                    `madatory`
	Sl_ReportAmount_r16   SL_PeriodicalReportConfig_r16_sl_ReportAmount_r16 `madatory`
	Sl_ReportQuantity_r16 SL_MeasReportQuantity_r16                         `madatory`
	Sl_RS_Type_r16        SL_RS_Type_r16                                    `madatory`
}

func (ie *SL_PeriodicalReportConfig_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.Sl_ReportInterval_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Sl_ReportInterval_r16", err)
	}
	if err = ie.Sl_ReportAmount_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Sl_ReportAmount_r16", err)
	}
	if err = ie.Sl_ReportQuantity_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Sl_ReportQuantity_r16", err)
	}
	if err = ie.Sl_RS_Type_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Sl_RS_Type_r16", err)
	}
	return nil
}

func (ie *SL_PeriodicalReportConfig_r16) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.Sl_ReportInterval_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Sl_ReportInterval_r16", err)
	}
	if err = ie.Sl_ReportAmount_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Sl_ReportAmount_r16", err)
	}
	if err = ie.Sl_ReportQuantity_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Sl_ReportQuantity_r16", err)
	}
	if err = ie.Sl_RS_Type_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Sl_RS_Type_r16", err)
	}
	return nil
}
