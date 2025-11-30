package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CG_Config_v1540_IEs struct {
	PSCellFrequency      *ARFCN_ValueNR                           `optional`
	ReportCGI_RequestNR  *CG_Config_v1540_IEs_reportCGI_RequestNR `optional`
	Ph_InfoSCG           *PH_TypeListSCG                          `optional`
	NonCriticalExtension *CG_Config_v1560_IEs                     `optional`
}

func (ie *CG_Config_v1540_IEs) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.PSCellFrequency != nil, ie.ReportCGI_RequestNR != nil, ie.Ph_InfoSCG != nil, ie.NonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.PSCellFrequency != nil {
		if err = ie.PSCellFrequency.Encode(w); err != nil {
			return utils.WrapError("Encode PSCellFrequency", err)
		}
	}
	if ie.ReportCGI_RequestNR != nil {
		if err = ie.ReportCGI_RequestNR.Encode(w); err != nil {
			return utils.WrapError("Encode ReportCGI_RequestNR", err)
		}
	}
	if ie.Ph_InfoSCG != nil {
		if err = ie.Ph_InfoSCG.Encode(w); err != nil {
			return utils.WrapError("Encode Ph_InfoSCG", err)
		}
	}
	if ie.NonCriticalExtension != nil {
		if err = ie.NonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode NonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *CG_Config_v1540_IEs) Decode(r *aper.AperReader) error {
	var err error
	var PSCellFrequencyPresent bool
	if PSCellFrequencyPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ReportCGI_RequestNRPresent bool
	if ReportCGI_RequestNRPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Ph_InfoSCGPresent bool
	if Ph_InfoSCGPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var NonCriticalExtensionPresent bool
	if NonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if PSCellFrequencyPresent {
		ie.PSCellFrequency = new(ARFCN_ValueNR)
		if err = ie.PSCellFrequency.Decode(r); err != nil {
			return utils.WrapError("Decode PSCellFrequency", err)
		}
	}
	if ReportCGI_RequestNRPresent {
		ie.ReportCGI_RequestNR = new(CG_Config_v1540_IEs_reportCGI_RequestNR)
		if err = ie.ReportCGI_RequestNR.Decode(r); err != nil {
			return utils.WrapError("Decode ReportCGI_RequestNR", err)
		}
	}
	if Ph_InfoSCGPresent {
		ie.Ph_InfoSCG = new(PH_TypeListSCG)
		if err = ie.Ph_InfoSCG.Decode(r); err != nil {
			return utils.WrapError("Decode Ph_InfoSCG", err)
		}
	}
	if NonCriticalExtensionPresent {
		ie.NonCriticalExtension = new(CG_Config_v1560_IEs)
		if err = ie.NonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode NonCriticalExtension", err)
		}
	}
	return nil
}
