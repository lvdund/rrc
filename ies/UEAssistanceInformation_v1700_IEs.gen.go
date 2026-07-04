// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: UEAssistanceInformation-v1700-IEs (line 2456).

var uEAssistanceInformationV1700IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "ul-GapFR2-Preference-r17", Optional: true},
		{Name: "musim-Assistance-r17", Optional: true},
		{Name: "overheatingAssistance-r17", Optional: true},
		{Name: "maxBW-PreferenceFR2-2-r17", Optional: true},
		{Name: "maxMIMO-LayerPreferenceFR2-2-r17", Optional: true},
		{Name: "minSchedulingOffsetPreferenceExt-r17", Optional: true},
		{Name: "rlm-MeasRelaxationState-r17", Optional: true},
		{Name: "bfd-MeasRelaxationState-r17", Optional: true},
		{Name: "nonSDT-DataIndication-r17", Optional: true},
		{Name: "scg-DeactivationPreference-r17", Optional: true},
		{Name: "uplinkData-r17", Optional: true},
		{Name: "rrm-MeasRelaxationFulfilment-r17", Optional: true},
		{Name: "propagationDelayDifference-r17", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

var uEAssistanceInformationV1700IEsBfdMeasRelaxationStateR17Constraints = per.SizeRange(1, common.MaxNrofServingCells)

const (
	UEAssistanceInformation_v1700_IEs_Scg_DeactivationPreference_r17_Scg_DeactivationPreferred = 0
	UEAssistanceInformation_v1700_IEs_Scg_DeactivationPreference_r17_NoPreference              = 1
)

var uEAssistanceInformationV1700IEsScgDeactivationPreferenceR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UEAssistanceInformation_v1700_IEs_Scg_DeactivationPreference_r17_Scg_DeactivationPreferred, UEAssistanceInformation_v1700_IEs_Scg_DeactivationPreference_r17_NoPreference},
}

const (
	UEAssistanceInformation_v1700_IEs_UplinkData_r17_True = 0
)

var uEAssistanceInformationV1700IEsUplinkDataR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UEAssistanceInformation_v1700_IEs_UplinkData_r17_True},
}

var uEAssistanceInformationV1700IEsNonSDTDataIndicationR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "resumeCause-r17", Optional: true},
	},
}

type UEAssistanceInformation_v1700_IEs struct {
	Ul_GapFR2_Preference_r17             *UL_GapFR2_Preference_r17
	Musim_Assistance_r17                 *MUSIM_Assistance_r17
	OverheatingAssistance_r17            *OverheatingAssistance_r17
	MaxBW_PreferenceFR2_2_r17            *MaxBW_PreferenceFR2_2_r17
	MaxMIMO_LayerPreferenceFR2_2_r17     *MaxMIMO_LayerPreferenceFR2_2_r17
	MinSchedulingOffsetPreferenceExt_r17 *MinSchedulingOffsetPreferenceExt_r17
	Rlm_MeasRelaxationState_r17          *bool
	Bfd_MeasRelaxationState_r17          *per.BitString
	NonSDT_DataIndication_r17            *struct{ ResumeCause_r17 *ResumeCause }
	Scg_DeactivationPreference_r17       *int64
	UplinkData_r17                       *int64
	Rrm_MeasRelaxationFulfilment_r17     *bool
	PropagationDelayDifference_r17       *PropagationDelayDifference_r17
	NonCriticalExtension                 *UEAssistanceInformation_v1800_IEs
}

