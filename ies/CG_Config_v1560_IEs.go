package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CG_Config_v1560_IEs struct {
	PSCellFrequencyEUTRA          *ARFCN_ValueEUTRA                           `optional`
	Scg_CellGroupConfigEUTRA      *[]byte                                     `optional`
	CandidateCellInfoListSN_EUTRA *[]byte                                     `optional`
	CandidateServingFreqListEUTRA *CandidateServingFreqListEUTRA              `optional`
	NeedForGaps                   *CG_Config_v1560_IEs_needForGaps            `optional`
	Drx_ConfigSCG                 *DRX_Config                                 `optional`
	ReportCGI_RequestEUTRA        *CG_Config_v1560_IEs_reportCGI_RequestEUTRA `optional`
	NonCriticalExtension          *CG_Config_v1590_IEs                        `optional`
}

func (ie *CG_Config_v1560_IEs) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.PSCellFrequencyEUTRA != nil, ie.Scg_CellGroupConfigEUTRA != nil, ie.CandidateCellInfoListSN_EUTRA != nil, ie.CandidateServingFreqListEUTRA != nil, ie.NeedForGaps != nil, ie.Drx_ConfigSCG != nil, ie.ReportCGI_RequestEUTRA != nil, ie.NonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.PSCellFrequencyEUTRA != nil {
		if err = ie.PSCellFrequencyEUTRA.Encode(w); err != nil {
			return utils.WrapError("Encode PSCellFrequencyEUTRA", err)
		}
	}
	if ie.Scg_CellGroupConfigEUTRA != nil {
		if err = w.WriteOctetString(*ie.Scg_CellGroupConfigEUTRA, &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode Scg_CellGroupConfigEUTRA", err)
		}
	}
	if ie.CandidateCellInfoListSN_EUTRA != nil {
		if err = w.WriteOctetString(*ie.CandidateCellInfoListSN_EUTRA, &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode CandidateCellInfoListSN_EUTRA", err)
		}
	}
	if ie.CandidateServingFreqListEUTRA != nil {
		if err = ie.CandidateServingFreqListEUTRA.Encode(w); err != nil {
			return utils.WrapError("Encode CandidateServingFreqListEUTRA", err)
		}
	}
	if ie.NeedForGaps != nil {
		if err = ie.NeedForGaps.Encode(w); err != nil {
			return utils.WrapError("Encode NeedForGaps", err)
		}
	}
	if ie.Drx_ConfigSCG != nil {
		if err = ie.Drx_ConfigSCG.Encode(w); err != nil {
			return utils.WrapError("Encode Drx_ConfigSCG", err)
		}
	}
	if ie.ReportCGI_RequestEUTRA != nil {
		if err = ie.ReportCGI_RequestEUTRA.Encode(w); err != nil {
			return utils.WrapError("Encode ReportCGI_RequestEUTRA", err)
		}
	}
	if ie.NonCriticalExtension != nil {
		if err = ie.NonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode NonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *CG_Config_v1560_IEs) Decode(r *aper.AperReader) error {
	var err error
	var PSCellFrequencyEUTRAPresent bool
	if PSCellFrequencyEUTRAPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Scg_CellGroupConfigEUTRAPresent bool
	if Scg_CellGroupConfigEUTRAPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var CandidateCellInfoListSN_EUTRAPresent bool
	if CandidateCellInfoListSN_EUTRAPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var CandidateServingFreqListEUTRAPresent bool
	if CandidateServingFreqListEUTRAPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var NeedForGapsPresent bool
	if NeedForGapsPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Drx_ConfigSCGPresent bool
	if Drx_ConfigSCGPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ReportCGI_RequestEUTRAPresent bool
	if ReportCGI_RequestEUTRAPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var NonCriticalExtensionPresent bool
	if NonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if PSCellFrequencyEUTRAPresent {
		ie.PSCellFrequencyEUTRA = new(ARFCN_ValueEUTRA)
		if err = ie.PSCellFrequencyEUTRA.Decode(r); err != nil {
			return utils.WrapError("Decode PSCellFrequencyEUTRA", err)
		}
	}
	if Scg_CellGroupConfigEUTRAPresent {
		var tmp_os_Scg_CellGroupConfigEUTRA []byte
		if tmp_os_Scg_CellGroupConfigEUTRA, err = r.ReadOctetString(&aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode Scg_CellGroupConfigEUTRA", err)
		}
		ie.Scg_CellGroupConfigEUTRA = &tmp_os_Scg_CellGroupConfigEUTRA
	}
	if CandidateCellInfoListSN_EUTRAPresent {
		var tmp_os_CandidateCellInfoListSN_EUTRA []byte
		if tmp_os_CandidateCellInfoListSN_EUTRA, err = r.ReadOctetString(&aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode CandidateCellInfoListSN_EUTRA", err)
		}
		ie.CandidateCellInfoListSN_EUTRA = &tmp_os_CandidateCellInfoListSN_EUTRA
	}
	if CandidateServingFreqListEUTRAPresent {
		ie.CandidateServingFreqListEUTRA = new(CandidateServingFreqListEUTRA)
		if err = ie.CandidateServingFreqListEUTRA.Decode(r); err != nil {
			return utils.WrapError("Decode CandidateServingFreqListEUTRA", err)
		}
	}
	if NeedForGapsPresent {
		ie.NeedForGaps = new(CG_Config_v1560_IEs_needForGaps)
		if err = ie.NeedForGaps.Decode(r); err != nil {
			return utils.WrapError("Decode NeedForGaps", err)
		}
	}
	if Drx_ConfigSCGPresent {
		ie.Drx_ConfigSCG = new(DRX_Config)
		if err = ie.Drx_ConfigSCG.Decode(r); err != nil {
			return utils.WrapError("Decode Drx_ConfigSCG", err)
		}
	}
	if ReportCGI_RequestEUTRAPresent {
		ie.ReportCGI_RequestEUTRA = new(CG_Config_v1560_IEs_reportCGI_RequestEUTRA)
		if err = ie.ReportCGI_RequestEUTRA.Decode(r); err != nil {
			return utils.WrapError("Decode ReportCGI_RequestEUTRA", err)
		}
	}
	if NonCriticalExtensionPresent {
		ie.NonCriticalExtension = new(CG_Config_v1590_IEs)
		if err = ie.NonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode NonCriticalExtension", err)
		}
	}
	return nil
}
