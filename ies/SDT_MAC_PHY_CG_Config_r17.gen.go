// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SDT-MAC-PHY-CG-Config-r17 (line 1378).

var sDTMACPHYCGConfigR17Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "cg-SDT-ConfigLCH-RestrictionToAddModList-r17", Optional: true},
		{Name: "cg-SDT-ConfigLCH-RestrictionToReleaseList-r17", Optional: true},
		{Name: "cg-SDT-ConfigInitialBWP-NUL-r17", Optional: true},
		{Name: "cg-SDT-ConfigInitialBWP-SUL-r17", Optional: true},
		{Name: "cg-SDT-ConfigInitialBWP-DL-r17", Optional: true},
		{Name: "cg-SDT-TimeAlignmentTimer-r17", Optional: true},
		{Name: "cg-SDT-RSRP-ThresholdSSB-r17", Optional: true},
		{Name: "cg-SDT-TA-ValidationConfig-r17", Optional: true},
		{Name: "cg-SDT-CS-RNTI-r17", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

var sDTMACPHYCGConfigR17CgSDTConfigLCHRestrictionToAddModListR17Constraints = per.SizeRange(1, common.MaxLC_ID)

var sDTMACPHYCGConfigR17CgSDTConfigLCHRestrictionToReleaseListR17Constraints = per.SizeRange(1, common.MaxLC_ID)

var sDT_MAC_PHY_CG_Config_r17CgSDTConfigInitialBWPNULR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	SDT_MAC_PHY_CG_Config_r17_Cg_SDT_ConfigInitialBWP_NUL_r17_Release = 0
	SDT_MAC_PHY_CG_Config_r17_Cg_SDT_ConfigInitialBWP_NUL_r17_Setup   = 1
)

var sDT_MAC_PHY_CG_Config_r17CgSDTConfigInitialBWPSULR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	SDT_MAC_PHY_CG_Config_r17_Cg_SDT_ConfigInitialBWP_SUL_r17_Release = 0
	SDT_MAC_PHY_CG_Config_r17_Cg_SDT_ConfigInitialBWP_SUL_r17_Setup   = 1
)

var sDT_MAC_PHY_CG_Config_r17CgSDTTAValidationConfigR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	SDT_MAC_PHY_CG_Config_r17_Cg_SDT_TA_ValidationConfig_r17_Release = 0
	SDT_MAC_PHY_CG_Config_r17_Cg_SDT_TA_ValidationConfig_r17_Setup   = 1
)

var sDTMACPHYCGConfigR17ExtCgSDTConfigLCHRestrictionToAddModListExtV1800Constraints = per.SizeRange(1, common.MaxLC_ID)

const (
	SDT_MAC_PHY_CG_Config_r17_Ext_Cg_MT_SDT_MaxDurationToNextCG_Occasion_r18_Ms10    = 0
	SDT_MAC_PHY_CG_Config_r17_Ext_Cg_MT_SDT_MaxDurationToNextCG_Occasion_r18_Ms100   = 1
	SDT_MAC_PHY_CG_Config_r17_Ext_Cg_MT_SDT_MaxDurationToNextCG_Occasion_r18_Sec1    = 2
	SDT_MAC_PHY_CG_Config_r17_Ext_Cg_MT_SDT_MaxDurationToNextCG_Occasion_r18_Sec10   = 3
	SDT_MAC_PHY_CG_Config_r17_Ext_Cg_MT_SDT_MaxDurationToNextCG_Occasion_r18_Sec60   = 4
	SDT_MAC_PHY_CG_Config_r17_Ext_Cg_MT_SDT_MaxDurationToNextCG_Occasion_r18_Sec100  = 5
	SDT_MAC_PHY_CG_Config_r17_Ext_Cg_MT_SDT_MaxDurationToNextCG_Occasion_r18_Sec300  = 6
	SDT_MAC_PHY_CG_Config_r17_Ext_Cg_MT_SDT_MaxDurationToNextCG_Occasion_r18_Sec600  = 7
	SDT_MAC_PHY_CG_Config_r17_Ext_Cg_MT_SDT_MaxDurationToNextCG_Occasion_r18_Sec1200 = 8
	SDT_MAC_PHY_CG_Config_r17_Ext_Cg_MT_SDT_MaxDurationToNextCG_Occasion_r18_Sec1800 = 9
	SDT_MAC_PHY_CG_Config_r17_Ext_Cg_MT_SDT_MaxDurationToNextCG_Occasion_r18_Sec3600 = 10
	SDT_MAC_PHY_CG_Config_r17_Ext_Cg_MT_SDT_MaxDurationToNextCG_Occasion_r18_Spare5  = 11
	SDT_MAC_PHY_CG_Config_r17_Ext_Cg_MT_SDT_MaxDurationToNextCG_Occasion_r18_Spare4  = 12
	SDT_MAC_PHY_CG_Config_r17_Ext_Cg_MT_SDT_MaxDurationToNextCG_Occasion_r18_Spare3  = 13
	SDT_MAC_PHY_CG_Config_r17_Ext_Cg_MT_SDT_MaxDurationToNextCG_Occasion_r18_Spare2  = 14
	SDT_MAC_PHY_CG_Config_r17_Ext_Cg_MT_SDT_MaxDurationToNextCG_Occasion_r18_Spare1  = 15
)

