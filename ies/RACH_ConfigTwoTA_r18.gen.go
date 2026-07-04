// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RACH-ConfigTwoTA-r18 (line 13040).

var rACHConfigTwoTAR18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "additionalPCI-andRACH-Index-r18"},
		{Name: "rach-ConfigGeneric-r18"},
		{Name: "ssb-perRACH-Occasion-r18", Optional: true},
		{Name: "prach-RootSequenceIndex-r18"},
		{Name: "msg1-SubcarrierSpacing-r18", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

const (
	RACH_ConfigTwoTA_r18_Ssb_PerRACH_Occasion_r18_OneEighth = 0
	RACH_ConfigTwoTA_r18_Ssb_PerRACH_Occasion_r18_OneFourth = 1
	RACH_ConfigTwoTA_r18_Ssb_PerRACH_Occasion_r18_OneHalf   = 2
	RACH_ConfigTwoTA_r18_Ssb_PerRACH_Occasion_r18_One       = 3
	RACH_ConfigTwoTA_r18_Ssb_PerRACH_Occasion_r18_Two       = 4
	RACH_ConfigTwoTA_r18_Ssb_PerRACH_Occasion_r18_Four      = 5
	RACH_ConfigTwoTA_r18_Ssb_PerRACH_Occasion_r18_Eight     = 6
	RACH_ConfigTwoTA_r18_Ssb_PerRACH_Occasion_r18_Sixteen   = 7
)

var rACHConfigTwoTAR18SsbPerRACHOccasionR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RACH_ConfigTwoTA_r18_Ssb_PerRACH_Occasion_r18_OneEighth, RACH_ConfigTwoTA_r18_Ssb_PerRACH_Occasion_r18_OneFourth, RACH_ConfigTwoTA_r18_Ssb_PerRACH_Occasion_r18_OneHalf, RACH_ConfigTwoTA_r18_Ssb_PerRACH_Occasion_r18_One, RACH_ConfigTwoTA_r18_Ssb_PerRACH_Occasion_r18_Two, RACH_ConfigTwoTA_r18_Ssb_PerRACH_Occasion_r18_Four, RACH_ConfigTwoTA_r18_Ssb_PerRACH_Occasion_r18_Eight, RACH_ConfigTwoTA_r18_Ssb_PerRACH_Occasion_r18_Sixteen},
}

var rACH_ConfigTwoTA_r18PrachRootSequenceIndexR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "l839"},
		{Name: "l139"},
		{Name: "l571"},
		{Name: "l1151"},
	},
}

const (
	RACH_ConfigTwoTA_r18_Prach_RootSequenceIndex_r18_L839  = 0
	RACH_ConfigTwoTA_r18_Prach_RootSequenceIndex_r18_L139  = 1
	RACH_ConfigTwoTA_r18_Prach_RootSequenceIndex_r18_L571  = 2
	RACH_ConfigTwoTA_r18_Prach_RootSequenceIndex_r18_L1151 = 3
)

const (
	RACH_ConfigTwoTA_r18_Ext_TwoTA_RestrictedSetConfig_r18_UnrestrictedSet    = 0
	RACH_ConfigTwoTA_r18_Ext_TwoTA_RestrictedSetConfig_r18_RestrictedSetTypeA = 1
	RACH_ConfigTwoTA_r18_Ext_TwoTA_RestrictedSetConfig_r18_RestrictedSetTypeB = 2
)

var rACHConfigTwoTAR18ExtTwoTARestrictedSetConfigR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RACH_ConfigTwoTA_r18_Ext_TwoTA_RestrictedSetConfig_r18_UnrestrictedSet, RACH_ConfigTwoTA_r18_Ext_TwoTA_RestrictedSetConfig_r18_RestrictedSetTypeA, RACH_ConfigTwoTA_r18_Ext_TwoTA_RestrictedSetConfig_r18_RestrictedSetTypeB},
}

type RACH_ConfigTwoTA_r18 struct {
	AdditionalPCI_AndRACH_Index_r18 AdditionalPCIIndex_r17
	Rach_ConfigGeneric_r18          RACH_ConfigGeneric
	Ssb_PerRACH_Occasion_r18        *int64
	Prach_RootSequenceIndex_r18     struct {
		Choice int
		L839   *int64
		L139   *int64
		L571   *int64
		L1151  *int64
	}
	Msg1_SubcarrierSpacing_r18    *SubcarrierSpacing
	TwoTA_RestrictedSetConfig_r18 *int64
}

