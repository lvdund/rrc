package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type ReportConfigNR struct {
	ReportType ReportConfigNR_reportType `madatory`
}

func (ie *ReportConfigNR) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.ReportType.Encode(w); err != nil {
		return utils.WrapError("Encode ReportType", err)
	}
	return nil
}

func (ie *ReportConfigNR) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.ReportType.Decode(r); err != nil {
		return utils.WrapError("Decode ReportType", err)
	}
	return nil
}