var sDTMACPHYCGConfigR17ExtCgMTSDTMaxDurationToNextCGOccasionR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SDT_MAC_PHY_CG_Config_r17_Ext_Cg_MT_SDT_MaxDurationToNextCG_Occasion_r18_Ms10, SDT_MAC_PHY_CG_Config_r17_Ext_Cg_MT_SDT_MaxDurationToNextCG_Occasion_r18_Ms100, SDT_MAC_PHY_CG_Config_r17_Ext_Cg_MT_SDT_MaxDurationToNextCG_Occasion_r18_Sec1, SDT_MAC_PHY_CG_Config_r17_Ext_Cg_MT_SDT_MaxDurationToNextCG_Occasion_r18_Sec10, SDT_MAC_PHY_CG_Config_r17_Ext_Cg_MT_SDT_MaxDurationToNextCG_Occasion_r18_Sec60, SDT_MAC_PHY_CG_Config_r17_Ext_Cg_MT_SDT_MaxDurationToNextCG_Occasion_r18_Sec100, SDT_MAC_PHY_CG_Config_r17_Ext_Cg_MT_SDT_MaxDurationToNextCG_Occasion_r18_Sec300, SDT_MAC_PHY_CG_Config_r17_Ext_Cg_MT_SDT_MaxDurationToNextCG_Occasion_r18_Sec600, SDT_MAC_PHY_CG_Config_r17_Ext_Cg_MT_SDT_MaxDurationToNextCG_Occasion_r18_Sec1200, SDT_MAC_PHY_CG_Config_r17_Ext_Cg_MT_SDT_MaxDurationToNextCG_Occasion_r18_Sec1800, SDT_MAC_PHY_CG_Config_r17_Ext_Cg_MT_SDT_MaxDurationToNextCG_Occasion_r18_Sec3600, SDT_MAC_PHY_CG_Config_r17_Ext_Cg_MT_SDT_MaxDurationToNextCG_Occasion_r18_Spare5, SDT_MAC_PHY_CG_Config_r17_Ext_Cg_MT_SDT_MaxDurationToNextCG_Occasion_r18_Spare4, SDT_MAC_PHY_CG_Config_r17_Ext_Cg_MT_SDT_MaxDurationToNextCG_Occasion_r18_Spare3, SDT_MAC_PHY_CG_Config_r17_Ext_Cg_MT_SDT_MaxDurationToNextCG_Occasion_r18_Spare2, SDT_MAC_PHY_CG_Config_r17_Ext_Cg_MT_SDT_MaxDurationToNextCG_Occasion_r18_Spare1},
}

type SDT_MAC_PHY_CG_Config_r17 struct {
	Cg_SDT_ConfigLCH_RestrictionToAddModList_r17  []CG_SDT_ConfigLCH_Restriction_r17
	Cg_SDT_ConfigLCH_RestrictionToReleaseList_r17 []LogicalChannelIdentity
	Cg_SDT_ConfigInitialBWP_NUL_r17               *struct {
		Choice  int
		Release *struct{}
		Setup   *BWP_UplinkDedicatedSDT_r17
	}
	Cg_SDT_ConfigInitialBWP_SUL_r17 *struct {
		Choice  int
		Release *struct{}
		Setup   *BWP_UplinkDedicatedSDT_r17
	}
	Cg_SDT_ConfigInitialBWP_DL_r17 *BWP_DownlinkDedicatedSDT_r17
	Cg_SDT_TimeAlignmentTimer_r17  *TimeAlignmentTimer
	Cg_SDT_RSRP_ThresholdSSB_r17   *RSRP_Range
	Cg_SDT_TA_ValidationConfig_r17 *struct {
		Choice  int
		Release *struct{}
		Setup   *CG_SDT_TA_ValidationConfig_r17
	}
	Cg_SDT_CS_RNTI_r17                                *RNTI_Value
	Cg_SDT_ConfigLCH_RestrictionToAddModListExt_v1800 []CG_SDT_ConfigLCH_RestrictionExt_v1800
	Cg_MT_SDT_MaxDurationToNextCG_Occasion_r18        *int64
}

