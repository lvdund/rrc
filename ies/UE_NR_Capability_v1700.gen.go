// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UE-NR-Capability-v1700 (line 25792).

var uENRCapabilityV1700Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "inactiveStatePO-Determination-r17", Optional: true},
		{Name: "highSpeedParameters-v1700", Optional: true},
		{Name: "powSav-Parameters-v1700", Optional: true},
		{Name: "mac-Parameters-v1700", Optional: true},
		{Name: "ims-Parameters-v1700", Optional: true},
		{Name: "measAndMobParameters-v1700"},
		{Name: "appLayerMeasParameters-r17", Optional: true},
		{Name: "redCapParameters-r17", Optional: true},
		{Name: "ra-SDT-r17", Optional: true},
		{Name: "srb-SDT-r17", Optional: true},
		{Name: "gNB-SideRTT-BasedPDC-r17", Optional: true},
		{Name: "bh-RLF-DetectionRecovery-Indication-r17", Optional: true},
		{Name: "nrdc-Parameters-v1700", Optional: true},
		{Name: "bap-Parameters-v1700", Optional: true},
		{Name: "musim-GapPreference-r17", Optional: true},
		{Name: "musimLeaveConnected-r17", Optional: true},
		{Name: "mbs-Parameters-r17"},
		{Name: "nonTerrestrialNetwork-r17", Optional: true},
		{Name: "ntn-ScenarioSupport-r17", Optional: true},
		{Name: "sliceInfoforCellReselection-r17", Optional: true},
		{Name: "ue-RadioPagingInfo-r17", Optional: true},
		{Name: "ul-GapFR2-Pattern-r17", Optional: true},
		{Name: "ntn-Parameters-r17", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

const (
	UE_NR_Capability_v1700_InactiveStatePO_Determination_r17_Supported = 0
)

var uENRCapabilityV1700InactiveStatePODeterminationR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_NR_Capability_v1700_InactiveStatePO_Determination_r17_Supported},
}

const (
	UE_NR_Capability_v1700_Ra_SDT_r17_Supported = 0
)

var uENRCapabilityV1700RaSDTR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_NR_Capability_v1700_Ra_SDT_r17_Supported},
}

const (
	UE_NR_Capability_v1700_Srb_SDT_r17_Supported = 0
)

var uENRCapabilityV1700SrbSDTR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_NR_Capability_v1700_Srb_SDT_r17_Supported},
}

const (
	UE_NR_Capability_v1700_GNB_SideRTT_BasedPDC_r17_Supported = 0
)

var uENRCapabilityV1700GNBSideRTTBasedPDCR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_NR_Capability_v1700_GNB_SideRTT_BasedPDC_r17_Supported},
}

const (
	UE_NR_Capability_v1700_Bh_RLF_DetectionRecovery_Indication_r17_Supported = 0
)

var uENRCapabilityV1700BhRLFDetectionRecoveryIndicationR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_NR_Capability_v1700_Bh_RLF_DetectionRecovery_Indication_r17_Supported},
}

const (
	UE_NR_Capability_v1700_Musim_GapPreference_r17_Supported = 0
)

var uENRCapabilityV1700MusimGapPreferenceR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_NR_Capability_v1700_Musim_GapPreference_r17_Supported},
}

const (
	UE_NR_Capability_v1700_MusimLeaveConnected_r17_Supported = 0
)

var uENRCapabilityV1700MusimLeaveConnectedR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_NR_Capability_v1700_MusimLeaveConnected_r17_Supported},
}

const (
	UE_NR_Capability_v1700_NonTerrestrialNetwork_r17_Supported = 0
)

var uENRCapabilityV1700NonTerrestrialNetworkR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_NR_Capability_v1700_NonTerrestrialNetwork_r17_Supported},
}

const (
	UE_NR_Capability_v1700_Ntn_ScenarioSupport_r17_Gso  = 0
	UE_NR_Capability_v1700_Ntn_ScenarioSupport_r17_Ngso = 1
)

