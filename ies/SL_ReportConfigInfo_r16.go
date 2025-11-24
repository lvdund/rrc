package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_ReportConfigInfo_r16 struct {
	Sl_ReportConfigId_r16 SL_ReportConfigId_r16 `madatory`
	Sl_ReportConfig_r16   SL_ReportConfig_r16   `madatory`
}

func (ie *SL_ReportConfigInfo_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.Sl_ReportConfigId_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Sl_ReportConfigId_r16", err)
	}
	if err = ie.Sl_ReportConfig_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Sl_ReportConfig_r16", err)
	}
	return nil
}

func (ie *SL_ReportConfigInfo_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.Sl_ReportConfigId_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Sl_ReportConfigId_r16", err)
	}
	if err = ie.Sl_ReportConfig_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Sl_ReportConfig_r16", err)
	}
	return nil
}
