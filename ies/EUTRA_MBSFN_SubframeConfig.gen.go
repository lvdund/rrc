// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: EUTRA-MBSFN-SubframeConfig (line 26141).

var eUTRAMBSFNSubframeConfigConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "radioframeAllocationPeriod"},
		{Name: "radioframeAllocationOffset"},
		{Name: "subframeAllocation1"},
		{Name: "subframeAllocation2", Optional: true},
	},
}

const (
	EUTRA_MBSFN_SubframeConfig_RadioframeAllocationPeriod_N1  = 0
	EUTRA_MBSFN_SubframeConfig_RadioframeAllocationPeriod_N2  = 1
	EUTRA_MBSFN_SubframeConfig_RadioframeAllocationPeriod_N4  = 2
	EUTRA_MBSFN_SubframeConfig_RadioframeAllocationPeriod_N8  = 3
	EUTRA_MBSFN_SubframeConfig_RadioframeAllocationPeriod_N16 = 4
	EUTRA_MBSFN_SubframeConfig_RadioframeAllocationPeriod_N32 = 5
)

var eUTRAMBSFNSubframeConfigRadioframeAllocationPeriodConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{EUTRA_MBSFN_SubframeConfig_RadioframeAllocationPeriod_N1, EUTRA_MBSFN_SubframeConfig_RadioframeAllocationPeriod_N2, EUTRA_MBSFN_SubframeConfig_RadioframeAllocationPeriod_N4, EUTRA_MBSFN_SubframeConfig_RadioframeAllocationPeriod_N8, EUTRA_MBSFN_SubframeConfig_RadioframeAllocationPeriod_N16, EUTRA_MBSFN_SubframeConfig_RadioframeAllocationPeriod_N32},
}

var eUTRAMBSFNSubframeConfigRadioframeAllocationOffsetConstraints = per.Constrained(0, 7)

var eUTRA_MBSFN_SubframeConfigSubframeAllocation1Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "oneFrame"},
		{Name: "fourFrames"},
	},
}

const (
	EUTRA_MBSFN_SubframeConfig_SubframeAllocation1_OneFrame   = 0
	EUTRA_MBSFN_SubframeConfig_SubframeAllocation1_FourFrames = 1
)

var eUTRA_MBSFN_SubframeConfigSubframeAllocation2Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "oneFrame"},
		{Name: "fourFrames"},
	},
}

const (
	EUTRA_MBSFN_SubframeConfig_SubframeAllocation2_OneFrame   = 0
	EUTRA_MBSFN_SubframeConfig_SubframeAllocation2_FourFrames = 1
)

type EUTRA_MBSFN_SubframeConfig struct {
	RadioframeAllocationPeriod int64
	RadioframeAllocationOffset int64
	SubframeAllocation1        struct {
		Choice     int
		OneFrame   *per.BitString
		FourFrames *per.BitString
	}
	SubframeAllocation2 *struct {
		Choice     int
		OneFrame   *per.BitString
		FourFrames *per.BitString
	}
}

