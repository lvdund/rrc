// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SL-Unlicensed-r18 (line 26773).

var sLUnlicensedR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-LBT-FailureRecoveryConfig-r18", Optional: true},
		{Name: "sl-StartingSymbolFirst-r18", Optional: true},
		{Name: "sl-StartingSymbolSecond-r18", Optional: true},
		{Name: "sl-TransmissionStructureForPSCCHandPSSCH-r18", Optional: true},
		{Name: "sl-GapOfAdditionalSSSB-Occasion-r18", Optional: true},
		{Name: "sl-AbsoluteFrequencySSB-NonAnchorList-r18", Optional: true},
		{Name: "sl-CPE-StartingPositionS-SSB-r18", Optional: true},
		{Name: "sl-CWS-ForPsschWithoutHarqAck-r18", Optional: true},
		{Name: "sl-NumOfAdditionalSSSBOccasion-r18", Optional: true},
		{Name: "sl-SSSBPowerOffsetOfAnchorRBSet-r18", Optional: true},
		{Name: "sl-RBSetConfigList-r18", Optional: true},
		{Name: "sl-IntraCellGuardBandsSL-List-r18", Optional: true},
	},
}

var sL_Unlicensed_r18SlLBTFailureRecoveryConfigR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	SL_Unlicensed_r18_Sl_LBT_FailureRecoveryConfig_r18_Release = 0
	SL_Unlicensed_r18_Sl_LBT_FailureRecoveryConfig_r18_Setup   = 1
)

const (
	SL_Unlicensed_r18_Sl_StartingSymbolFirst_r18_Sym0 = 0
	SL_Unlicensed_r18_Sl_StartingSymbolFirst_r18_Sym1 = 1
	SL_Unlicensed_r18_Sl_StartingSymbolFirst_r18_Sym2 = 2
	SL_Unlicensed_r18_Sl_StartingSymbolFirst_r18_Sym3 = 3
	SL_Unlicensed_r18_Sl_StartingSymbolFirst_r18_Sym4 = 4
	SL_Unlicensed_r18_Sl_StartingSymbolFirst_r18_Sym5 = 5
	SL_Unlicensed_r18_Sl_StartingSymbolFirst_r18_Sym6 = 6
)

var sLUnlicensedR18SlStartingSymbolFirstR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_Unlicensed_r18_Sl_StartingSymbolFirst_r18_Sym0, SL_Unlicensed_r18_Sl_StartingSymbolFirst_r18_Sym1, SL_Unlicensed_r18_Sl_StartingSymbolFirst_r18_Sym2, SL_Unlicensed_r18_Sl_StartingSymbolFirst_r18_Sym3, SL_Unlicensed_r18_Sl_StartingSymbolFirst_r18_Sym4, SL_Unlicensed_r18_Sl_StartingSymbolFirst_r18_Sym5, SL_Unlicensed_r18_Sl_StartingSymbolFirst_r18_Sym6},
}

const (
	SL_Unlicensed_r18_Sl_StartingSymbolSecond_r18_Sym3 = 0
	SL_Unlicensed_r18_Sl_StartingSymbolSecond_r18_Sym4 = 1
	SL_Unlicensed_r18_Sl_StartingSymbolSecond_r18_Sym5 = 2
	SL_Unlicensed_r18_Sl_StartingSymbolSecond_r18_Sym6 = 3
	SL_Unlicensed_r18_Sl_StartingSymbolSecond_r18_Sym7 = 4
)

var sLUnlicensedR18SlStartingSymbolSecondR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_Unlicensed_r18_Sl_StartingSymbolSecond_r18_Sym3, SL_Unlicensed_r18_Sl_StartingSymbolSecond_r18_Sym4, SL_Unlicensed_r18_Sl_StartingSymbolSecond_r18_Sym5, SL_Unlicensed_r18_Sl_StartingSymbolSecond_r18_Sym6, SL_Unlicensed_r18_Sl_StartingSymbolSecond_r18_Sym7},
}

const (
	SL_Unlicensed_r18_Sl_TransmissionStructureForPSCCHandPSSCH_r18_ContiguousRB = 0
	SL_Unlicensed_r18_Sl_TransmissionStructureForPSCCHandPSSCH_r18_InterlaceRB  = 1
)

var sLUnlicensedR18SlTransmissionStructureForPSCCHandPSSCHR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_Unlicensed_r18_Sl_TransmissionStructureForPSCCHandPSSCH_r18_ContiguousRB, SL_Unlicensed_r18_Sl_TransmissionStructureForPSCCHandPSSCH_r18_InterlaceRB},
}

var sLUnlicensedR18SlGapOfAdditionalSSSBOccasionR18Constraints = per.Constrained(0, 639)