func (ie *UEAssistanceInformation_v1700_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uEAssistanceInformationV1700IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ul_GapFR2_Preference_r17 != nil, ie.Musim_Assistance_r17 != nil, ie.OverheatingAssistance_r17 != nil, ie.MaxBW_PreferenceFR2_2_r17 != nil, ie.MaxMIMO_LayerPreferenceFR2_2_r17 != nil, ie.MinSchedulingOffsetPreferenceExt_r17 != nil, ie.Rlm_MeasRelaxationState_r17 != nil, ie.Bfd_MeasRelaxationState_r17 != nil, ie.NonSDT_DataIndication_r17 != nil, ie.Scg_DeactivationPreference_r17 != nil, ie.UplinkData_r17 != nil, ie.Rrm_MeasRelaxationFulfilment_r17 != nil, ie.PropagationDelayDifference_r17 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. ul-GapFR2-Preference-r17: ref
	{
		if ie.Ul_GapFR2_Preference_r17 != nil {
			if err := ie.Ul_GapFR2_Preference_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. musim-Assistance-r17: ref
	{
		if ie.Musim_Assistance_r17 != nil {
			if err := ie.Musim_Assistance_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. overheatingAssistance-r17: ref
	{
		if ie.OverheatingAssistance_r17 != nil {
			if err := ie.OverheatingAssistance_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. maxBW-PreferenceFR2-2-r17: ref
	{
		if ie.MaxBW_PreferenceFR2_2_r17 != nil {
			if err := ie.MaxBW_PreferenceFR2_2_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. maxMIMO-LayerPreferenceFR2-2-r17: ref
	{
		if ie.MaxMIMO_LayerPreferenceFR2_2_r17 != nil {
			if err := ie.MaxMIMO_LayerPreferenceFR2_2_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 7. minSchedulingOffsetPreferenceExt-r17: ref
	{
		if ie.MinSchedulingOffsetPreferenceExt_r17 != nil {
			if err := ie.MinSchedulingOffsetPreferenceExt_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 8. rlm-MeasRelaxationState-r17: boolean
	{
		if ie.Rlm_MeasRelaxationState_r17 != nil {
			if err := e.EncodeBoolean(*ie.Rlm_MeasRelaxationState_r17); err != nil {
				return err
			}
		}
	}

	// 9. bfd-MeasRelaxationState-r17: bit-string
	{
		if ie.Bfd_MeasRelaxationState_r17 != nil {
			if err := e.EncodeBitString(*ie.Bfd_MeasRelaxationState_r17, uEAssistanceInformationV1700IEsBfdMeasRelaxationStateR17Constraints); err != nil {
				return err
			}
		}
	}

	// 10. nonSDT-DataIndication-r17: sequence
	{
		if ie.NonSDT_DataIndication_r17 != nil {
			c := ie.NonSDT_DataIndication_r17
			uEAssistanceInformationV1700IEsNonSDTDataIndicationR17Seq := e.NewSequenceEncoder(uEAssistanceInformationV1700IEsNonSDTDataIndicationR17Constraints)
			if err := uEAssistanceInformationV1700IEsNonSDTDataIndicationR17Seq.EncodePreamble([]bool{c.ResumeCause_r17 != nil}); err != nil {
				return err
			}
			if c.ResumeCause_r17 != nil {
				if err := c.ResumeCause_r17.Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 11. scg-DeactivationPreference-r17: enumerated
	{
		if ie.Scg_DeactivationPreference_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Scg_DeactivationPreference_r17, uEAssistanceInformationV1700IEsScgDeactivationPreferenceR17Constraints); err != nil {
				return err
			}
		}
	}

	// 12. uplinkData-r17: enumerated
	{
		if ie.UplinkData_r17 != nil {
			if err := e.EncodeEnumerated(*ie.UplinkData_r17, uEAssistanceInformationV1700IEsUplinkDataR17Constraints); err != nil {
				return err
			}
		}
	}

	// 13. rrm-MeasRelaxationFulfilment-r17: boolean
	{
		if ie.Rrm_MeasRelaxationFulfilment_r17 != nil {
			if err := e.EncodeBoolean(*ie.Rrm_MeasRelaxationFulfilment_r17); err != nil {
				return err
			}
		}
	}

	// 14. propagationDelayDifference-r17: ref
	{
		if ie.PropagationDelayDifference_r17 != nil {
			if err := ie.PropagationDelayDifference_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 15. nonCriticalExtension: ref
	{
		if ie.NonCriticalExtension != nil {
			if err := ie.NonCriticalExtension.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *UEAssistanceInformation_v1700_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uEAssistanceInformationV1700IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. ul-GapFR2-Preference-r17: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Ul_GapFR2_Preference_r17 = new(UL_GapFR2_Preference_r17)
			if err := ie.Ul_GapFR2_Preference_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. musim-Assistance-r17: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Musim_Assistance_r17 = new(MUSIM_Assistance_r17)
			if err := ie.Musim_Assistance_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. overheatingAssistance-r17: ref
	{
		if seq.IsComponentPresent(2) {
			ie.OverheatingAssistance_r17 = new(OverheatingAssistance_r17)
			if err := ie.OverheatingAssistance_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. maxBW-PreferenceFR2-2-r17: ref
	{
		if seq.IsComponentPresent(3) {
			ie.MaxBW_PreferenceFR2_2_r17 = new(MaxBW_PreferenceFR2_2_r17)
			if err := ie.MaxBW_PreferenceFR2_2_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. maxMIMO-LayerPreferenceFR2-2-r17: ref
	{
		if seq.IsComponentPresent(4) {
			ie.MaxMIMO_LayerPreferenceFR2_2_r17 = new(MaxMIMO_LayerPreferenceFR2_2_r17)
			if err := ie.MaxMIMO_LayerPreferenceFR2_2_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 7. minSchedulingOffsetPreferenceExt-r17: ref
	{
		if seq.IsComponentPresent(5) {
			ie.MinSchedulingOffsetPreferenceExt_r17 = new(MinSchedulingOffsetPreferenceExt_r17)
			if err := ie.MinSchedulingOffsetPreferenceExt_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 8. rlm-MeasRelaxationState-r17: boolean
	{
		if seq.IsComponentPresent(6) {
			v, err := d.DecodeBoolean()
			if err != nil {
				return err
			}
			ie.Rlm_MeasRelaxationState_r17 = &v
		}
	}

	// 9. bfd-MeasRelaxationState-r17: bit-string
	{
		if seq.IsComponentPresent(7) {
			v, err := d.DecodeBitString(uEAssistanceInformationV1700IEsBfdMeasRelaxationStateR17Constraints)
			if err != nil {
				return err
			}
			ie.Bfd_MeasRelaxationState_r17 = &v
		}
	}

	// 10. nonSDT-DataIndication-r17: sequence
	{
		if seq.IsComponentPresent(8) {
			ie.NonSDT_DataIndication_r17 = &struct{ ResumeCause_r17 *ResumeCause }{}
			c := ie.NonSDT_DataIndication_r17
			uEAssistanceInformationV1700IEsNonSDTDataIndicationR17Seq := d.NewSequenceDecoder(uEAssistanceInformationV1700IEsNonSDTDataIndicationR17Constraints)
			if err := uEAssistanceInformationV1700IEsNonSDTDataIndicationR17Seq.DecodePreamble(); err != nil {
				return err
			}
			if uEAssistanceInformationV1700IEsNonSDTDataIndicationR17Seq.IsComponentPresent(0) {
				c.ResumeCause_r17 = new(ResumeCause)
				if err := (*c.ResumeCause_r17).Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 11. scg-DeactivationPreference-r17: enumerated
	{
		if seq.IsComponentPresent(9) {
			idx, err := d.DecodeEnumerated(uEAssistanceInformationV1700IEsScgDeactivationPreferenceR17Constraints)
			if err != nil {
				return err
			}
			ie.Scg_DeactivationPreference_r17 = &idx
		}
	}

	// 12. uplinkData-r17: enumerated
	{
		if seq.IsComponentPresent(10) {
			idx, err := d.DecodeEnumerated(uEAssistanceInformationV1700IEsUplinkDataR17Constraints)
			if err != nil {
				return err
			}
			ie.UplinkData_r17 = &idx
		}
	}

	// 13. rrm-MeasRelaxationFulfilment-r17: boolean
	{
		if seq.IsComponentPresent(11) {
			v, err := d.DecodeBoolean()
			if err != nil {
				return err
			}
			ie.Rrm_MeasRelaxationFulfilment_r17 = &v
		}
	}

	// 14. propagationDelayDifference-r17: ref
	{
		if seq.IsComponentPresent(12) {
			ie.PropagationDelayDifference_r17 = new(PropagationDelayDifference_r17)
			if err := ie.PropagationDelayDifference_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 15. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(13) {
			ie.NonCriticalExtension = new(UEAssistanceInformation_v1800_IEs)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