func (ie *EUTRA_MBSFN_SubframeConfig) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(eUTRAMBSFNSubframeConfigConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.SubframeAllocation2 != nil}); err != nil {
		return err
	}

	// 3. radioframeAllocationPeriod: enumerated
	{
		if err := e.EncodeEnumerated(ie.RadioframeAllocationPeriod, eUTRAMBSFNSubframeConfigRadioframeAllocationPeriodConstraints); err != nil {
			return err
		}
	}

	// 4. radioframeAllocationOffset: integer
	{
		if err := e.EncodeInteger(ie.RadioframeAllocationOffset, eUTRAMBSFNSubframeConfigRadioframeAllocationOffsetConstraints); err != nil {
			return err
		}
	}

	// 5. subframeAllocation1: choice
	{
		choiceEnc := e.NewChoiceEncoder(eUTRA_MBSFN_SubframeConfigSubframeAllocation1Constraints)
		if err := choiceEnc.EncodeChoice(int64(ie.SubframeAllocation1.Choice), false, nil); err != nil {
			return err
		}
		switch ie.SubframeAllocation1.Choice {
		case EUTRA_MBSFN_SubframeConfig_SubframeAllocation1_OneFrame:
			if err := e.EncodeBitString((*ie.SubframeAllocation1.OneFrame), per.FixedSize(6)); err != nil {
				return err
			}
		case EUTRA_MBSFN_SubframeConfig_SubframeAllocation1_FourFrames:
			if err := e.EncodeBitString((*ie.SubframeAllocation1.FourFrames), per.FixedSize(24)); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.SubframeAllocation1.Choice), Constraint: "index out of range"}
		}
	}

	// 6. subframeAllocation2: choice
	{
		if ie.SubframeAllocation2 != nil {
			choiceEnc := e.NewChoiceEncoder(eUTRA_MBSFN_SubframeConfigSubframeAllocation2Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.SubframeAllocation2).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.SubframeAllocation2).Choice {
			case EUTRA_MBSFN_SubframeConfig_SubframeAllocation2_OneFrame:
				if err := e.EncodeBitString((*(*ie.SubframeAllocation2).OneFrame), per.FixedSize(2)); err != nil {
					return err
				}
			case EUTRA_MBSFN_SubframeConfig_SubframeAllocation2_FourFrames:
				if err := e.EncodeBitString((*(*ie.SubframeAllocation2).FourFrames), per.FixedSize(8)); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.SubframeAllocation2).Choice), Constraint: "index out of range"}
			}
		}
	}

	return nil
}

func (ie *EUTRA_MBSFN_SubframeConfig) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(eUTRAMBSFNSubframeConfigConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. radioframeAllocationPeriod: enumerated
	{
		v0, err := d.DecodeEnumerated(eUTRAMBSFNSubframeConfigRadioframeAllocationPeriodConstraints)
		if err != nil {
			return err
		}
		ie.RadioframeAllocationPeriod = v0
	}

	// 4. radioframeAllocationOffset: integer
	{
		v1, err := d.DecodeInteger(eUTRAMBSFNSubframeConfigRadioframeAllocationOffsetConstraints)
		if err != nil {
			return err
		}
		ie.RadioframeAllocationOffset = v1
	}

	// 5. subframeAllocation1: choice
	{
		choiceDec := d.NewChoiceDecoder(eUTRA_MBSFN_SubframeConfigSubframeAllocation1Constraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.SubframeAllocation1.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case EUTRA_MBSFN_SubframeConfig_SubframeAllocation1_OneFrame:
			ie.SubframeAllocation1.OneFrame = new(per.BitString)
			v, err := d.DecodeBitString(per.FixedSize(6))
			if err != nil {
				return err
			}
			(*ie.SubframeAllocation1.OneFrame) = v
		case EUTRA_MBSFN_SubframeConfig_SubframeAllocation1_FourFrames:
			ie.SubframeAllocation1.FourFrames = new(per.BitString)
			v, err := d.DecodeBitString(per.FixedSize(24))
			if err != nil {
				return err
			}
			(*ie.SubframeAllocation1.FourFrames) = v
		}
	}

	// 6. subframeAllocation2: choice
	{
		if seq.IsComponentPresent(3) {
			ie.SubframeAllocation2 = &struct {
				Choice     int
				OneFrame   *per.BitString
				FourFrames *per.BitString
			}{}
			choiceDec := d.NewChoiceDecoder(eUTRA_MBSFN_SubframeConfigSubframeAllocation2Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.SubframeAllocation2).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case EUTRA_MBSFN_SubframeConfig_SubframeAllocation2_OneFrame:
				(*ie.SubframeAllocation2).OneFrame = new(per.BitString)
				v, err := d.DecodeBitString(per.FixedSize(2))
				if err != nil {
					return err
				}
				(*(*ie.SubframeAllocation2).OneFrame) = v
			case EUTRA_MBSFN_SubframeConfig_SubframeAllocation2_FourFrames:
				(*ie.SubframeAllocation2).FourFrames = new(per.BitString)
				v, err := d.DecodeBitString(per.FixedSize(8))
				if err != nil {
					return err
				}
				(*(*ie.SubframeAllocation2).FourFrames) = v
			}
		}
	}

	return nil
}
