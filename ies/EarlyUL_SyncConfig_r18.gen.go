// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: EarlyUL-SyncConfig-r18 (line 8237).

var earlyULSyncConfigR18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "frequencyInfoUL-r18"},
		{Name: "rach-ConfigGeneric-r18"},
		{Name: "bwp-GenericParameters-r18"},
		{Name: "ssb-PerRACH-Occasion-r18", Optional: true},
		{Name: "prach-RootSequenceIndex-r18", Optional: true},
		{Name: "ltm-PRACH-SubcarrierSpacing-r18", Optional: true},
		{Name: "n-TimingAdvanceOffset-r18", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
	},
}

const (
	EarlyUL_SyncConfig_r18_Ssb_PerRACH_Occasion_r18_OneEighth = 0
	EarlyUL_SyncConfig_r18_Ssb_PerRACH_Occasion_r18_OneFourth = 1
	EarlyUL_SyncConfig_r18_Ssb_PerRACH_Occasion_r18_OneHalf   = 2
	EarlyUL_SyncConfig_r18_Ssb_PerRACH_Occasion_r18_One       = 3
	EarlyUL_SyncConfig_r18_Ssb_PerRACH_Occasion_r18_Two       = 4
	EarlyUL_SyncConfig_r18_Ssb_PerRACH_Occasion_r18_Four      = 5
	EarlyUL_SyncConfig_r18_Ssb_PerRACH_Occasion_r18_Eight     = 6
	EarlyUL_SyncConfig_r18_Ssb_PerRACH_Occasion_r18_Sixteen   = 7
)

var earlyULSyncConfigR18SsbPerRACHOccasionR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{EarlyUL_SyncConfig_r18_Ssb_PerRACH_Occasion_r18_OneEighth, EarlyUL_SyncConfig_r18_Ssb_PerRACH_Occasion_r18_OneFourth, EarlyUL_SyncConfig_r18_Ssb_PerRACH_Occasion_r18_OneHalf, EarlyUL_SyncConfig_r18_Ssb_PerRACH_Occasion_r18_One, EarlyUL_SyncConfig_r18_Ssb_PerRACH_Occasion_r18_Two, EarlyUL_SyncConfig_r18_Ssb_PerRACH_Occasion_r18_Four, EarlyUL_SyncConfig_r18_Ssb_PerRACH_Occasion_r18_Eight, EarlyUL_SyncConfig_r18_Ssb_PerRACH_Occasion_r18_Sixteen},
}

var earlyUL_SyncConfig_r18PrachRootSequenceIndexR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "l839"},
		{Name: "l139"},
	},
}

const (
	EarlyUL_SyncConfig_r18_Prach_RootSequenceIndex_r18_L839 = 0
	EarlyUL_SyncConfig_r18_Prach_RootSequenceIndex_r18_L139 = 1
)

const (
	EarlyUL_SyncConfig_r18_N_TimingAdvanceOffset_r18_N0     = 0
	EarlyUL_SyncConfig_r18_N_TimingAdvanceOffset_r18_N25600 = 1
	EarlyUL_SyncConfig_r18_N_TimingAdvanceOffset_r18_N39936 = 2
	EarlyUL_SyncConfig_r18_N_TimingAdvanceOffset_r18_Spare1 = 3
)

var earlyULSyncConfigR18NTimingAdvanceOffsetR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{EarlyUL_SyncConfig_r18_N_TimingAdvanceOffset_r18_N0, EarlyUL_SyncConfig_r18_N_TimingAdvanceOffset_r18_N25600, EarlyUL_SyncConfig_r18_N_TimingAdvanceOffset_r18_N39936, EarlyUL_SyncConfig_r18_N_TimingAdvanceOffset_r18_Spare1},
}

const (
	EarlyUL_SyncConfig_r18_Ext_Ltm_RestrictedSetConfig_r18_UnrestrictedSet    = 0
	EarlyUL_SyncConfig_r18_Ext_Ltm_RestrictedSetConfig_r18_RestrictedSetTypeA = 1
	EarlyUL_SyncConfig_r18_Ext_Ltm_RestrictedSetConfig_r18_RestrictedSetTypeB = 2
)

var earlyULSyncConfigR18ExtLtmRestrictedSetConfigR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{EarlyUL_SyncConfig_r18_Ext_Ltm_RestrictedSetConfig_r18_UnrestrictedSet, EarlyUL_SyncConfig_r18_Ext_Ltm_RestrictedSetConfig_r18_RestrictedSetTypeA, EarlyUL_SyncConfig_r18_Ext_Ltm_RestrictedSetConfig_r18_RestrictedSetTypeB},
}

