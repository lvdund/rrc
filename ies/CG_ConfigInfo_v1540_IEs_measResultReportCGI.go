package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CG_ConfigInfo_v1540_IEs_measResultReportCGI struct {
	SsbFrequency            ARFCN_ValueNR `madatory`
	CellForWhichToReportCGI PhysCellId    `madatory`
	Cgi_Info                CGI_InfoNR    `madatory`
}

func (ie *CG_ConfigInfo_v1540_IEs_measResultReportCGI) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.SsbFrequency.Encode(w); err != nil {
		return utils.WrapError("Encode SsbFrequency", err)
	}
	if err = ie.CellForWhichToReportCGI.Encode(w); err != nil {
		return utils.WrapError("Encode CellForWhichToReportCGI", err)
	}
	if err = ie.Cgi_Info.Encode(w); err != nil {
		return utils.WrapError("Encode Cgi_Info", err)
	}
	return nil
}

func (ie *CG_ConfigInfo_v1540_IEs_measResultReportCGI) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.SsbFrequency.Decode(r); err != nil {
		return utils.WrapError("Decode SsbFrequency", err)
	}
	if err = ie.CellForWhichToReportCGI.Decode(r); err != nil {
		return utils.WrapError("Decode CellForWhichToReportCGI", err)
	}
	if err = ie.Cgi_Info.Decode(r); err != nil {
		return utils.WrapError("Decode Cgi_Info", err)
	}
	return nil
}
