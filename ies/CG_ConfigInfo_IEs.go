package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CG_ConfigInfo_IEs struct {
	Ue_CapabilityInfo         *[]byte                           `optional`
	CandidateCellInfoListMN   *MeasResultList2NR                `optional`
	CandidateCellInfoListSN   *[]byte                           `optional`
	MeasResultCellListSFTD_NR *MeasResultCellListSFTD_NR        `optional`
	ScgFailureInfo            *CG_ConfigInfo_IEs_scgFailureInfo `optional`
	ConfigRestrictInfo        *ConfigRestrictInfoSCG            `optional`
	Drx_InfoMCG               *DRX_Info                         `optional`
	MeasConfigMN              *MeasConfigMN                     `optional`
	SourceConfigSCG           *[]byte                           `optional`
	Scg_RB_Config             *[]byte                           `optional`
	Mcg_RB_Config             *[]byte                           `optional`
	Mrdc_AssistanceInfo       *MRDC_AssistanceInfo              `optional`
	NonCriticalExtension      *CG_ConfigInfo_v1540_IEs          `optional`
}

func (ie *CG_ConfigInfo_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Ue_CapabilityInfo != nil, ie.CandidateCellInfoListMN != nil, ie.CandidateCellInfoListSN != nil, ie.MeasResultCellListSFTD_NR != nil, ie.ScgFailureInfo != nil, ie.ConfigRestrictInfo != nil, ie.Drx_InfoMCG != nil, ie.MeasConfigMN != nil, ie.SourceConfigSCG != nil, ie.Scg_RB_Config != nil, ie.Mcg_RB_Config != nil, ie.Mrdc_AssistanceInfo != nil, ie.NonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Ue_CapabilityInfo != nil {
		if err = w.WriteOctetString(*ie.Ue_CapabilityInfo, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode Ue_CapabilityInfo", err)
		}
	}
	if ie.CandidateCellInfoListMN != nil {
		if err = ie.CandidateCellInfoListMN.Encode(w); err != nil {
			return utils.WrapError("Encode CandidateCellInfoListMN", err)
		}
	}
	if ie.CandidateCellInfoListSN != nil {
		if err = w.WriteOctetString(*ie.CandidateCellInfoListSN, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode CandidateCellInfoListSN", err)
		}
	}
	if ie.MeasResultCellListSFTD_NR != nil {
		if err = ie.MeasResultCellListSFTD_NR.Encode(w); err != nil {
			return utils.WrapError("Encode MeasResultCellListSFTD_NR", err)
		}
	}
	if ie.ScgFailureInfo != nil {
		if err = ie.ScgFailureInfo.Encode(w); err != nil {
			return utils.WrapError("Encode ScgFailureInfo", err)
		}
	}
	if ie.ConfigRestrictInfo != nil {
		if err = ie.ConfigRestrictInfo.Encode(w); err != nil {
			return utils.WrapError("Encode ConfigRestrictInfo", err)
		}
	}
	if ie.Drx_InfoMCG != nil {
		if err = ie.Drx_InfoMCG.Encode(w); err != nil {
			return utils.WrapError("Encode Drx_InfoMCG", err)
		}
	}
	if ie.MeasConfigMN != nil {
		if err = ie.MeasConfigMN.Encode(w); err != nil {
			return utils.WrapError("Encode MeasConfigMN", err)
		}
	}
	if ie.SourceConfigSCG != nil {
		if err = w.WriteOctetString(*ie.SourceConfigSCG, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode SourceConfigSCG", err)
		}
	}
	if ie.Scg_RB_Config != nil {
		if err = w.WriteOctetString(*ie.Scg_RB_Config, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode Scg_RB_Config", err)
		}
	}
	if ie.Mcg_RB_Config != nil {
		if err = w.WriteOctetString(*ie.Mcg_RB_Config, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode Mcg_RB_Config", err)
		}
	}
	if ie.Mrdc_AssistanceInfo != nil {
		if err = ie.Mrdc_AssistanceInfo.Encode(w); err != nil {
			return utils.WrapError("Encode Mrdc_AssistanceInfo", err)
		}
	}
	if ie.NonCriticalExtension != nil {
		if err = ie.NonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode NonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *CG_ConfigInfo_IEs) Decode(r *uper.UperReader) error {
	var err error
	var Ue_CapabilityInfoPresent bool
	if Ue_CapabilityInfoPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var CandidateCellInfoListMNPresent bool
	if CandidateCellInfoListMNPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var CandidateCellInfoListSNPresent bool
	if CandidateCellInfoListSNPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var MeasResultCellListSFTD_NRPresent bool
	if MeasResultCellListSFTD_NRPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ScgFailureInfoPresent bool
	if ScgFailureInfoPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ConfigRestrictInfoPresent bool
	if ConfigRestrictInfoPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Drx_InfoMCGPresent bool
	if Drx_InfoMCGPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var MeasConfigMNPresent bool
	if MeasConfigMNPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var SourceConfigSCGPresent bool
	if SourceConfigSCGPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Scg_RB_ConfigPresent bool
	if Scg_RB_ConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Mcg_RB_ConfigPresent bool
	if Mcg_RB_ConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Mrdc_AssistanceInfoPresent bool
	if Mrdc_AssistanceInfoPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var NonCriticalExtensionPresent bool
	if NonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Ue_CapabilityInfoPresent {
		var tmp_os_Ue_CapabilityInfo []byte
		if tmp_os_Ue_CapabilityInfo, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode Ue_CapabilityInfo", err)
		}
		ie.Ue_CapabilityInfo = &tmp_os_Ue_CapabilityInfo
	}
	if CandidateCellInfoListMNPresent {
		ie.CandidateCellInfoListMN = new(MeasResultList2NR)
		if err = ie.CandidateCellInfoListMN.Decode(r); err != nil {
			return utils.WrapError("Decode CandidateCellInfoListMN", err)
		}
	}
	if CandidateCellInfoListSNPresent {
		var tmp_os_CandidateCellInfoListSN []byte
		if tmp_os_CandidateCellInfoListSN, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode CandidateCellInfoListSN", err)
		}
		ie.CandidateCellInfoListSN = &tmp_os_CandidateCellInfoListSN
	}
	if MeasResultCellListSFTD_NRPresent {
		ie.MeasResultCellListSFTD_NR = new(MeasResultCellListSFTD_NR)
		if err = ie.MeasResultCellListSFTD_NR.Decode(r); err != nil {
			return utils.WrapError("Decode MeasResultCellListSFTD_NR", err)
		}
	}
	if ScgFailureInfoPresent {
		ie.ScgFailureInfo = new(CG_ConfigInfo_IEs_scgFailureInfo)
		if err = ie.ScgFailureInfo.Decode(r); err != nil {
			return utils.WrapError("Decode ScgFailureInfo", err)
		}
	}
	if ConfigRestrictInfoPresent {
		ie.ConfigRestrictInfo = new(ConfigRestrictInfoSCG)
		if err = ie.ConfigRestrictInfo.Decode(r); err != nil {
			return utils.WrapError("Decode ConfigRestrictInfo", err)
		}
	}
	if Drx_InfoMCGPresent {
		ie.Drx_InfoMCG = new(DRX_Info)
		if err = ie.Drx_InfoMCG.Decode(r); err != nil {
			return utils.WrapError("Decode Drx_InfoMCG", err)
		}
	}
	if MeasConfigMNPresent {
		ie.MeasConfigMN = new(MeasConfigMN)
		if err = ie.MeasConfigMN.Decode(r); err != nil {
			return utils.WrapError("Decode MeasConfigMN", err)
		}
	}
	if SourceConfigSCGPresent {
		var tmp_os_SourceConfigSCG []byte
		if tmp_os_SourceConfigSCG, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode SourceConfigSCG", err)
		}
		ie.SourceConfigSCG = &tmp_os_SourceConfigSCG
	}
	if Scg_RB_ConfigPresent {
		var tmp_os_Scg_RB_Config []byte
		if tmp_os_Scg_RB_Config, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode Scg_RB_Config", err)
		}
		ie.Scg_RB_Config = &tmp_os_Scg_RB_Config
	}
	if Mcg_RB_ConfigPresent {
		var tmp_os_Mcg_RB_Config []byte
		if tmp_os_Mcg_RB_Config, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode Mcg_RB_Config", err)
		}
		ie.Mcg_RB_Config = &tmp_os_Mcg_RB_Config
	}
	if Mrdc_AssistanceInfoPresent {
		ie.Mrdc_AssistanceInfo = new(MRDC_AssistanceInfo)
		if err = ie.Mrdc_AssistanceInfo.Decode(r); err != nil {
			return utils.WrapError("Decode Mrdc_AssistanceInfo", err)
		}
	}
	if NonCriticalExtensionPresent {
		ie.NonCriticalExtension = new(CG_ConfigInfo_v1540_IEs)
		if err = ie.NonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode NonCriticalExtension", err)
		}
	}
	return nil
}