var uENRCapabilityV1700NtnScenarioSupportR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_NR_Capability_v1700_Ntn_ScenarioSupport_r17_Gso, UE_NR_Capability_v1700_Ntn_ScenarioSupport_r17_Ngso},
}

const (
	UE_NR_Capability_v1700_SliceInfoforCellReselection_r17_Supported = 0
)

var uENRCapabilityV1700SliceInfoforCellReselectionR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_NR_Capability_v1700_SliceInfoforCellReselection_r17_Supported},
}

var uENRCapabilityV1700UlGapFR2PatternR17Constraints = per.FixedSize(4)

type UE_NR_Capability_v1700 struct {
	InactiveStatePO_Determination_r17       *int64
	HighSpeedParameters_v1700               *HighSpeedParameters_v1700
	PowSav_Parameters_v1700                 *PowSav_Parameters_v1700
	Mac_Parameters_v1700                    *MAC_Parameters_v1700
	Ims_Parameters_v1700                    *IMS_Parameters_v1700
	MeasAndMobParameters_v1700              MeasAndMobParameters_v1700
	AppLayerMeasParameters_r17              *AppLayerMeasParameters_r17
	RedCapParameters_r17                    *RedCapParameters_r17
	Ra_SDT_r17                              *int64
	Srb_SDT_r17                             *int64
	GNB_SideRTT_BasedPDC_r17                *int64
	Bh_RLF_DetectionRecovery_Indication_r17 *int64
	Nrdc_Parameters_v1700                   *NRDC_Parameters_v1700
	Bap_Parameters_v1700                    *BAP_Parameters_v1700
	Musim_GapPreference_r17                 *int64
	MusimLeaveConnected_r17                 *int64
	Mbs_Parameters_r17                      MBS_Parameters_r17
	NonTerrestrialNetwork_r17               *int64
	Ntn_ScenarioSupport_r17                 *int64
	SliceInfoforCellReselection_r17         *int64
	Ue_RadioPagingInfo_r17                  *UE_RadioPagingInfo_r17
	Ul_GapFR2_Pattern_r17                   *per.BitString
	Ntn_Parameters_r17                      *NTN_Parameters_r17
	NonCriticalExtension                    *UE_NR_Capability_v1740
}