func (ie *RACH_ConfigTwoTA_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rACHConfigTwoTAR18Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.TwoTA_RestrictedSetConfig_r18 != nil
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ssb_PerRACH_Occasion_r18 != nil, ie.Msg1_SubcarrierSpacing_r18 != nil}); err != nil {
		return err
	}

	// 3. additionalPCI-andRACH-Index-r18: ref
	{
		if err := ie.AdditionalPCI_AndRACH_Index_r18.Encode(e); err != nil {
			return err
		}
	}

	// 4. rach-ConfigGeneric-r18: ref
	{
		if err := ie.Rach_ConfigGeneric_r18.Encode(e); err != nil {
			return err
		}
	}

	// 5. ssb-perRACH-Occasion-r18: enumerated
	{
		if ie.Ssb_PerRACH_Occasion_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Ssb_PerRACH_Occasion_r18, rACHConfigTwoTAR18SsbPerRACHOccasionR18Constraints); err != nil {
				return err
			}
		}
	}

	// 6. prach-RootSequenceIndex-r18: choice
	{
		choiceEnc := e.NewChoiceEncoder(rACH_ConfigTwoTA_r18PrachRootSequenceIndexR18Constraints)
		if err := choiceEnc.EncodeChoice(int64(ie.Prach_RootSequenceIndex_r18.Choice), false, nil); err != nil {
			return err
		}
		switch ie.Prach_RootSequenceIndex_r18.Choice {
		case RACH_ConfigTwoTA_r18_Prach_RootSequenceIndex_r18_L839:
			if err := e.EncodeInteger((*ie.Prach_RootSequenceIndex_r18.L839), per.Constrained(0, 837)); err != nil {
				return err
			}
		case RACH_ConfigTwoTA_r18_Prach_RootSequenceIndex_r18_L139:
			if err := e.EncodeInteger((*ie.Prach_RootSequenceIndex_r18.L139), per.Constrained(0, 137)); err != nil {
				return err
			}
		case RACH_ConfigTwoTA_r18_Prach_RootSequenceIndex_r18_L571:
			if err := e.EncodeInteger((*ie.Prach_RootSequenceIndex_r18.L571), per.Constrained(0, 569)); err != nil {
				return err
			}
		case RACH_ConfigTwoTA_r18_Prach_RootSequenceIndex_r18_L1151:
			if err := e.EncodeInteger((*ie.Prach_RootSequenceIndex_r18.L1151), per.Constrained(0, 1149)); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.Prach_RootSequenceIndex_r18.Choice), Constraint: "index out of range"}
		}
	}

	// 7. msg1-SubcarrierSpacing-r18: ref
	{
		if ie.Msg1_SubcarrierSpacing_r18 != nil {
			if err := ie.Msg1_SubcarrierSpacing_r18.Encode(e); err != nil {
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
					{Name: "twoTA-restrictedSetConfig-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.TwoTA_RestrictedSetConfig_r18 != nil}); err != nil {
				return err
			}

			if ie.TwoTA_RestrictedSetConfig_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.TwoTA_RestrictedSetConfig_r18, rACHConfigTwoTAR18ExtTwoTARestrictedSetConfigR18Constraints); err != nil {
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

func (ie *RACH_ConfigTwoTA_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rACHConfigTwoTAR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. additionalPCI-andRACH-Index-r18: ref
	{
		if err := ie.AdditionalPCI_AndRACH_Index_r18.Decode(d); err != nil {
			return err
		}
	}

	// 4. rach-ConfigGeneric-r18: ref
	{
		if err := ie.Rach_ConfigGeneric_r18.Decode(d); err != nil {
			return err
		}
	}

	// 5. ssb-perRACH-Occasion-r18: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(rACHConfigTwoTAR18SsbPerRACHOccasionR18Constraints)
			if err != nil {
				return err
			}
			ie.Ssb_PerRACH_Occasion_r18 = &idx
		}
	}

	// 6. prach-RootSequenceIndex-r18: choice
	{
		choiceDec := d.NewChoiceDecoder(rACH_ConfigTwoTA_r18PrachRootSequenceIndexR18Constraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.Prach_RootSequenceIndex_r18.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case RACH_ConfigTwoTA_r18_Prach_RootSequenceIndex_r18_L839:
			ie.Prach_RootSequenceIndex_r18.L839 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 837))
			if err != nil {
				return err
			}
			(*ie.Prach_RootSequenceIndex_r18.L839) = v
		case RACH_ConfigTwoTA_r18_Prach_RootSequenceIndex_r18_L139:
			ie.Prach_RootSequenceIndex_r18.L139 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 137))
			if err != nil {
				return err
			}
			(*ie.Prach_RootSequenceIndex_r18.L139) = v
		case RACH_ConfigTwoTA_r18_Prach_RootSequenceIndex_r18_L571:
			ie.Prach_RootSequenceIndex_r18.L571 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 569))
			if err != nil {
				return err
			}
			(*ie.Prach_RootSequenceIndex_r18.L571) = v
		case RACH_ConfigTwoTA_r18_Prach_RootSequenceIndex_r18_L1151:
			ie.Prach_RootSequenceIndex_r18.L1151 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 1149))
			if err != nil {
				return err
			}
			(*ie.Prach_RootSequenceIndex_r18.L1151) = v
		}
	}

	// 7. msg1-SubcarrierSpacing-r18: ref
	{
		if seq.IsComponentPresent(4) {
			ie.Msg1_SubcarrierSpacing_r18 = new(SubcarrierSpacing)
			if err := ie.Msg1_SubcarrierSpacing_r18.Decode(d); err != nil {
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
				{Name: "twoTA-restrictedSetConfig-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(rACHConfigTwoTAR18ExtTwoTARestrictedSetConfigR18Constraints)
			if err != nil {
				return err
			}
			ie.TwoTA_RestrictedSetConfig_r18 = &v
		}
	}

	return nil
}
