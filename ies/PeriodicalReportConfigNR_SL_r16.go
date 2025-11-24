package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PeriodicalReportConfigNR_SL_r16 struct {
	ReportInterval_r16 ReportInterval                                   `madatory`
	ReportAmount_r16   PeriodicalReportConfigNR_SL_r16_reportAmount_r16 `madatory`
	ReportQuantity_r16 MeasReportQuantity_r16                           `madatory`
}

func (ie *PeriodicalReportConfigNR_SL_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.ReportInterval_r16.Encode(w); err != nil {
		return utils.WrapError("Encode ReportInterval_r16", err)
	}
	if err = ie.ReportAmount_r16.Encode(w); err != nil {
		return utils.WrapError("Encode ReportAmount_r16", err)
	}
	if err = ie.ReportQuantity_r16.Encode(w); err != nil {
		return utils.WrapError("Encode ReportQuantity_r16", err)
	}
	return nil
}

func (ie *PeriodicalReportConfigNR_SL_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.ReportInterval_r16.Decode(r); err != nil {
		return utils.WrapError("Decode ReportInterval_r16", err)
	}
	if err = ie.ReportAmount_r16.Decode(r); err != nil {
		return utils.WrapError("Decode ReportAmount_r16", err)
	}
	if err = ie.ReportQuantity_r16.Decode(r); err != nil {
		return utils.WrapError("Decode ReportQuantity_r16", err)
	}
	return nil
}
