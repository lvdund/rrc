package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CLI_PeriodicalReportConfig_r16 struct {
	ReportInterval_r16    ReportInterval                                  `madatory`
	ReportAmount_r16      CLI_PeriodicalReportConfig_r16_reportAmount_r16 `madatory`
	ReportQuantityCLI_r16 MeasReportQuantityCLI_r16                       `madatory`
	MaxReportCLI_r16      int64                                           `lb:1,ub:maxCLI_Report_r16,madatory`
}

func (ie *CLI_PeriodicalReportConfig_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.ReportInterval_r16.Encode(w); err != nil {
		return utils.WrapError("Encode ReportInterval_r16", err)
	}
	if err = ie.ReportAmount_r16.Encode(w); err != nil {
		return utils.WrapError("Encode ReportAmount_r16", err)
	}
	if err = ie.ReportQuantityCLI_r16.Encode(w); err != nil {
		return utils.WrapError("Encode ReportQuantityCLI_r16", err)
	}
	if err = w.WriteInteger(ie.MaxReportCLI_r16, &aper.Constraint{Lb: 1, Ub: maxCLI_Report_r16}, false); err != nil {
		return utils.WrapError("WriteInteger MaxReportCLI_r16", err)
	}
	return nil
}

func (ie *CLI_PeriodicalReportConfig_r16) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.ReportInterval_r16.Decode(r); err != nil {
		return utils.WrapError("Decode ReportInterval_r16", err)
	}
	if err = ie.ReportAmount_r16.Decode(r); err != nil {
		return utils.WrapError("Decode ReportAmount_r16", err)
	}
	if err = ie.ReportQuantityCLI_r16.Decode(r); err != nil {
		return utils.WrapError("Decode ReportQuantityCLI_r16", err)
	}
	var tmp_int_MaxReportCLI_r16 int64
	if tmp_int_MaxReportCLI_r16, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: maxCLI_Report_r16}, false); err != nil {
		return utils.WrapError("ReadInteger MaxReportCLI_r16", err)
	}
	ie.MaxReportCLI_r16 = tmp_int_MaxReportCLI_r16
	return nil
}
