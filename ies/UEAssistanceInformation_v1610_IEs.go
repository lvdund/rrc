package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UEAssistanceInformation_v1610_IEs struct {
	Idc_Assistance_r16                *IDC_Assistance_r16                `optional`
	Drx_Preference_r16                *DRX_Preference_r16                `optional`
	MaxBW_Preference_r16              *MaxBW_Preference_r16              `optional`
	MaxCC_Preference_r16              *MaxCC_Preference_r16              `optional`
	MaxMIMO_LayerPreference_r16       *MaxMIMO_LayerPreference_r16       `optional`
	MinSchedulingOffsetPreference_r16 *MinSchedulingOffsetPreference_r16 `optional`
	ReleasePreference_r16             *ReleasePreference_r16             `optional`
	Sl_UE_AssistanceInformationNR_r16 *SL_UE_AssistanceInformationNR_r16 `optional`
	ReferenceTimeInfoPreference_r16   *bool                              `optional`
	NonCriticalExtension              *UEAssistanceInformation_v1700_IEs `optional`
}

func (ie *UEAssistanceInformation_v1610_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Idc_Assistance_r16 != nil, ie.Drx_Preference_r16 != nil, ie.MaxBW_Preference_r16 != nil, ie.MaxCC_Preference_r16 != nil, ie.MaxMIMO_LayerPreference_r16 != nil, ie.MinSchedulingOffsetPreference_r16 != nil, ie.ReleasePreference_r16 != nil, ie.Sl_UE_AssistanceInformationNR_r16 != nil, ie.ReferenceTimeInfoPreference_r16 != nil, ie.NonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Idc_Assistance_r16 != nil {
		if err = ie.Idc_Assistance_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Idc_Assistance_r16", err)
		}
	}
	if ie.Drx_Preference_r16 != nil {
		if err = ie.Drx_Preference_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Drx_Preference_r16", err)
		}
	}
	if ie.MaxBW_Preference_r16 != nil {
		if err = ie.MaxBW_Preference_r16.Encode(w); err != nil {
			return utils.WrapError("Encode MaxBW_Preference_r16", err)
		}
	}
	if ie.MaxCC_Preference_r16 != nil {
		if err = ie.MaxCC_Preference_r16.Encode(w); err != nil {
			return utils.WrapError("Encode MaxCC_Preference_r16", err)
		}
	}
	if ie.MaxMIMO_LayerPreference_r16 != nil {
		if err = ie.MaxMIMO_LayerPreference_r16.Encode(w); err != nil {
			return utils.WrapError("Encode MaxMIMO_LayerPreference_r16", err)
		}
	}
	if ie.MinSchedulingOffsetPreference_r16 != nil {
		if err = ie.MinSchedulingOffsetPreference_r16.Encode(w); err != nil {
			return utils.WrapError("Encode MinSchedulingOffsetPreference_r16", err)
		}
	}
	if ie.ReleasePreference_r16 != nil {
		if err = ie.ReleasePreference_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ReleasePreference_r16", err)
		}
	}
	if ie.Sl_UE_AssistanceInformationNR_r16 != nil {
		if err = ie.Sl_UE_AssistanceInformationNR_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_UE_AssistanceInformationNR_r16", err)
		}
	}
	if ie.ReferenceTimeInfoPreference_r16 != nil {
		if err = w.WriteBoolean(*ie.ReferenceTimeInfoPreference_r16); err != nil {
			return utils.WrapError("Encode ReferenceTimeInfoPreference_r16", err)
		}
	}
	if ie.NonCriticalExtension != nil {
		if err = ie.NonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode NonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *UEAssistanceInformation_v1610_IEs) Decode(r *uper.UperReader) error {
	var err error
	var Idc_Assistance_r16Present bool
	if Idc_Assistance_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Drx_Preference_r16Present bool
	if Drx_Preference_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MaxBW_Preference_r16Present bool
	if MaxBW_Preference_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MaxCC_Preference_r16Present bool
	if MaxCC_Preference_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MaxMIMO_LayerPreference_r16Present bool
	if MaxMIMO_LayerPreference_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MinSchedulingOffsetPreference_r16Present bool
	if MinSchedulingOffsetPreference_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ReleasePreference_r16Present bool
	if ReleasePreference_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_UE_AssistanceInformationNR_r16Present bool
	if Sl_UE_AssistanceInformationNR_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ReferenceTimeInfoPreference_r16Present bool
	if ReferenceTimeInfoPreference_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var NonCriticalExtensionPresent bool
	if NonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Idc_Assistance_r16Present {
		ie.Idc_Assistance_r16 = new(IDC_Assistance_r16)
		if err = ie.Idc_Assistance_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Idc_Assistance_r16", err)
		}
	}
	if Drx_Preference_r16Present {
		ie.Drx_Preference_r16 = new(DRX_Preference_r16)
		if err = ie.Drx_Preference_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Drx_Preference_r16", err)
		}
	}
	if MaxBW_Preference_r16Present {
		ie.MaxBW_Preference_r16 = new(MaxBW_Preference_r16)
		if err = ie.MaxBW_Preference_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MaxBW_Preference_r16", err)
		}
	}
	if MaxCC_Preference_r16Present {
		ie.MaxCC_Preference_r16 = new(MaxCC_Preference_r16)
		if err = ie.MaxCC_Preference_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MaxCC_Preference_r16", err)
		}
	}
	if MaxMIMO_LayerPreference_r16Present {
		ie.MaxMIMO_LayerPreference_r16 = new(MaxMIMO_LayerPreference_r16)
		if err = ie.MaxMIMO_LayerPreference_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MaxMIMO_LayerPreference_r16", err)
		}
	}
	if MinSchedulingOffsetPreference_r16Present {
		ie.MinSchedulingOffsetPreference_r16 = new(MinSchedulingOffsetPreference_r16)
		if err = ie.MinSchedulingOffsetPreference_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MinSchedulingOffsetPreference_r16", err)
		}
	}
	if ReleasePreference_r16Present {
		ie.ReleasePreference_r16 = new(ReleasePreference_r16)
		if err = ie.ReleasePreference_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ReleasePreference_r16", err)
		}
	}
	if Sl_UE_AssistanceInformationNR_r16Present {
		ie.Sl_UE_AssistanceInformationNR_r16 = new(SL_UE_AssistanceInformationNR_r16)
		if err = ie.Sl_UE_AssistanceInformationNR_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_UE_AssistanceInformationNR_r16", err)
		}
	}
	if ReferenceTimeInfoPreference_r16Present {
		var tmp_bool_ReferenceTimeInfoPreference_r16 bool
		if tmp_bool_ReferenceTimeInfoPreference_r16, err = r.ReadBoolean(); err != nil {
			return utils.WrapError("Decode ReferenceTimeInfoPreference_r16", err)
		}
		ie.ReferenceTimeInfoPreference_r16 = &tmp_bool_ReferenceTimeInfoPreference_r16
	}
	if NonCriticalExtensionPresent {
		ie.NonCriticalExtension = new(UEAssistanceInformation_v1700_IEs)
		if err = ie.NonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode NonCriticalExtension", err)
		}
	}
	return nil
}
