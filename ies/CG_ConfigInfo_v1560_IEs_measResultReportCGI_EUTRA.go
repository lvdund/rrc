package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CG_ConfigInfo_v1560_IEs_measResultReportCGI_EUTRA struct {
	EutraFrequency                ARFCN_ValueEUTRA `madatory`
	CellForWhichToReportCGI_EUTRA EUTRA_PhysCellId `madatory`
	Cgi_InfoEUTRA                 CGI_InfoEUTRA    `madatory`
}

func (ie *CG_ConfigInfo_v1560_IEs_measResultReportCGI_EUTRA) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.EutraFrequency.Encode(w); err != nil {
		return utils.WrapError("Encode EutraFrequency", err)
	}
	if err = ie.CellForWhichToReportCGI_EUTRA.Encode(w); err != nil {
		return utils.WrapError("Encode CellForWhichToReportCGI_EUTRA", err)
	}
	if err = ie.Cgi_InfoEUTRA.Encode(w); err != nil {
		return utils.WrapError("Encode Cgi_InfoEUTRA", err)
	}
	return nil
}

func (ie *CG_ConfigInfo_v1560_IEs_measResultReportCGI_EUTRA) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.EutraFrequency.Decode(r); err != nil {
		return utils.WrapError("Decode EutraFrequency", err)
	}
	if err = ie.CellForWhichToReportCGI_EUTRA.Decode(r); err != nil {
		return utils.WrapError("Decode CellForWhichToReportCGI_EUTRA", err)
	}
	if err = ie.Cgi_InfoEUTRA.Decode(r); err != nil {
		return utils.WrapError("Decode Cgi_InfoEUTRA", err)
	}
	return nil
}