type EarlyUL_SyncConfig_r18 struct {
	FrequencyInfoUL_r18         FrequencyInfoUL
	Rach_ConfigGeneric_r18      RACH_ConfigGeneric
	Bwp_GenericParameters_r18   BWP
	Ssb_PerRACH_Occasion_r18    *int64
	Prach_RootSequenceIndex_r18 *struct {
		Choice int
		L839   *int64
		L139   *int64
	}
	Ltm_PRACH_SubcarrierSpacing_r18       *SubcarrierSpacing
	N_TimingAdvanceOffset_r18             *int64
	Ltm_Tdd_UL_DL_ConfigurationCommon_r18 *TDD_UL_DL_ConfigCommon
	Ltm_RestrictedSetConfig_r18           *int64
	Ltm_TimeAlignmentTimer_r19            *TimeAlignmentTimer
	Ltm_TimeAlignmentTimerTag2_r19        *TimeAlignmentTimer
}

func (ie *EarlyUL_SyncConfig_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(earlyULSyncConfigR18Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Ltm_Tdd_UL_DL_ConfigurationCommon_r18 != nil || ie.Ltm_RestrictedSetConfig_r18 != nil
	hasExtGroup1 := ie.Ltm_TimeAlignmentTimer_r19 != nil || ie.Ltm_TimeAlignmentTimerTag2_r19 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ssb_PerRACH_Occasion_r18 != nil, ie.Prach_RootSequenceIndex_r18 != nil, ie.Ltm_PRACH_SubcarrierSpacing_r18 != nil, ie.N_TimingAdvanceOffset_r18 != nil}); err != nil {
		return err
	}

	// 3. frequencyInfoUL-r18: ref
	{
		if err := ie.FrequencyInfoUL_r18.Encode(e); err != nil {
			return err
		}
	}

	// 4. rach-ConfigGeneric-r18: ref
	{
		if err := ie.Rach_ConfigGeneric_r18.Encode(e); err != nil {
			return err
		}
	}

	// 5. bwp-GenericParameters-r18: ref
	{
		if err := ie.Bwp_GenericParameters_r18.Encode(e); err != nil {
			return err
		}
	}

	// 6. ssb-PerRACH-Occasion-r18: enumerated
	{
		if ie.Ssb_PerRACH_Occasion_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Ssb_PerRACH_Occasion_r18, earlyULSyncConfigR18SsbPerRACHOccasionR18Constraints); err != nil {
				return err
			}
		}
	}

	// 7. prach-RootSequenceIndex-r18: choice
	{
		if ie.Prach_RootSequenceIndex_r18 != nil {
			choiceEnc := e.NewChoiceEncoder(earlyUL_SyncConfig_r18PrachRootSequenceIndexR18Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Prach_RootSequenceIndex_r18).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Prach_RootSequenceIndex_r18).Choice {
			case EarlyUL_SyncConfig_r18_Prach_RootSequenceIndex_r18_L839:
				if err := e.EncodeInteger((*(*ie.Prach_RootSequenceIndex_r18).L839), per.Constrained(0, 837)); err != nil {
					return err
				}
			case EarlyUL_SyncConfig_r18_Prach_RootSequenceIndex_r18_L139:
				if err := e.EncodeInteger((*(*ie.Prach_RootSequenceIndex_r18).L139), per.Constrained(0, 137)); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Prach_RootSequenceIndex_r18).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 8. ltm-PRACH-SubcarrierSpacing-r18: ref
	{
		if ie.Ltm_PRACH_SubcarrierSpacing_r18 != nil {
			if err := ie.Ltm_PRACH_SubcarrierSpacing_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 9. n-TimingAdvanceOffset-r18: enumerated
	{
		if ie.N_TimingAdvanceOffset_r18 != nil {
			if err := e.EncodeEnumerated(*ie.N_TimingAdvanceOffset_r18, earlyULSyncConfigR18NTimingAdvanceOffsetR18Constraints); err != nil {
				return err
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0, hasExtGroup1}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "ltm-tdd-UL-DL-ConfigurationCommon-r18", Optional: true},
					{Name: "ltm-restrictedSetConfig-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Ltm_Tdd_UL_DL_ConfigurationCommon_r18 != nil, ie.Ltm_RestrictedSetConfig_r18 != nil}); err != nil {
				return err
			}

			if ie.Ltm_Tdd_UL_DL_ConfigurationCommon_r18 != nil {
				if err := ie.Ltm_Tdd_UL_DL_ConfigurationCommon_r18.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Ltm_RestrictedSetConfig_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Ltm_RestrictedSetConfig_r18, earlyULSyncConfigR18ExtLtmRestrictedSetConfigR18Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 1:
		if hasExtGroup1 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "ltm-TimeAlignmentTimer-r19", Optional: true},
					{Name: "ltm-TimeAlignmentTimerTag2-r19", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Ltm_TimeAlignmentTimer_r19 != nil, ie.Ltm_TimeAlignmentTimerTag2_r19 != nil}); err != nil {
				return err
			}

			if ie.Ltm_TimeAlignmentTimer_r19 != nil {
				if err := ie.Ltm_TimeAlignmentTimer_r19.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Ltm_TimeAlignmentTimerTag2_r19 != nil {
				if err := ie.Ltm_TimeAlignmentTimerTag2_r19.Encode(ex); err != nil {
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

func (ie *EarlyUL_SyncConfig_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(earlyULSyncConfigR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. frequencyInfoUL-r18: ref
	{
		if err := ie.FrequencyInfoUL_r18.Decode(d); err != nil {
			return err
		}
	}

	// 4. rach-ConfigGeneric-r18: ref
	{
		if err := ie.Rach_ConfigGeneric_r18.Decode(d); err != nil {
			return err
		}
	}

	// 5. bwp-GenericParameters-r18: ref
	{
		if err := ie.Bwp_GenericParameters_r18.Decode(d); err != nil {
			return err
		}
	}

	// 6. ssb-PerRACH-Occasion-r18: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(earlyULSyncConfigR18SsbPerRACHOccasionR18Constraints)
			if err != nil {
				return err
			}
			ie.Ssb_PerRACH_Occasion_r18 = &idx
		}
	}

	// 7. prach-RootSequenceIndex-r18: choice
	{
		if seq.IsComponentPresent(4) {
			ie.Prach_RootSequenceIndex_r18 = &struct {
				Choice int
				L839   *int64
				L139   *int64
			}{}
			choiceDec := d.NewChoiceDecoder(earlyUL_SyncConfig_r18PrachRootSequenceIndexR18Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Prach_RootSequenceIndex_r18).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case EarlyUL_SyncConfig_r18_Prach_RootSequenceIndex_r18_L839:
				(*ie.Prach_RootSequenceIndex_r18).L839 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 837))
				if err != nil {
					return err
				}
				(*(*ie.Prach_RootSequenceIndex_r18).L839) = v
			case EarlyUL_SyncConfig_r18_Prach_RootSequenceIndex_r18_L139:
				(*ie.Prach_RootSequenceIndex_r18).L139 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 137))
				if err != nil {
					return err
				}
				(*(*ie.Prach_RootSequenceIndex_r18).L139) = v
			}
		}
	}

	// 8. ltm-PRACH-SubcarrierSpacing-r18: ref
	{
		if seq.IsComponentPresent(5) {
			ie.Ltm_PRACH_SubcarrierSpacing_r18 = new(SubcarrierSpacing)
			if err := ie.Ltm_PRACH_SubcarrierSpacing_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 9. n-TimingAdvanceOffset-r18: enumerated
	{
		if seq.IsComponentPresent(6) {
			idx, err := d.DecodeEnumerated(earlyULSyncConfigR18NTimingAdvanceOffsetR18Constraints)
			if err != nil {
				return err
			}
			ie.N_TimingAdvanceOffset_r18 = &idx
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
				{Name: "ltm-tdd-UL-DL-ConfigurationCommon-r18", Optional: true},
				{Name: "ltm-restrictedSetConfig-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Ltm_Tdd_UL_DL_ConfigurationCommon_r18 = new(TDD_UL_DL_ConfigCommon)
			if err := ie.Ltm_Tdd_UL_DL_ConfigurationCommon_r18.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(earlyULSyncConfigR18ExtLtmRestrictedSetConfigR18Constraints)
			if err != nil {
				return err
			}
			ie.Ltm_RestrictedSetConfig_r18 = &v
		}
	}

	// Extension group 1:
	if seq.IsExtensionPresent(1) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "ltm-TimeAlignmentTimer-r19", Optional: true},
				{Name: "ltm-TimeAlignmentTimerTag2-r19", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Ltm_TimeAlignmentTimer_r19 = new(TimeAlignmentTimer)
			if err := ie.Ltm_TimeAlignmentTimer_r19.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.Ltm_TimeAlignmentTimerTag2_r19 = new(TimeAlignmentTimer)
			if err := ie.Ltm_TimeAlignmentTimerTag2_r19.Decode(dx); err != nil {
				return err
			}
		}
	}

	return nil
}
