package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CG_ConfigInfo_v1540_IEs struct {
	Ph_InfoMCG           *PH_TypeListMCG                              `optional`
	MeasResultReportCGI  *CG_ConfigInfo_v1540_IEs_measResultReportCGI `optional`
	NonCriticalExtension *CG_ConfigInfo_v1560_IEs                     `optional`
}

func (ie *CG_ConfigInfo_v1540_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Ph_InfoMCG != nil, ie.MeasResultReportCGI != nil, ie.NonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Ph_InfoMCG != nil {
		if err = ie.Ph_InfoMCG.Encode(w); err != nil {
			return utils.WrapError("Encode Ph_InfoMCG", err)
		}
	}
	if ie.MeasResultReportCGI != nil {
		if err = ie.MeasResultReportCGI.Encode(w); err != nil {
			return utils.WrapError("Encode MeasResultReportCGI", err)
		}
	}
	if ie.NonCriticalExtension != nil {
		if err = ie.NonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode NonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *CG_ConfigInfo_v1540_IEs) Decode(r *uper.UperReader) error {
	var err error
	var Ph_InfoMCGPresent bool
	if Ph_InfoMCGPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var MeasResultReportCGIPresent bool
	if MeasResultReportCGIPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var NonCriticalExtensionPresent bool
	if NonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Ph_InfoMCGPresent {
		ie.Ph_InfoMCG = new(PH_TypeListMCG)
		if err = ie.Ph_InfoMCG.Decode(r); err != nil {
			return utils.WrapError("Decode Ph_InfoMCG", err)
		}
	}
	if MeasResultReportCGIPresent {
		ie.MeasResultReportCGI = new(CG_ConfigInfo_v1540_IEs_measResultReportCGI)
		if err = ie.MeasResultReportCGI.Decode(r); err != nil {
			return utils.WrapError("Decode MeasResultReportCGI", err)
		}
	}
	if NonCriticalExtensionPresent {
		ie.NonCriticalExtension = new(CG_ConfigInfo_v1560_IEs)
		if err = ie.NonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode NonCriticalExtension", err)
		}
	}
	return nil
}