func (ie *UE_NR_Capability_v1700) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uENRCapabilityV1700Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.InactiveStatePO_Determination_r17 != nil, ie.HighSpeedParameters_v1700 != nil, ie.PowSav_Parameters_v1700 != nil, ie.Mac_Parameters_v1700 != nil, ie.Ims_Parameters_v1700 != nil, ie.AppLayerMeasParameters_r17 != nil, ie.RedCapParameters_r17 != nil, ie.Ra_SDT_r17 != nil, ie.Srb_SDT_r17 != nil, ie.GNB_SideRTT_BasedPDC_r17 != nil, ie.Bh_RLF_DetectionRecovery_Indication_r17 != nil, ie.Nrdc_Parameters_v1700 != nil, ie.Bap_Parameters_v1700 != nil, ie.Musim_GapPreference_r17 != nil, ie.MusimLeaveConnected_r17 != nil, ie.NonTerrestrialNetwork_r17 != nil, ie.Ntn_ScenarioSupport_r17 != nil, ie.SliceInfoforCellReselection_r17 != nil, ie.Ue_RadioPagingInfo_r17 != nil, ie.Ul_GapFR2_Pattern_r17 != nil, ie.Ntn_Parameters_r17 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. inactiveStatePO-Determination-r17: enumerated
	{
		if ie.InactiveStatePO_Determination_r17 != nil {
			if err := e.EncodeEnumerated(*ie.InactiveStatePO_Determination_r17, uENRCapabilityV1700InactiveStatePODeterminationR17Constraints); err != nil {
				return err
			}
		}
	}

	// 3. highSpeedParameters-v1700: ref
	{
		if ie.HighSpeedParameters_v1700 != nil {
			if err := ie.HighSpeedParameters_v1700.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. powSav-Parameters-v1700: ref
	{
		if ie.PowSav_Parameters_v1700 != nil {
			if err := ie.PowSav_Parameters_v1700.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. mac-Parameters-v1700: ref
	{
		if ie.Mac_Parameters_v1700 != nil {
			if err := ie.Mac_Parameters_v1700.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. ims-Parameters-v1700: ref
	{
		if ie.Ims_Parameters_v1700 != nil {
			if err := ie.Ims_Parameters_v1700.Encode(e); err != nil {
				return err
			}
		}
	}

	// 7. measAndMobParameters-v1700: ref
	{
		if err := ie.MeasAndMobParameters_v1700.Encode(e); err != nil {
			return err
		}
	}

	// 8. appLayerMeasParameters-r17: ref
	{
		if ie.AppLayerMeasParameters_r17 != nil {
			if err := ie.AppLayerMeasParameters_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 9. redCapParameters-r17: ref
	{
		if ie.RedCapParameters_r17 != nil {
			if err := ie.RedCapParameters_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 10. ra-SDT-r17: enumerated
	{
		if ie.Ra_SDT_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Ra_SDT_r17, uENRCapabilityV1700RaSDTR17Constraints); err != nil {
				return err
			}
		}
	}

	// 11. srb-SDT-r17: enumerated
	{
		if ie.Srb_SDT_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Srb_SDT_r17, uENRCapabilityV1700SrbSDTR17Constraints); err != nil {
				return err
			}
		}
	}

	// 12. gNB-SideRTT-BasedPDC-r17: enumerated
	{
		if ie.GNB_SideRTT_BasedPDC_r17 != nil {
			if err := e.EncodeEnumerated(*ie.GNB_SideRTT_BasedPDC_r17, uENRCapabilityV1700GNBSideRTTBasedPDCR17Constraints); err != nil {
				return err
			}
		}
	}

	// 13. bh-RLF-DetectionRecovery-Indication-r17: enumerated
	{
		if ie.Bh_RLF_DetectionRecovery_Indication_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Bh_RLF_DetectionRecovery_Indication_r17, uENRCapabilityV1700BhRLFDetectionRecoveryIndicationR17Constraints); err != nil {
				return err
			}
		}
	}

	// 14. nrdc-Parameters-v1700: ref
	{
		if ie.Nrdc_Parameters_v1700 != nil {
			if err := ie.Nrdc_Parameters_v1700.Encode(e); err != nil {
				return err
			}
		}
	}

	// 15. bap-Parameters-v1700: ref
	{
		if ie.Bap_Parameters_v1700 != nil {
			if err := ie.Bap_Parameters_v1700.Encode(e); err != nil {
				return err
			}
		}
	}

	// 16. musim-GapPreference-r17: enumerated
	{
		if ie.Musim_GapPreference_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Musim_GapPreference_r17, uENRCapabilityV1700MusimGapPreferenceR17Constraints); err != nil {
				return err
			}
		}
	}

	// 17. musimLeaveConnected-r17: enumerated
	{
		if ie.MusimLeaveConnected_r17 != nil {
			if err := e.EncodeEnumerated(*ie.MusimLeaveConnected_r17, uENRCapabilityV1700MusimLeaveConnectedR17Constraints); err != nil {
				return err
			}
		}
	}

	// 18. mbs-Parameters-r17: ref
	{
		if err := ie.Mbs_Parameters_r17.Encode(e); err != nil {
			return err
		}
	}

	// 19. nonTerrestrialNetwork-r17: enumerated
	{
		if ie.NonTerrestrialNetwork_r17 != nil {
			if err := e.EncodeEnumerated(*ie.NonTerrestrialNetwork_r17, uENRCapabilityV1700NonTerrestrialNetworkR17Constraints); err != nil {
				return err
			}
		}
	}

	// 20. ntn-ScenarioSupport-r17: enumerated
	{
		if ie.Ntn_ScenarioSupport_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Ntn_ScenarioSupport_r17, uENRCapabilityV1700NtnScenarioSupportR17Constraints); err != nil {
				return err
			}
		}
	}

	// 21. sliceInfoforCellReselection-r17: enumerated
	{
		if ie.SliceInfoforCellReselection_r17 != nil {
			if err := e.EncodeEnumerated(*ie.SliceInfoforCellReselection_r17, uENRCapabilityV1700SliceInfoforCellReselectionR17Constraints); err != nil {
				return err
			}
		}
	}

	// 22. ue-RadioPagingInfo-r17: ref
	{
		if ie.Ue_RadioPagingInfo_r17 != nil {
			if err := ie.Ue_RadioPagingInfo_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 23. ul-GapFR2-Pattern-r17: bit-string
	{
		if ie.Ul_GapFR2_Pattern_r17 != nil {
			if err := e.EncodeBitString(*ie.Ul_GapFR2_Pattern_r17, uENRCapabilityV1700UlGapFR2PatternR17Constraints); err != nil {
				return err
			}
		}
	}

	// 24. ntn-Parameters-r17: ref
	{
		if ie.Ntn_Parameters_r17 != nil {
			if err := ie.Ntn_Parameters_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 25. nonCriticalExtension: ref
	{
		if ie.NonCriticalExtension != nil {
			if err := ie.NonCriticalExtension.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *UE_NR_Capability_v1700) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uENRCapabilityV1700Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. inactiveStatePO-Determination-r17: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(uENRCapabilityV1700InactiveStatePODeterminationR17Constraints)
			if err != nil {
				return err
			}
			ie.InactiveStatePO_Determination_r17 = &idx
		}
	}

	// 3. highSpeedParameters-v1700: ref
	{
		if seq.IsComponentPresent(1) {
			ie.HighSpeedParameters_v1700 = new(HighSpeedParameters_v1700)
			if err := ie.HighSpeedParameters_v1700.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. powSav-Parameters-v1700: ref
	{
		if seq.IsComponentPresent(2) {
			ie.PowSav_Parameters_v1700 = new(PowSav_Parameters_v1700)
			if err := ie.PowSav_Parameters_v1700.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. mac-Parameters-v1700: ref
	{
		if seq.IsComponentPresent(3) {
			ie.Mac_Parameters_v1700 = new(MAC_Parameters_v1700)
			if err := ie.Mac_Parameters_v1700.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. ims-Parameters-v1700: ref
	{
		if seq.IsComponentPresent(4) {
			ie.Ims_Parameters_v1700 = new(IMS_Parameters_v1700)
			if err := ie.Ims_Parameters_v1700.Decode(d); err != nil {
				return err
			}
		}
	}

	// 7. measAndMobParameters-v1700: ref
	{
		if err := ie.MeasAndMobParameters_v1700.Decode(d); err != nil {
			return err
		}
	}

	// 8. appLayerMeasParameters-r17: ref
	{
		if seq.IsComponentPresent(6) {
			ie.AppLayerMeasParameters_r17 = new(AppLayerMeasParameters_r17)
			if err := ie.AppLayerMeasParameters_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 9. redCapParameters-r17: ref
	{
		if seq.IsComponentPresent(7) {
			ie.RedCapParameters_r17 = new(RedCapParameters_r17)
			if err := ie.RedCapParameters_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 10. ra-SDT-r17: enumerated
	{
		if seq.IsComponentPresent(8) {
			idx, err := d.DecodeEnumerated(uENRCapabilityV1700RaSDTR17Constraints)
			if err != nil {
				return err
			}
			ie.Ra_SDT_r17 = &idx
		}
	}

	// 11. srb-SDT-r17: enumerated
	{
		if seq.IsComponentPresent(9) {
			idx, err := d.DecodeEnumerated(uENRCapabilityV1700SrbSDTR17Constraints)
			if err != nil {
				return err
			}
			ie.Srb_SDT_r17 = &idx
		}
	}

	// 12. gNB-SideRTT-BasedPDC-r17: enumerated
	{
		if seq.IsComponentPresent(10) {
			idx, err := d.DecodeEnumerated(uENRCapabilityV1700GNBSideRTTBasedPDCR17Constraints)
			if err != nil {
				return err
			}
			ie.GNB_SideRTT_BasedPDC_r17 = &idx
		}
	}

	// 13. bh-RLF-DetectionRecovery-Indication-r17: enumerated
	{
		if seq.IsComponentPresent(11) {
			idx, err := d.DecodeEnumerated(uENRCapabilityV1700BhRLFDetectionRecoveryIndicationR17Constraints)
			if err != nil {
				return err
			}
			ie.Bh_RLF_DetectionRecovery_Indication_r17 = &idx
		}
	}

	// 14. nrdc-Parameters-v1700: ref
	{
		if seq.IsComponentPresent(12) {
			ie.Nrdc_Parameters_v1700 = new(NRDC_Parameters_v1700)
			if err := ie.Nrdc_Parameters_v1700.Decode(d); err != nil {
				return err
			}
		}
	}

	// 15. bap-Parameters-v1700: ref
	{
		if seq.IsComponentPresent(13) {
			ie.Bap_Parameters_v1700 = new(BAP_Parameters_v1700)
			if err := ie.Bap_Parameters_v1700.Decode(d); err != nil {
				return err
			}
		}
	}

	// 16. musim-GapPreference-r17: enumerated
	{
		if seq.IsComponentPresent(14) {
			idx, err := d.DecodeEnumerated(uENRCapabilityV1700MusimGapPreferenceR17Constraints)
			if err != nil {
				return err
			}
			ie.Musim_GapPreference_r17 = &idx
		}
	}

	// 17. musimLeaveConnected-r17: enumerated
	{
		if seq.IsComponentPresent(15) {
			idx, err := d.DecodeEnumerated(uENRCapabilityV1700MusimLeaveConnectedR17Constraints)
			if err != nil {
				return err
			}
			ie.MusimLeaveConnected_r17 = &idx
		}
	}

	// 18. mbs-Parameters-r17: ref
	{
		if err := ie.Mbs_Parameters_r17.Decode(d); err != nil {
			return err
		}
	}

	// 19. nonTerrestrialNetwork-r17: enumerated
	{
		if seq.IsComponentPresent(17) {
			idx, err := d.DecodeEnumerated(uENRCapabilityV1700NonTerrestrialNetworkR17Constraints)
			if err != nil {
				return err
			}
			ie.NonTerrestrialNetwork_r17 = &idx
		}
	}

	// 20. ntn-ScenarioSupport-r17: enumerated
	{
		if seq.IsComponentPresent(18) {
			idx, err := d.DecodeEnumerated(uENRCapabilityV1700NtnScenarioSupportR17Constraints)
			if err != nil {
				return err
			}
			ie.Ntn_ScenarioSupport_r17 = &idx
		}
	}

	// 21. sliceInfoforCellReselection-r17: enumerated
	{
		if seq.IsComponentPresent(19) {
			idx, err := d.DecodeEnumerated(uENRCapabilityV1700SliceInfoforCellReselectionR17Constraints)
			if err != nil {
				return err
			}
			ie.SliceInfoforCellReselection_r17 = &idx
		}
	}

	// 22. ue-RadioPagingInfo-r17: ref
	{
		if seq.IsComponentPresent(20) {
			ie.Ue_RadioPagingInfo_r17 = new(UE_RadioPagingInfo_r17)
			if err := ie.Ue_RadioPagingInfo_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 23. ul-GapFR2-Pattern-r17: bit-string
	{
		if seq.IsComponentPresent(21) {
			v, err := d.DecodeBitString(uENRCapabilityV1700UlGapFR2PatternR17Constraints)
			if err != nil {
				return err
			}
			ie.Ul_GapFR2_Pattern_r17 = &v
		}
	}

	// 24. ntn-Parameters-r17: ref
	{
		if seq.IsComponentPresent(22) {
			ie.Ntn_Parameters_r17 = new(NTN_Parameters_r17)
			if err := ie.Ntn_Parameters_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 25. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(23) {
			ie.NonCriticalExtension = new(UE_NR_Capability_v1740)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
