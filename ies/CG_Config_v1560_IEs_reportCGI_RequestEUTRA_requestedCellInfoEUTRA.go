package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CG_Config_v1560_IEs_reportCGI_RequestEUTRA_requestedCellInfoEUTRA struct {
	EutraFrequency                ARFCN_ValueEUTRA `madatory`
	CellForWhichToReportCGI_EUTRA EUTRA_PhysCellId `madatory`
}

func (ie *CG_Config_v1560_IEs_reportCGI_RequestEUTRA_requestedCellInfoEUTRA) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.EutraFrequency.Encode(w); err != nil {
		return utils.WrapError("Encode EutraFrequency", err)
	}
	if err = ie.CellForWhichToReportCGI_EUTRA.Encode(w); err != nil {
		return utils.WrapError("Encode CellForWhichToReportCGI_EUTRA", err)
	}
	return nil
}

func (ie *CG_Config_v1560_IEs_reportCGI_RequestEUTRA_requestedCellInfoEUTRA) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.EutraFrequency.Decode(r); err != nil {
		return utils.WrapError("Decode EutraFrequency", err)
	}
	if err = ie.CellForWhichToReportCGI_EUTRA.Decode(r); err != nil {
		return utils.WrapError("Decode CellForWhichToReportCGI_EUTRA", err)
	}
	return nil
}