func (ie *SDT_MAC_PHY_CG_Config_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sDTMACPHYCGConfigR17Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Cg_SDT_ConfigLCH_RestrictionToAddModListExt_v1800 != nil || ie.Cg_MT_SDT_MaxDurationToNextCG_Occasion_r18 != nil
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Cg_SDT_ConfigLCH_RestrictionToAddModList_r17 != nil, ie.Cg_SDT_ConfigLCH_RestrictionToReleaseList_r17 != nil, ie.Cg_SDT_ConfigInitialBWP_NUL_r17 != nil, ie.Cg_SDT_ConfigInitialBWP_SUL_r17 != nil, ie.Cg_SDT_ConfigInitialBWP_DL_r17 != nil, ie.Cg_SDT_TimeAlignmentTimer_r17 != nil, ie.Cg_SDT_RSRP_ThresholdSSB_r17 != nil, ie.Cg_SDT_TA_ValidationConfig_r17 != nil, ie.Cg_SDT_CS_RNTI_r17 != nil}); err != nil {
		return err
	}

	// 3. cg-SDT-ConfigLCH-RestrictionToAddModList-r17: sequence-of
	{
		if ie.Cg_SDT_ConfigLCH_RestrictionToAddModList_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(sDTMACPHYCGConfigR17CgSDTConfigLCHRestrictionToAddModListR17Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Cg_SDT_ConfigLCH_RestrictionToAddModList_r17))); err != nil {
				return err
			}
			for i := range ie.Cg_SDT_ConfigLCH_RestrictionToAddModList_r17 {
				if err := ie.Cg_SDT_ConfigLCH_RestrictionToAddModList_r17[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 4. cg-SDT-ConfigLCH-RestrictionToReleaseList-r17: sequence-of
	{
		if ie.Cg_SDT_ConfigLCH_RestrictionToReleaseList_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(sDTMACPHYCGConfigR17CgSDTConfigLCHRestrictionToReleaseListR17Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Cg_SDT_ConfigLCH_RestrictionToReleaseList_r17))); err != nil {
				return err
			}
			for i := range ie.Cg_SDT_ConfigLCH_RestrictionToReleaseList_r17 {
				if err := ie.Cg_SDT_ConfigLCH_RestrictionToReleaseList_r17[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 5. cg-SDT-ConfigInitialBWP-NUL-r17: choice
	{
		if ie.Cg_SDT_ConfigInitialBWP_NUL_r17 != nil {
			choiceEnc := e.NewChoiceEncoder(sDT_MAC_PHY_CG_Config_r17CgSDTConfigInitialBWPNULR17Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Cg_SDT_ConfigInitialBWP_NUL_r17).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Cg_SDT_ConfigInitialBWP_NUL_r17).Choice {
			case SDT_MAC_PHY_CG_Config_r17_Cg_SDT_ConfigInitialBWP_NUL_r17_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case SDT_MAC_PHY_CG_Config_r17_Cg_SDT_ConfigInitialBWP_NUL_r17_Setup:
				if err := (*ie.Cg_SDT_ConfigInitialBWP_NUL_r17).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Cg_SDT_ConfigInitialBWP_NUL_r17).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 6. cg-SDT-ConfigInitialBWP-SUL-r17: choice
	{
		if ie.Cg_SDT_ConfigInitialBWP_SUL_r17 != nil {
			choiceEnc := e.NewChoiceEncoder(sDT_MAC_PHY_CG_Config_r17CgSDTConfigInitialBWPSULR17Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Cg_SDT_ConfigInitialBWP_SUL_r17).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Cg_SDT_ConfigInitialBWP_SUL_r17).Choice {
			case SDT_MAC_PHY_CG_Config_r17_Cg_SDT_ConfigInitialBWP_SUL_r17_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case SDT_MAC_PHY_CG_Config_r17_Cg_SDT_ConfigInitialBWP_SUL_r17_Setup:
				if err := (*ie.Cg_SDT_ConfigInitialBWP_SUL_r17).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Cg_SDT_ConfigInitialBWP_SUL_r17).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 7. cg-SDT-ConfigInitialBWP-DL-r17: ref
	{
		if ie.Cg_SDT_ConfigInitialBWP_DL_r17 != nil {
			if err := ie.Cg_SDT_ConfigInitialBWP_DL_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 8. cg-SDT-TimeAlignmentTimer-r17: ref
	{
		if ie.Cg_SDT_TimeAlignmentTimer_r17 != nil {
			if err := ie.Cg_SDT_TimeAlignmentTimer_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 9. cg-SDT-RSRP-ThresholdSSB-r17: ref
	{
		if ie.Cg_SDT_RSRP_ThresholdSSB_r17 != nil {
			if err := ie.Cg_SDT_RSRP_ThresholdSSB_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 10. cg-SDT-TA-ValidationConfig-r17: choice
	{
		if ie.Cg_SDT_TA_ValidationConfig_r17 != nil {
			choiceEnc := e.NewChoiceEncoder(sDT_MAC_PHY_CG_Config_r17CgSDTTAValidationConfigR17Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Cg_SDT_TA_ValidationConfig_r17).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Cg_SDT_TA_ValidationConfig_r17).Choice {
			case SDT_MAC_PHY_CG_Config_r17_Cg_SDT_TA_ValidationConfig_r17_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case SDT_MAC_PHY_CG_Config_r17_Cg_SDT_TA_ValidationConfig_r17_Setup:
				if err := (*ie.Cg_SDT_TA_ValidationConfig_r17).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Cg_SDT_TA_ValidationConfig_r17).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 11. cg-SDT-CS-RNTI-r17: ref
	{
		if ie.Cg_SDT_CS_RNTI_r17 != nil {
			if err := ie.Cg_SDT_CS_RNTI_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "cg-SDT-ConfigLCH-RestrictionToAddModListExt-v1800", Optional: true},
					{Name: "cg-MT-SDT-MaxDurationToNextCG-Occasion-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Cg_SDT_ConfigLCH_RestrictionToAddModListExt_v1800 != nil, ie.Cg_MT_SDT_MaxDurationToNextCG_Occasion_r18 != nil}); err != nil {
				return err
			}

			if ie.Cg_SDT_ConfigLCH_RestrictionToAddModListExt_v1800 != nil {
				seqOf := ex.NewSequenceOfEncoder(sDTMACPHYCGConfigR17ExtCgSDTConfigLCHRestrictionToAddModListExtV1800Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.Cg_SDT_ConfigLCH_RestrictionToAddModListExt_v1800))); err != nil {
					return err
				}
				for i := range ie.Cg_SDT_ConfigLCH_RestrictionToAddModListExt_v1800 {
					if err := ie.Cg_SDT_ConfigLCH_RestrictionToAddModListExt_v1800[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Cg_MT_SDT_MaxDurationToNextCG_Occasion_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Cg_MT_SDT_MaxDurationToNextCG_Occasion_r18, sDTMACPHYCGConfigR17ExtCgMTSDTMaxDurationToNextCGOccasionR18Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		if err := seq.EncodeExtensionAdditions(extPresence, extValues); err != nil {
			return err
		}
	}

	return nil
}

func (ie *SDT_MAC_PHY_CG_Config_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sDTMACPHYCGConfigR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. cg-SDT-ConfigLCH-RestrictionToAddModList-r17: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(sDTMACPHYCGConfigR17CgSDTConfigLCHRestrictionToAddModListR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Cg_SDT_ConfigLCH_RestrictionToAddModList_r17 = make([]CG_SDT_ConfigLCH_Restriction_r17, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Cg_SDT_ConfigLCH_RestrictionToAddModList_r17[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 4. cg-SDT-ConfigLCH-RestrictionToReleaseList-r17: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(sDTMACPHYCGConfigR17CgSDTConfigLCHRestrictionToReleaseListR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Cg_SDT_ConfigLCH_RestrictionToReleaseList_r17 = make([]LogicalChannelIdentity, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Cg_SDT_ConfigLCH_RestrictionToReleaseList_r17[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 5. cg-SDT-ConfigInitialBWP-NUL-r17: choice
	{
		if seq.IsComponentPresent(2) {
			ie.Cg_SDT_ConfigInitialBWP_NUL_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *BWP_UplinkDedicatedSDT_r17
			}{}
			choiceDec := d.NewChoiceDecoder(sDT_MAC_PHY_CG_Config_r17CgSDTConfigInitialBWPNULR17Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Cg_SDT_ConfigInitialBWP_NUL_r17).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case SDT_MAC_PHY_CG_Config_r17_Cg_SDT_ConfigInitialBWP_NUL_r17_Release:
				(*ie.Cg_SDT_ConfigInitialBWP_NUL_r17).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case SDT_MAC_PHY_CG_Config_r17_Cg_SDT_ConfigInitialBWP_NUL_r17_Setup:
				(*ie.Cg_SDT_ConfigInitialBWP_NUL_r17).Setup = new(BWP_UplinkDedicatedSDT_r17)
				if err := (*ie.Cg_SDT_ConfigInitialBWP_NUL_r17).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 6. cg-SDT-ConfigInitialBWP-SUL-r17: choice
	{
		if seq.IsComponentPresent(3) {
			ie.Cg_SDT_ConfigInitialBWP_SUL_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *BWP_UplinkDedicatedSDT_r17
			}{}
			choiceDec := d.NewChoiceDecoder(sDT_MAC_PHY_CG_Config_r17CgSDTConfigInitialBWPSULR17Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Cg_SDT_ConfigInitialBWP_SUL_r17).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case SDT_MAC_PHY_CG_Config_r17_Cg_SDT_ConfigInitialBWP_SUL_r17_Release:
				(*ie.Cg_SDT_ConfigInitialBWP_SUL_r17).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case SDT_MAC_PHY_CG_Config_r17_Cg_SDT_ConfigInitialBWP_SUL_r17_Setup:
				(*ie.Cg_SDT_ConfigInitialBWP_SUL_r17).Setup = new(BWP_UplinkDedicatedSDT_r17)
				if err := (*ie.Cg_SDT_ConfigInitialBWP_SUL_r17).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 7. cg-SDT-ConfigInitialBWP-DL-r17: ref
	{
		if seq.IsComponentPresent(4) {
			ie.Cg_SDT_ConfigInitialBWP_DL_r17 = new(BWP_DownlinkDedicatedSDT_r17)
			if err := ie.Cg_SDT_ConfigInitialBWP_DL_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 8. cg-SDT-TimeAlignmentTimer-r17: ref
	{
		if seq.IsComponentPresent(5) {
			ie.Cg_SDT_TimeAlignmentTimer_r17 = new(TimeAlignmentTimer)
			if err := ie.Cg_SDT_TimeAlignmentTimer_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 9. cg-SDT-RSRP-ThresholdSSB-r17: ref
	{
		if seq.IsComponentPresent(6) {
			ie.Cg_SDT_RSRP_ThresholdSSB_r17 = new(RSRP_Range)
			if err := ie.Cg_SDT_RSRP_ThresholdSSB_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 10. cg-SDT-TA-ValidationConfig-r17: choice
	{
		if seq.IsComponentPresent(7) {
			ie.Cg_SDT_TA_ValidationConfig_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *CG_SDT_TA_ValidationConfig_r17
			}{}
			choiceDec := d.NewChoiceDecoder(sDT_MAC_PHY_CG_Config_r17CgSDTTAValidationConfigR17Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Cg_SDT_TA_ValidationConfig_r17).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case SDT_MAC_PHY_CG_Config_r17_Cg_SDT_TA_ValidationConfig_r17_Release:
				(*ie.Cg_SDT_TA_ValidationConfig_r17).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case SDT_MAC_PHY_CG_Config_r17_Cg_SDT_TA_ValidationConfig_r17_Setup:
				(*ie.Cg_SDT_TA_ValidationConfig_r17).Setup = new(CG_SDT_TA_ValidationConfig_r17)
				if err := (*ie.Cg_SDT_TA_ValidationConfig_r17).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 11. cg-SDT-CS-RNTI-r17: ref
	{
		if seq.IsComponentPresent(8) {
			ie.Cg_SDT_CS_RNTI_r17 = new(RNTI_Value)
			if err := ie.Cg_SDT_CS_RNTI_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// Extension additions as open types.
	extValues, err := seq.DecodeExtensionAdditions()
	if err != nil {
		return err
	}
	extValueIndex := 0

	// Extension group 0:
	if seq.IsExtensionPresent(0) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "cg-SDT-ConfigLCH-RestrictionToAddModListExt-v1800", Optional: true},
				{Name: "cg-MT-SDT-MaxDurationToNextCG-Occasion-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			seqOf := dx.NewSequenceOfDecoder(sDTMACPHYCGConfigR17ExtCgSDTConfigLCHRestrictionToAddModListExtV1800Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Cg_SDT_ConfigLCH_RestrictionToAddModListExt_v1800 = make([]CG_SDT_ConfigLCH_RestrictionExt_v1800, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Cg_SDT_ConfigLCH_RestrictionToAddModListExt_v1800[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(sDTMACPHYCGConfigR17ExtCgMTSDTMaxDurationToNextCGOccasionR18Constraints)
			if err != nil {
				return err
			}
			ie.Cg_MT_SDT_MaxDurationToNextCG_Occasion_r18 = &v
		}
	}

	return nil
}
