package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CG_Config_v1540_IEs_reportCGI_RequestNR_requestedCellInfo struct {
	SsbFrequency            ARFCN_ValueNR `madatory`
	CellForWhichToReportCGI PhysCellId    `madatory`
}

func (ie *CG_Config_v1540_IEs_reportCGI_RequestNR_requestedCellInfo) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.SsbFrequency.Encode(w); err != nil {
		return utils.WrapError("Encode SsbFrequency", err)
	}
	if err = ie.CellForWhichToReportCGI.Encode(w); err != nil {
		return utils.WrapError("Encode CellForWhichToReportCGI", err)
	}
	return nil
}

func (ie *CG_Config_v1540_IEs_reportCGI_RequestNR_requestedCellInfo) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.SsbFrequency.Decode(r); err != nil {
		return utils.WrapError("Decode SsbFrequency", err)
	}
	if err = ie.CellForWhichToReportCGI.Decode(r); err != nil {
		return utils.WrapError("Decode CellForWhichToReportCGI", err)
	}
	return nil
}
