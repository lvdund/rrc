package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CG_Config_IEs struct {
	Scg_CellGroupConfig        *[]byte                     `optional`
	Scg_RB_Config              *[]byte                     `optional`
	ConfigRestrictModReq       *ConfigRestrictModReqSCG    `optional`
	Drx_InfoSCG                *DRX_Info                   `optional`
	CandidateCellInfoListSN    *[]byte                     `optional`
	MeasConfigSN               *MeasConfigSN               `optional`
	SelectedBandCombination    *BandCombinationInfoSN      `optional`
	Fr_InfoListSCG             *FR_InfoList                `optional`
	CandidateServingFreqListNR *CandidateServingFreqListNR `optional`
	NonCriticalExtension       *CG_Config_v1540_IEs        `optional`
}

func (ie *CG_Config_IEs) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Scg_CellGroupConfig != nil, ie.Scg_RB_Config != nil, ie.ConfigRestrictModReq != nil, ie.Drx_InfoSCG != nil, ie.CandidateCellInfoListSN != nil, ie.MeasConfigSN != nil, ie.SelectedBandCombination != nil, ie.Fr_InfoListSCG != nil, ie.CandidateServingFreqListNR != nil, ie.NonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Scg_CellGroupConfig != nil {
		if err = w.WriteOctetString(*ie.Scg_CellGroupConfig, nil, false); err != nil {
			return utils.WrapError("Encode Scg_CellGroupConfig", err)
		}
	}
	if ie.Scg_RB_Config != nil {
		if err = w.WriteOctetString(*ie.Scg_RB_Config, nil, false); err != nil {
			return utils.WrapError("Encode Scg_RB_Config", err)
		}
	}
	if ie.ConfigRestrictModReq != nil {
		if err = ie.ConfigRestrictModReq.Encode(w); err != nil {
			return utils.WrapError("Encode ConfigRestrictModReq", err)
		}
	}
	if ie.Drx_InfoSCG != nil {
		if err = ie.Drx_InfoSCG.Encode(w); err != nil {
			return utils.WrapError("Encode Drx_InfoSCG", err)
		}
	}
	if ie.CandidateCellInfoListSN != nil {
		if err = w.WriteOctetString(*ie.CandidateCellInfoListSN, nil, false); err != nil {
			return utils.WrapError("Encode CandidateCellInfoListSN", err)
		}
	}
	if ie.MeasConfigSN != nil {
		if err = ie.MeasConfigSN.Encode(w); err != nil {
			return utils.WrapError("Encode MeasConfigSN", err)
		}
	}
	if ie.SelectedBandCombination != nil {
		if err = ie.SelectedBandCombination.Encode(w); err != nil {
			return utils.WrapError("Encode SelectedBandCombination", err)
		}
	}
	if ie.Fr_InfoListSCG != nil {
		if err = ie.Fr_InfoListSCG.Encode(w); err != nil {
			return utils.WrapError("Encode Fr_InfoListSCG", err)
		}
	}
	if ie.CandidateServingFreqListNR != nil {
		if err = ie.CandidateServingFreqListNR.Encode(w); err != nil {
			return utils.WrapError("Encode CandidateServingFreqListNR", err)
		}
	}
	if ie.NonCriticalExtension != nil {
		if err = ie.NonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode NonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *CG_Config_IEs) Decode(r *aper.AperReader) error {
	var err error
	var Scg_CellGroupConfigPresent bool
	if Scg_CellGroupConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Scg_RB_ConfigPresent bool
	if Scg_RB_ConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ConfigRestrictModReqPresent bool
	if ConfigRestrictModReqPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Drx_InfoSCGPresent bool
	if Drx_InfoSCGPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var CandidateCellInfoListSNPresent bool
	if CandidateCellInfoListSNPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var MeasConfigSNPresent bool
	if MeasConfigSNPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var SelectedBandCombinationPresent bool
	if SelectedBandCombinationPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Fr_InfoListSCGPresent bool
	if Fr_InfoListSCGPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var CandidateServingFreqListNRPresent bool
	if CandidateServingFreqListNRPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var NonCriticalExtensionPresent bool
	if NonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Scg_CellGroupConfigPresent {
		var tmp_os_Scg_CellGroupConfig []byte
		if tmp_os_Scg_CellGroupConfig, err = r.ReadOctetString(nil, false); err != nil {
			return utils.WrapError("Decode Scg_CellGroupConfig", err)
		}
		ie.Scg_CellGroupConfig = &tmp_os_Scg_CellGroupConfig
	}
	if Scg_RB_ConfigPresent {
		var tmp_os_Scg_RB_Config []byte
		if tmp_os_Scg_RB_Config, err = r.ReadOctetString(nil, false); err != nil {
			return utils.WrapError("Decode Scg_RB_Config", err)
		}
		ie.Scg_RB_Config = &tmp_os_Scg_RB_Config
	}
	if ConfigRestrictModReqPresent {
		ie.ConfigRestrictModReq = new(ConfigRestrictModReqSCG)
		if err = ie.ConfigRestrictModReq.Decode(r); err != nil {
			return utils.WrapError("Decode ConfigRestrictModReq", err)
		}
	}
	if Drx_InfoSCGPresent {
		ie.Drx_InfoSCG = new(DRX_Info)
		if err = ie.Drx_InfoSCG.Decode(r); err != nil {
			return utils.WrapError("Decode Drx_InfoSCG", err)
		}
	}
	if CandidateCellInfoListSNPresent {
		var tmp_os_CandidateCellInfoListSN []byte
		if tmp_os_CandidateCellInfoListSN, err = r.ReadOctetString(nil, false); err != nil {
			return utils.WrapError("Decode CandidateCellInfoListSN", err)
		}
		ie.CandidateCellInfoListSN = &tmp_os_CandidateCellInfoListSN
	}
	if MeasConfigSNPresent {
		ie.MeasConfigSN = new(MeasConfigSN)
		if err = ie.MeasConfigSN.Decode(r); err != nil {
			return utils.WrapError("Decode MeasConfigSN", err)
		}
	}
	if SelectedBandCombinationPresent {
		ie.SelectedBandCombination = new(BandCombinationInfoSN)
		if err = ie.SelectedBandCombination.Decode(r); err != nil {
			return utils.WrapError("Decode SelectedBandCombination", err)
		}
	}
	if Fr_InfoListSCGPresent {
		ie.Fr_InfoListSCG = new(FR_InfoList)
		if err = ie.Fr_InfoListSCG.Decode(r); err != nil {
			return utils.WrapError("Decode Fr_InfoListSCG", err)
		}
	}
	if CandidateServingFreqListNRPresent {
		ie.CandidateServingFreqListNR = new(CandidateServingFreqListNR)
		if err = ie.CandidateServingFreqListNR.Decode(r); err != nil {
			return utils.WrapError("Decode CandidateServingFreqListNR", err)
		}
	}
	if NonCriticalExtensionPresent {
		ie.NonCriticalExtension = new(CG_Config_v1540_IEs)
		if err = ie.NonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode NonCriticalExtension", err)
		}
	}
	return nil
}
