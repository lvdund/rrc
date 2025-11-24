package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ReportSFTD_EUTRA struct {
	ReportSFTD_Meas bool `madatory`
	ReportRSRP      bool `madatory`
}

func (ie *ReportSFTD_EUTRA) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteBoolean(ie.ReportSFTD_Meas); err != nil {
		return utils.WrapError("WriteBoolean ReportSFTD_Meas", err)
	}
	if err = w.WriteBoolean(ie.ReportRSRP); err != nil {
		return utils.WrapError("WriteBoolean ReportRSRP", err)
	}
	return nil
}

func (ie *ReportSFTD_EUTRA) Decode(r *uper.UperReader) error {
	var err error
	var tmp_bool_ReportSFTD_Meas bool
	if tmp_bool_ReportSFTD_Meas, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean ReportSFTD_Meas", err)
	}
	ie.ReportSFTD_Meas = tmp_bool_ReportSFTD_Meas
	var tmp_bool_ReportRSRP bool
	if tmp_bool_ReportRSRP, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean ReportRSRP", err)
	}
	ie.ReportRSRP = tmp_bool_ReportRSRP
	return nil
}