var sLUnlicensedR18SlAbsoluteFrequencySSBNonAnchorListR18Constraints = per.SizeRange(1, common.MaxSL_NonAnchorRBsets)

var sLUnlicensedR18SlCPEStartingPositionSSSBR18Constraints = per.Constrained(1, 9)

const (
	SL_Unlicensed_r18_Sl_CWS_ForPsschWithoutHarqAck_r18_T1       = 0
	SL_Unlicensed_r18_Sl_CWS_ForPsschWithoutHarqAck_r18_T8       = 1
	SL_Unlicensed_r18_Sl_CWS_ForPsschWithoutHarqAck_r18_T16      = 2
	SL_Unlicensed_r18_Sl_CWS_ForPsschWithoutHarqAck_r18_T32      = 3
	SL_Unlicensed_r18_Sl_CWS_ForPsschWithoutHarqAck_r18_Infinity = 4
)

var sLUnlicensedR18SlCWSForPsschWithoutHarqAckR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_Unlicensed_r18_Sl_CWS_ForPsschWithoutHarqAck_r18_T1, SL_Unlicensed_r18_Sl_CWS_ForPsschWithoutHarqAck_r18_T8, SL_Unlicensed_r18_Sl_CWS_ForPsschWithoutHarqAck_r18_T16, SL_Unlicensed_r18_Sl_CWS_ForPsschWithoutHarqAck_r18_T32, SL_Unlicensed_r18_Sl_CWS_ForPsschWithoutHarqAck_r18_Infinity},
}

var sLUnlicensedR18SlNumOfAdditionalSSSBOccasionR18Constraints = per.Constrained(0, 4)

const (
	SL_Unlicensed_r18_Sl_SSSBPowerOffsetOfAnchorRBSet_r18_Value1 = 0
	SL_Unlicensed_r18_Sl_SSSBPowerOffsetOfAnchorRBSet_r18_Value2 = 1
)

var sLUnlicensedR18SlSSSBPowerOffsetOfAnchorRBSetR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_Unlicensed_r18_Sl_SSSBPowerOffsetOfAnchorRBSet_r18_Value1, SL_Unlicensed_r18_Sl_SSSBPowerOffsetOfAnchorRBSet_r18_Value2},
}

var sLUnlicensedR18SlRBSetConfigListR18Constraints = per.SizeRange(1, 5)

var sLUnlicensedR18SlIntraCellGuardBandsSLListR18Constraints = per.SizeRange(1, common.MaxSCSs)

type SL_Unlicensed_r18 struct {
	Sl_LBT_FailureRecoveryConfig_r18 *struct {
		Choice  int
		Release *struct{}
		Setup   *SL_LBT_FailureRecoveryConfig_r18
	}
	Sl_StartingSymbolFirst_r18                   *int64
	Sl_StartingSymbolSecond_r18                  *int64
	Sl_TransmissionStructureForPSCCHandPSSCH_r18 *int64
	Sl_GapOfAdditionalSSSB_Occasion_r18          *int64
	Sl_AbsoluteFrequencySSB_NonAnchorList_r18    []ARFCN_ValueNR
	Sl_CPE_StartingPositionS_SSB_r18             *int64
	Sl_CWS_ForPsschWithoutHarqAck_r18            *int64
	Sl_NumOfAdditionalSSSBOccasion_r18           *int64
	Sl_SSSBPowerOffsetOfAnchorRBSet_r18          *int64
	Sl_RBSetConfigList_r18                       []SL_RBSetConfig_r18
	Sl_IntraCellGuardBandsSL_List_r18            []IntraCellGuardBandsPerSCS_r16
}

