package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CG_ConfigInfo_v1560_IEs struct {
	CandidateCellInfoListMN_EUTRA *[]byte                                            `optional`
	CandidateCellInfoListSN_EUTRA *[]byte                                            `optional`
	SourceConfigSCG_EUTRA         *[]byte                                            `optional`
	ScgFailureInfoEUTRA           *CG_ConfigInfo_v1560_IEs_scgFailureInfoEUTRA       `optional`
	Drx_ConfigMCG                 *DRX_Config                                        `optional`
	MeasResultReportCGI_EUTRA     *CG_ConfigInfo_v1560_IEs_measResultReportCGI_EUTRA `optional`
	MeasResultCellListSFTD_EUTRA  *MeasResultCellListSFTD_EUTRA                      `optional`
	Fr_InfoListMCG                *FR_InfoList                                       `optional`
	NonCriticalExtension          *CG_ConfigInfo_v1570_IEs                           `optional`
}

func (ie *CG_ConfigInfo_v1560_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.CandidateCellInfoListMN_EUTRA != nil, ie.CandidateCellInfoListSN_EUTRA != nil, ie.SourceConfigSCG_EUTRA != nil, ie.ScgFailureInfoEUTRA != nil, ie.Drx_ConfigMCG != nil, ie.MeasResultReportCGI_EUTRA != nil, ie.MeasResultCellListSFTD_EUTRA != nil, ie.Fr_InfoListMCG != nil, ie.NonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.CandidateCellInfoListMN_EUTRA != nil {
		if err = w.WriteOctetString(*ie.CandidateCellInfoListMN_EUTRA, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode CandidateCellInfoListMN_EUTRA", err)
		}
	}
	if ie.CandidateCellInfoListSN_EUTRA != nil {
		if err = w.WriteOctetString(*ie.CandidateCellInfoListSN_EUTRA, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode CandidateCellInfoListSN_EUTRA", err)
		}
	}
	if ie.SourceConfigSCG_EUTRA != nil {
		if err = w.WriteOctetString(*ie.SourceConfigSCG_EUTRA, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode SourceConfigSCG_EUTRA", err)
		}
	}
	if ie.ScgFailureInfoEUTRA != nil {
		if err = ie.ScgFailureInfoEUTRA.Encode(w); err != nil {
			return utils.WrapError("Encode ScgFailureInfoEUTRA", err)
		}
	}
	if ie.Drx_ConfigMCG != nil {
		if err = ie.Drx_ConfigMCG.Encode(w); err != nil {
			return utils.WrapError("Encode Drx_ConfigMCG", err)
		}
	}
	if ie.MeasResultReportCGI_EUTRA != nil {
		if err = ie.MeasResultReportCGI_EUTRA.Encode(w); err != nil {
			return utils.WrapError("Encode MeasResultReportCGI_EUTRA", err)
		}
	}
	if ie.MeasResultCellListSFTD_EUTRA != nil {
		if err = ie.MeasResultCellListSFTD_EUTRA.Encode(w); err != nil {
			return utils.WrapError("Encode MeasResultCellListSFTD_EUTRA", err)
		}
	}
	if ie.Fr_InfoListMCG != nil {
		if err = ie.Fr_InfoListMCG.Encode(w); err != nil {
			return utils.WrapError("Encode Fr_InfoListMCG", err)
		}
	}
	if ie.NonCriticalExtension != nil {
		if err = ie.NonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode NonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *CG_ConfigInfo_v1560_IEs) Decode(r *uper.UperReader) error {
	var err error
	var CandidateCellInfoListMN_EUTRAPresent bool
	if CandidateCellInfoListMN_EUTRAPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var CandidateCellInfoListSN_EUTRAPresent bool
	if CandidateCellInfoListSN_EUTRAPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var SourceConfigSCG_EUTRAPresent bool
	if SourceConfigSCG_EUTRAPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ScgFailureInfoEUTRAPresent bool
	if ScgFailureInfoEUTRAPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Drx_ConfigMCGPresent bool
	if Drx_ConfigMCGPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var MeasResultReportCGI_EUTRAPresent bool
	if MeasResultReportCGI_EUTRAPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var MeasResultCellListSFTD_EUTRAPresent bool
	if MeasResultCellListSFTD_EUTRAPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Fr_InfoListMCGPresent bool
	if Fr_InfoListMCGPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var NonCriticalExtensionPresent bool
	if NonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if CandidateCellInfoListMN_EUTRAPresent {
		var tmp_os_CandidateCellInfoListMN_EUTRA []byte
		if tmp_os_CandidateCellInfoListMN_EUTRA, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode CandidateCellInfoListMN_EUTRA", err)
		}
		ie.CandidateCellInfoListMN_EUTRA = &tmp_os_CandidateCellInfoListMN_EUTRA
	}
	if CandidateCellInfoListSN_EUTRAPresent {
		var tmp_os_CandidateCellInfoListSN_EUTRA []byte
		if tmp_os_CandidateCellInfoListSN_EUTRA, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode CandidateCellInfoListSN_EUTRA", err)
		}
		ie.CandidateCellInfoListSN_EUTRA = &tmp_os_CandidateCellInfoListSN_EUTRA
	}
	if SourceConfigSCG_EUTRAPresent {
		var tmp_os_SourceConfigSCG_EUTRA []byte
		if tmp_os_SourceConfigSCG_EUTRA, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode SourceConfigSCG_EUTRA", err)
		}
		ie.SourceConfigSCG_EUTRA = &tmp_os_SourceConfigSCG_EUTRA
	}
	if ScgFailureInfoEUTRAPresent {
		ie.ScgFailureInfoEUTRA = new(CG_ConfigInfo_v1560_IEs_scgFailureInfoEUTRA)
		if err = ie.ScgFailureInfoEUTRA.Decode(r); err != nil {
			return utils.WrapError("Decode ScgFailureInfoEUTRA", err)
		}
	}
	if Drx_ConfigMCGPresent {
		ie.Drx_ConfigMCG = new(DRX_Config)
		if err = ie.Drx_ConfigMCG.Decode(r); err != nil {
			return utils.WrapError("Decode Drx_ConfigMCG", err)
		}
	}
	if MeasResultReportCGI_EUTRAPresent {
		ie.MeasResultReportCGI_EUTRA = new(CG_ConfigInfo_v1560_IEs_measResultReportCGI_EUTRA)
		if err = ie.MeasResultReportCGI_EUTRA.Decode(r); err != nil {
			return utils.WrapError("Decode MeasResultReportCGI_EUTRA", err)
		}
	}
	if MeasResultCellListSFTD_EUTRAPresent {
		ie.MeasResultCellListSFTD_EUTRA = new(MeasResultCellListSFTD_EUTRA)
		if err = ie.MeasResultCellListSFTD_EUTRA.Decode(r); err != nil {
			return utils.WrapError("Decode MeasResultCellListSFTD_EUTRA", err)
		}
	}
	if Fr_InfoListMCGPresent {
		ie.Fr_InfoListMCG = new(FR_InfoList)
		if err = ie.Fr_InfoListMCG.Decode(r); err != nil {
			return utils.WrapError("Decode Fr_InfoListMCG", err)
		}
	}
	if NonCriticalExtensionPresent {
		ie.NonCriticalExtension = new(CG_ConfigInfo_v1570_IEs)
		if err = ie.NonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode NonCriticalExtension", err)
		}
	}
	return nil
}
