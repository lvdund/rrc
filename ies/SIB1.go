package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SIB1 struct {
	CellSelectionInfo        *SIB1_cellSelectionInfo     `lb:1,ub:8,optional`
	CellAccessRelatedInfo    CellAccessRelatedInfo       `madatory`
	ConnEstFailureControl    *ConnEstFailureControl      `optional`
	Si_SchedulingInfo        *SI_SchedulingInfo          `optional`
	ServingCellConfigCommon  *ServingCellConfigCommonSIB `optional`
	Ims_EmergencySupport     *SIB1_ims_EmergencySupport  `optional`
	ECallOverIMS_Support     *SIB1_eCallOverIMS_Support  `optional`
	Ue_TimersAndConstants    *UE_TimersAndConstants      `optional`
	Uac_BarringInfo          *SIB1_uac_BarringInfo       `lb:2,ub:maxPLMN,optional`
	UseFullResumeID          *SIB1_useFullResumeID       `optional`
	LateNonCriticalExtension *[]byte                     `optional`
	NonCriticalExtension     *SIB1_v1610_IEs             `optional`
}

func (ie *SIB1) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.CellSelectionInfo != nil, ie.ConnEstFailureControl != nil, ie.Si_SchedulingInfo != nil, ie.ServingCellConfigCommon != nil, ie.Ims_EmergencySupport != nil, ie.ECallOverIMS_Support != nil, ie.Ue_TimersAndConstants != nil, ie.Uac_BarringInfo != nil, ie.UseFullResumeID != nil, ie.LateNonCriticalExtension != nil, ie.NonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.CellSelectionInfo != nil {
		if err = ie.CellSelectionInfo.Encode(w); err != nil {
			return utils.WrapError("Encode CellSelectionInfo", err)
		}
	}
	if err = ie.CellAccessRelatedInfo.Encode(w); err != nil {
		return utils.WrapError("Encode CellAccessRelatedInfo", err)
	}
	if ie.ConnEstFailureControl != nil {
		if err = ie.ConnEstFailureControl.Encode(w); err != nil {
			return utils.WrapError("Encode ConnEstFailureControl", err)
		}
	}
	if ie.Si_SchedulingInfo != nil {
		if err = ie.Si_SchedulingInfo.Encode(w); err != nil {
			return utils.WrapError("Encode Si_SchedulingInfo", err)
		}
	}
	if ie.ServingCellConfigCommon != nil {
		if err = ie.ServingCellConfigCommon.Encode(w); err != nil {
			return utils.WrapError("Encode ServingCellConfigCommon", err)
		}
	}
	if ie.Ims_EmergencySupport != nil {
		if err = ie.Ims_EmergencySupport.Encode(w); err != nil {
			return utils.WrapError("Encode Ims_EmergencySupport", err)
		}
	}
	if ie.ECallOverIMS_Support != nil {
		if err = ie.ECallOverIMS_Support.Encode(w); err != nil {
			return utils.WrapError("Encode ECallOverIMS_Support", err)
		}
	}
	if ie.Ue_TimersAndConstants != nil {
		if err = ie.Ue_TimersAndConstants.Encode(w); err != nil {
			return utils.WrapError("Encode Ue_TimersAndConstants", err)
		}
	}
	if ie.Uac_BarringInfo != nil {
		if err = ie.Uac_BarringInfo.Encode(w); err != nil {
			return utils.WrapError("Encode Uac_BarringInfo", err)
		}
	}
	if ie.UseFullResumeID != nil {
		if err = ie.UseFullResumeID.Encode(w); err != nil {
			return utils.WrapError("Encode UseFullResumeID", err)
		}
	}
	if ie.LateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.LateNonCriticalExtension, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode LateNonCriticalExtension", err)
		}
	}
	if ie.NonCriticalExtension != nil {
		if err = ie.NonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode NonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *SIB1) Decode(r *uper.UperReader) error {
	var err error
	var CellSelectionInfoPresent bool
	if CellSelectionInfoPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ConnEstFailureControlPresent bool
	if ConnEstFailureControlPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Si_SchedulingInfoPresent bool
	if Si_SchedulingInfoPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ServingCellConfigCommonPresent bool
	if ServingCellConfigCommonPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Ims_EmergencySupportPresent bool
	if Ims_EmergencySupportPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ECallOverIMS_SupportPresent bool
	if ECallOverIMS_SupportPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Ue_TimersAndConstantsPresent bool
	if Ue_TimersAndConstantsPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Uac_BarringInfoPresent bool
	if Uac_BarringInfoPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var UseFullResumeIDPresent bool
	if UseFullResumeIDPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var LateNonCriticalExtensionPresent bool
	if LateNonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var NonCriticalExtensionPresent bool
	if NonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if CellSelectionInfoPresent {
		ie.CellSelectionInfo = new(SIB1_cellSelectionInfo)
		if err = ie.CellSelectionInfo.Decode(r); err != nil {
			return utils.WrapError("Decode CellSelectionInfo", err)
		}
	}
	if err = ie.CellAccessRelatedInfo.Decode(r); err != nil {
		return utils.WrapError("Decode CellAccessRelatedInfo", err)
	}
	if ConnEstFailureControlPresent {
		ie.ConnEstFailureControl = new(ConnEstFailureControl)
		if err = ie.ConnEstFailureControl.Decode(r); err != nil {
			return utils.WrapError("Decode ConnEstFailureControl", err)
		}
	}
	if Si_SchedulingInfoPresent {
		ie.Si_SchedulingInfo = new(SI_SchedulingInfo)
		if err = ie.Si_SchedulingInfo.Decode(r); err != nil {
			return utils.WrapError("Decode Si_SchedulingInfo", err)
		}
	}
	if ServingCellConfigCommonPresent {
		ie.ServingCellConfigCommon = new(ServingCellConfigCommonSIB)
		if err = ie.ServingCellConfigCommon.Decode(r); err != nil {
			return utils.WrapError("Decode ServingCellConfigCommon", err)
		}
	}
	if Ims_EmergencySupportPresent {
		ie.Ims_EmergencySupport = new(SIB1_ims_EmergencySupport)
		if err = ie.Ims_EmergencySupport.Decode(r); err != nil {
			return utils.WrapError("Decode Ims_EmergencySupport", err)
		}
	}
	if ECallOverIMS_SupportPresent {
		ie.ECallOverIMS_Support = new(SIB1_eCallOverIMS_Support)
		if err = ie.ECallOverIMS_Support.Decode(r); err != nil {
			return utils.WrapError("Decode ECallOverIMS_Support", err)
		}
	}
	if Ue_TimersAndConstantsPresent {
		ie.Ue_TimersAndConstants = new(UE_TimersAndConstants)
		if err = ie.Ue_TimersAndConstants.Decode(r); err != nil {
			return utils.WrapError("Decode Ue_TimersAndConstants", err)
		}
	}
	if Uac_BarringInfoPresent {
		ie.Uac_BarringInfo = new(SIB1_uac_BarringInfo)
		if err = ie.Uac_BarringInfo.Decode(r); err != nil {
			return utils.WrapError("Decode Uac_BarringInfo", err)
		}
	}
	if UseFullResumeIDPresent {
		ie.UseFullResumeID = new(SIB1_useFullResumeID)
		if err = ie.UseFullResumeID.Decode(r); err != nil {
			return utils.WrapError("Decode UseFullResumeID", err)
		}
	}
	if LateNonCriticalExtensionPresent {
		var tmp_os_LateNonCriticalExtension []byte
		if tmp_os_LateNonCriticalExtension, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode LateNonCriticalExtension", err)
		}
		ie.LateNonCriticalExtension = &tmp_os_LateNonCriticalExtension
	}
	if NonCriticalExtensionPresent {
		ie.NonCriticalExtension = new(SIB1_v1610_IEs)
		if err = ie.NonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode NonCriticalExtension", err)
		}
	}
	return nil
}