func (ie *SL_Unlicensed_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLUnlicensedR18Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_LBT_FailureRecoveryConfig_r18 != nil, ie.Sl_StartingSymbolFirst_r18 != nil, ie.Sl_StartingSymbolSecond_r18 != nil, ie.Sl_TransmissionStructureForPSCCHandPSSCH_r18 != nil, ie.Sl_GapOfAdditionalSSSB_Occasion_r18 != nil, ie.Sl_AbsoluteFrequencySSB_NonAnchorList_r18 != nil, ie.Sl_CPE_StartingPositionS_SSB_r18 != nil, ie.Sl_CWS_ForPsschWithoutHarqAck_r18 != nil, ie.Sl_NumOfAdditionalSSSBOccasion_r18 != nil, ie.Sl_SSSBPowerOffsetOfAnchorRBSet_r18 != nil, ie.Sl_RBSetConfigList_r18 != nil, ie.Sl_IntraCellGuardBandsSL_List_r18 != nil}); err != nil {
		return err
	}

	// 2. sl-LBT-FailureRecoveryConfig-r18: choice
	{
		if ie.Sl_LBT_FailureRecoveryConfig_r18 != nil {
			choiceEnc := e.NewChoiceEncoder(sL_Unlicensed_r18SlLBTFailureRecoveryConfigR18Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Sl_LBT_FailureRecoveryConfig_r18).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Sl_LBT_FailureRecoveryConfig_r18).Choice {
			case SL_Unlicensed_r18_Sl_LBT_FailureRecoveryConfig_r18_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case SL_Unlicensed_r18_Sl_LBT_FailureRecoveryConfig_r18_Setup:
				if err := (*ie.Sl_LBT_FailureRecoveryConfig_r18).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Sl_LBT_FailureRecoveryConfig_r18).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 3. sl-StartingSymbolFirst-r18: enumerated
	{
		if ie.Sl_StartingSymbolFirst_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_StartingSymbolFirst_r18, sLUnlicensedR18SlStartingSymbolFirstR18Constraints); err != nil {
				return err
			}
		}
	}

	// 4. sl-StartingSymbolSecond-r18: enumerated
	{
		if ie.Sl_StartingSymbolSecond_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_StartingSymbolSecond_r18, sLUnlicensedR18SlStartingSymbolSecondR18Constraints); err != nil {
				return err
			}
		}
	}

	// 5. sl-TransmissionStructureForPSCCHandPSSCH-r18: enumerated
	{
		if ie.Sl_TransmissionStructureForPSCCHandPSSCH_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_TransmissionStructureForPSCCHandPSSCH_r18, sLUnlicensedR18SlTransmissionStructureForPSCCHandPSSCHR18Constraints); err != nil {
				return err
			}
		}
	}

	// 6. sl-GapOfAdditionalSSSB-Occasion-r18: integer
	{
		if ie.Sl_GapOfAdditionalSSSB_Occasion_r18 != nil {
			if err := e.EncodeInteger(*ie.Sl_GapOfAdditionalSSSB_Occasion_r18, sLUnlicensedR18SlGapOfAdditionalSSSBOccasionR18Constraints); err != nil {
				return err
			}
		}
	}

	// 7. sl-AbsoluteFrequencySSB-NonAnchorList-r18: sequence-of
	{
		if ie.Sl_AbsoluteFrequencySSB_NonAnchorList_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(sLUnlicensedR18SlAbsoluteFrequencySSBNonAnchorListR18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sl_AbsoluteFrequencySSB_NonAnchorList_r18))); err != nil {
				return err
			}
			for i := range ie.Sl_AbsoluteFrequencySSB_NonAnchorList_r18 {
				if err := ie.Sl_AbsoluteFrequencySSB_NonAnchorList_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 8. sl-CPE-StartingPositionS-SSB-r18: integer
	{
		if ie.Sl_CPE_StartingPositionS_SSB_r18 != nil {
			if err := e.EncodeInteger(*ie.Sl_CPE_StartingPositionS_SSB_r18, sLUnlicensedR18SlCPEStartingPositionSSSBR18Constraints); err != nil {
				return err
			}
		}
	}

	// 9. sl-CWS-ForPsschWithoutHarqAck-r18: enumerated
	{
		if ie.Sl_CWS_ForPsschWithoutHarqAck_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_CWS_ForPsschWithoutHarqAck_r18, sLUnlicensedR18SlCWSForPsschWithoutHarqAckR18Constraints); err != nil {
				return err
			}
		}
	}

	// 10. sl-NumOfAdditionalSSSBOccasion-r18: integer
	{
		if ie.Sl_NumOfAdditionalSSSBOccasion_r18 != nil {
			if err := e.EncodeInteger(*ie.Sl_NumOfAdditionalSSSBOccasion_r18, sLUnlicensedR18SlNumOfAdditionalSSSBOccasionR18Constraints); err != nil {
				return err
			}
		}
	}

	// 11. sl-SSSBPowerOffsetOfAnchorRBSet-r18: enumerated
	{
		if ie.Sl_SSSBPowerOffsetOfAnchorRBSet_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_SSSBPowerOffsetOfAnchorRBSet_r18, sLUnlicensedR18SlSSSBPowerOffsetOfAnchorRBSetR18Constraints); err != nil {
				return err
			}
		}
	}

	// 12. sl-RBSetConfigList-r18: sequence-of
	{
		if ie.Sl_RBSetConfigList_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(sLUnlicensedR18SlRBSetConfigListR18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sl_RBSetConfigList_r18))); err != nil {
				return err
			}
			for i := range ie.Sl_RBSetConfigList_r18 {
				if err := ie.Sl_RBSetConfigList_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 13. sl-IntraCellGuardBandsSL-List-r18: sequence-of
	{
		if ie.Sl_IntraCellGuardBandsSL_List_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(sLUnlicensedR18SlIntraCellGuardBandsSLListR18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sl_IntraCellGuardBandsSL_List_r18))); err != nil {
				return err
			}
			for i := range ie.Sl_IntraCellGuardBandsSL_List_r18 {
				if err := ie.Sl_IntraCellGuardBandsSL_List_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *SL_Unlicensed_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLUnlicensedR18Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. sl-LBT-FailureRecoveryConfig-r18: choice
	{
		if seq.IsComponentPresent(0) {
			ie.Sl_LBT_FailureRecoveryConfig_r18 = &struct {
				Choice  int
				Release *struct{}
				Setup   *SL_LBT_FailureRecoveryConfig_r18
			}{}
			choiceDec := d.NewChoiceDecoder(sL_Unlicensed_r18SlLBTFailureRecoveryConfigR18Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Sl_LBT_FailureRecoveryConfig_r18).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case SL_Unlicensed_r18_Sl_LBT_FailureRecoveryConfig_r18_Release:
				(*ie.Sl_LBT_FailureRecoveryConfig_r18).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case SL_Unlicensed_r18_Sl_LBT_FailureRecoveryConfig_r18_Setup:
				(*ie.Sl_LBT_FailureRecoveryConfig_r18).Setup = new(SL_LBT_FailureRecoveryConfig_r18)
				if err := (*ie.Sl_LBT_FailureRecoveryConfig_r18).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 3. sl-StartingSymbolFirst-r18: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(sLUnlicensedR18SlStartingSymbolFirstR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_StartingSymbolFirst_r18 = &idx
		}
	}

	// 4. sl-StartingSymbolSecond-r18: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(sLUnlicensedR18SlStartingSymbolSecondR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_StartingSymbolSecond_r18 = &idx
		}
	}

	// 5. sl-TransmissionStructureForPSCCHandPSSCH-r18: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(sLUnlicensedR18SlTransmissionStructureForPSCCHandPSSCHR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_TransmissionStructureForPSCCHandPSSCH_r18 = &idx
		}
	}

	// 6. sl-GapOfAdditionalSSSB-Occasion-r18: integer
	{
		if seq.IsComponentPresent(4) {
			v, err := d.DecodeInteger(sLUnlicensedR18SlGapOfAdditionalSSSBOccasionR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_GapOfAdditionalSSSB_Occasion_r18 = &v
		}
	}

	// 7. sl-AbsoluteFrequencySSB-NonAnchorList-r18: sequence-of
	{
		if seq.IsComponentPresent(5) {
			seqOf := d.NewSequenceOfDecoder(sLUnlicensedR18SlAbsoluteFrequencySSBNonAnchorListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_AbsoluteFrequencySSB_NonAnchorList_r18 = make([]ARFCN_ValueNR, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_AbsoluteFrequencySSB_NonAnchorList_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 8. sl-CPE-StartingPositionS-SSB-r18: integer
	{
		if seq.IsComponentPresent(6) {
			v, err := d.DecodeInteger(sLUnlicensedR18SlCPEStartingPositionSSSBR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_CPE_StartingPositionS_SSB_r18 = &v
		}
	}

	// 9. sl-CWS-ForPsschWithoutHarqAck-r18: enumerated
	{
		if seq.IsComponentPresent(7) {
			idx, err := d.DecodeEnumerated(sLUnlicensedR18SlCWSForPsschWithoutHarqAckR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_CWS_ForPsschWithoutHarqAck_r18 = &idx
		}
	}

	// 10. sl-NumOfAdditionalSSSBOccasion-r18: integer
	{
		if seq.IsComponentPresent(8) {
			v, err := d.DecodeInteger(sLUnlicensedR18SlNumOfAdditionalSSSBOccasionR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_NumOfAdditionalSSSBOccasion_r18 = &v
		}
	}

	// 11. sl-SSSBPowerOffsetOfAnchorRBSet-r18: enumerated
	{
		if seq.IsComponentPresent(9) {
			idx, err := d.DecodeEnumerated(sLUnlicensedR18SlSSSBPowerOffsetOfAnchorRBSetR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_SSSBPowerOffsetOfAnchorRBSet_r18 = &idx
		}
	}

	// 12. sl-RBSetConfigList-r18: sequence-of
	{
		if seq.IsComponentPresent(10) {
			seqOf := d.NewSequenceOfDecoder(sLUnlicensedR18SlRBSetConfigListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_RBSetConfigList_r18 = make([]SL_RBSetConfig_r18, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_RBSetConfigList_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 13. sl-IntraCellGuardBandsSL-List-r18: sequence-of
	{
		if seq.IsComponentPresent(11) {
			seqOf := d.NewSequenceOfDecoder(sLUnlicensedR18SlIntraCellGuardBandsSLListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_IntraCellGuardBandsSL_List_r18 = make([]IntraCellGuardBandsPerSCS_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_IntraCellGuardBandsSL_List_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
