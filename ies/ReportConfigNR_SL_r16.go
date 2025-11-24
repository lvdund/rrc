package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ReportConfigNR_SL_r16 struct {
	ReportType_r16 ReportConfigNR_SL_r16_reportType_r16 `madatory`
}

func (ie *ReportConfigNR_SL_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.ReportType_r16.Encode(w); err != nil {
		return utils.WrapError("Encode ReportType_r16", err)
	}
	return nil
}

func (ie *ReportConfigNR_SL_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.ReportType_r16.Decode(r); err != nil {
		return utils.WrapError("Decode ReportType_r16", err)
	}
	return nil
}
