// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UEAssistanceInformation-v1610-IEs (line 2443).

var uEAssistanceInformationV1610IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "idc-Assistance-r16", Optional: true},
		{Name: "drx-Preference-r16", Optional: true},
		{Name: "maxBW-Preference-r16", Optional: true},
		{Name: "maxCC-Preference-r16", Optional: true},
		{Name: "maxMIMO-LayerPreference-r16", Optional: true},
		{Name: "minSchedulingOffsetPreference-r16", Optional: true},
		{Name: "releasePreference-r16", Optional: true},
		{Name: "sl-UE-AssistanceInformationNR-r16", Optional: true},
		{Name: "referenceTimeInfoPreference-r16", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

type UEAssistanceInformation_v1610_IEs struct {
	Idc_Assistance_r16                *IDC_Assistance_r16
	Drx_Preference_r16                *DRX_Preference_r16
	MaxBW_Preference_r16              *MaxBW_Preference_r16
	MaxCC_Preference_r16              *MaxCC_Preference_r16
	MaxMIMO_LayerPreference_r16       *MaxMIMO_LayerPreference_r16
	MinSchedulingOffsetPreference_r16 *MinSchedulingOffsetPreference_r16
	ReleasePreference_r16             *ReleasePreference_r16
	Sl_UE_AssistanceInformationNR_r16 *SL_UE_AssistanceInformationNR_r16
	ReferenceTimeInfoPreference_r16   *bool
	NonCriticalExtension              *UEAssistanceInformation_v1700_IEs
}

func (ie *UEAssistanceInformation_v1610_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uEAssistanceInformationV1610IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Idc_Assistance_r16 != nil, ie.Drx_Preference_r16 != nil, ie.MaxBW_Preference_r16 != nil, ie.MaxCC_Preference_r16 != nil, ie.MaxMIMO_LayerPreference_r16 != nil, ie.MinSchedulingOffsetPreference_r16 != nil, ie.ReleasePreference_r16 != nil, ie.Sl_UE_AssistanceInformationNR_r16 != nil, ie.ReferenceTimeInfoPreference_r16 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. idc-Assistance-r16: ref
	{
		if ie.Idc_Assistance_r16 != nil {
			if err := ie.Idc_Assistance_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. drx-Preference-r16: ref
	{
		if ie.Drx_Preference_r16 != nil {
			if err := ie.Drx_Preference_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. maxBW-Preference-r16: ref
	{
		if ie.MaxBW_Preference_r16 != nil {
			if err := ie.MaxBW_Preference_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. maxCC-Preference-r16: ref
	{
		if ie.MaxCC_Preference_r16 != nil {
			if err := ie.MaxCC_Preference_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. maxMIMO-LayerPreference-r16: ref
	{
		if ie.MaxMIMO_LayerPreference_r16 != nil {
			if err := ie.MaxMIMO_LayerPreference_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 7. minSchedulingOffsetPreference-r16: ref
	{
		if ie.MinSchedulingOffsetPreference_r16 != nil {
			if err := ie.MinSchedulingOffsetPreference_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 8. releasePreference-r16: ref
	{
		if ie.ReleasePreference_r16 != nil {
			if err := ie.ReleasePreference_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 9. sl-UE-AssistanceInformationNR-r16: ref
	{
		if ie.Sl_UE_AssistanceInformationNR_r16 != nil {
			if err := ie.Sl_UE_AssistanceInformationNR_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 10. referenceTimeInfoPreference-r16: boolean
	{
		if ie.ReferenceTimeInfoPreference_r16 != nil {
			if err := e.EncodeBoolean(*ie.ReferenceTimeInfoPreference_r16); err != nil {
				return err
			}
		}
	}

	// 11. nonCriticalExtension: ref
	{
		if ie.NonCriticalExtension != nil {
			if err := ie.NonCriticalExtension.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *UEAssistanceInformation_v1610_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uEAssistanceInformationV1610IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. idc-Assistance-r16: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Idc_Assistance_r16 = new(IDC_Assistance_r16)
			if err := ie.Idc_Assistance_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. drx-Preference-r16: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Drx_Preference_r16 = new(DRX_Preference_r16)
			if err := ie.Drx_Preference_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. maxBW-Preference-r16: ref
	{
		if seq.IsComponentPresent(2) {
			ie.MaxBW_Preference_r16 = new(MaxBW_Preference_r16)
			if err := ie.MaxBW_Preference_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. maxCC-Preference-r16: ref
	{
		if seq.IsComponentPresent(3) {
			ie.MaxCC_Preference_r16 = new(MaxCC_Preference_r16)
			if err := ie.MaxCC_Preference_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. maxMIMO-LayerPreference-r16: ref
	{
		if seq.IsComponentPresent(4) {
			ie.MaxMIMO_LayerPreference_r16 = new(MaxMIMO_LayerPreference_r16)
			if err := ie.MaxMIMO_LayerPreference_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 7. minSchedulingOffsetPreference-r16: ref
	{
		if seq.IsComponentPresent(5) {
			ie.MinSchedulingOffsetPreference_r16 = new(MinSchedulingOffsetPreference_r16)
			if err := ie.MinSchedulingOffsetPreference_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 8. releasePreference-r16: ref
	{
		if seq.IsComponentPresent(6) {
			ie.ReleasePreference_r16 = new(ReleasePreference_r16)
			if err := ie.ReleasePreference_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 9. sl-UE-AssistanceInformationNR-r16: ref
	{
		if seq.IsComponentPresent(7) {
			ie.Sl_UE_AssistanceInformationNR_r16 = new(SL_UE_AssistanceInformationNR_r16)
			if err := ie.Sl_UE_AssistanceInformationNR_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 10. referenceTimeInfoPreference-r16: boolean
	{
		if seq.IsComponentPresent(8) {
			v, err := d.DecodeBoolean()
			if err != nil {
				return err
			}
			ie.ReferenceTimeInfoPreference_r16 = &v
		}
	}

	// 11. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(9) {
			ie.NonCriticalExtension = new(UEAssistanceInformation_v1700_IEs)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
