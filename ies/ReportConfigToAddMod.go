package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type ReportConfigToAddMod struct {
	ReportConfigId ReportConfigId                    `madatory`
	ReportConfig   ReportConfigToAddMod_reportConfig `madatory`
}

func (ie *ReportConfigToAddMod) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.ReportConfigId.Encode(w); err != nil {
		return utils.WrapError("Encode ReportConfigId", err)
	}
	if err = ie.ReportConfig.Encode(w); err != nil {
		return utils.WrapError("Encode ReportConfig", err)
	}
	return nil
}

func (ie *ReportConfigToAddMod) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.ReportConfigId.Decode(r); err != nil {
		return utils.WrapError("Decode ReportConfigId", err)
	}
	if err = ie.ReportConfig.Decode(r); err != nil {
		return utils.WrapError("Decode ReportConfig", err)
	}
	return nil
}
