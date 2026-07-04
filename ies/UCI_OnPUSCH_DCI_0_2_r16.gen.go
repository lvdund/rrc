// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UCI-OnPUSCH-DCI-0-2-r16 (line 12463).

var uCIOnPUSCHDCI02R16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "betaOffsetsDCI-0-2-r16", Optional: true},
		{Name: "scalingDCI-0-2-r16"},
	},
}

var uCI_OnPUSCH_DCI_0_2_r16BetaOffsetsDCI02R16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "dynamicDCI-0-2-r16"},
		{Name: "semiStaticDCI-0-2-r16"},
	},
}

const (
	UCI_OnPUSCH_DCI_0_2_r16_BetaOffsetsDCI_0_2_r16_DynamicDCI_0_2_r16    = 0
	UCI_OnPUSCH_DCI_0_2_r16_BetaOffsetsDCI_0_2_r16_SemiStaticDCI_0_2_r16 = 1
)

const (
	UCI_OnPUSCH_DCI_0_2_r16_ScalingDCI_0_2_r16_F0p5  = 0
	UCI_OnPUSCH_DCI_0_2_r16_ScalingDCI_0_2_r16_F0p65 = 1
	UCI_OnPUSCH_DCI_0_2_r16_ScalingDCI_0_2_r16_F0p8  = 2
	UCI_OnPUSCH_DCI_0_2_r16_ScalingDCI_0_2_r16_F1    = 3
)

var uCIOnPUSCHDCI02R16ScalingDCI02R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UCI_OnPUSCH_DCI_0_2_r16_ScalingDCI_0_2_r16_F0p5, UCI_OnPUSCH_DCI_0_2_r16_ScalingDCI_0_2_r16_F0p65, UCI_OnPUSCH_DCI_0_2_r16_ScalingDCI_0_2_r16_F0p8, UCI_OnPUSCH_DCI_0_2_r16_ScalingDCI_0_2_r16_F1},
}

var uCIOnPUSCHDCI02R16BetaOffsetsDCI02R16DynamicDCI02R16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "oneBit-r16"},
		{Name: "twoBits-r16"},
	},
}

const (
	UCI_OnPUSCH_DCI_0_2_r16_BetaOffsetsDCI_0_2_r16_DynamicDCI_0_2_r16_OneBit_r16  = 0
	UCI_OnPUSCH_DCI_0_2_r16_BetaOffsetsDCI_0_2_r16_DynamicDCI_0_2_r16_TwoBits_r16 = 1
)

var uCIOnPUSCHDCI02R16BetaOffsetsDCI02R16DynamicDCI02R16OneBitR16Constraints = per.FixedSize(2)

var uCIOnPUSCHDCI02R16BetaOffsetsDCI02R16DynamicDCI02R16TwoBitsR16Constraints = per.FixedSize(4)

type UCI_OnPUSCH_DCI_0_2_r16 struct {
	BetaOffsetsDCI_0_2_r16 *struct {
		Choice             int
		DynamicDCI_0_2_r16 *struct {
			Choice      int
			OneBit_r16  *[]BetaOffsets
			TwoBits_r16 *[]BetaOffsets
		}
		SemiStaticDCI_0_2_r16 *BetaOffsets
	}
	ScalingDCI_0_2_r16 int64
}

func (ie *UCI_OnPUSCH_DCI_0_2_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uCIOnPUSCHDCI02R16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.BetaOffsetsDCI_0_2_r16 != nil}); err != nil {
		return err
	}

	// 2. betaOffsetsDCI-0-2-r16: choice
	{
		if ie.BetaOffsetsDCI_0_2_r16 != nil {
			choiceEnc := e.NewChoiceEncoder(uCI_OnPUSCH_DCI_0_2_r16BetaOffsetsDCI02R16Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.BetaOffsetsDCI_0_2_r16).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.BetaOffsetsDCI_0_2_r16).Choice {
			case UCI_OnPUSCH_DCI_0_2_r16_BetaOffsetsDCI_0_2_r16_DynamicDCI_0_2_r16:
				choiceEnc := e.NewChoiceEncoder(uCIOnPUSCHDCI02R16BetaOffsetsDCI02R16DynamicDCI02R16Constraints)
				if err := choiceEnc.EncodeChoice(int64((*(*ie.BetaOffsetsDCI_0_2_r16).DynamicDCI_0_2_r16).Choice), false, nil); err != nil {
					return err
				}
				switch (*(*ie.BetaOffsetsDCI_0_2_r16).DynamicDCI_0_2_r16).Choice {
				case UCI_OnPUSCH_DCI_0_2_r16_BetaOffsetsDCI_0_2_r16_DynamicDCI_0_2_r16_OneBit_r16:
					seqOf := e.NewSequenceOfEncoder(uCIOnPUSCHDCI02R16BetaOffsetsDCI02R16DynamicDCI02R16OneBitR16Constraints)
					if err := seqOf.EncodeLength(int64(len((*(*(*ie.BetaOffsetsDCI_0_2_r16).DynamicDCI_0_2_r16).OneBit_r16)))); err != nil {
						return err
					}
					for i := range *(*(*ie.BetaOffsetsDCI_0_2_r16).DynamicDCI_0_2_r16).OneBit_r16 {
						if err := (*(*(*ie.BetaOffsetsDCI_0_2_r16).DynamicDCI_0_2_r16).OneBit_r16)[i].Encode(e); err != nil {
							return err
						}
					}
				case UCI_OnPUSCH_DCI_0_2_r16_BetaOffsetsDCI_0_2_r16_DynamicDCI_0_2_r16_TwoBits_r16:
					seqOf := e.NewSequenceOfEncoder(uCIOnPUSCHDCI02R16BetaOffsetsDCI02R16DynamicDCI02R16TwoBitsR16Constraints)
					if err := seqOf.EncodeLength(int64(len((*(*(*ie.BetaOffsetsDCI_0_2_r16).DynamicDCI_0_2_r16).TwoBits_r16)))); err != nil {
						return err
					}
					for i := range *(*(*ie.BetaOffsetsDCI_0_2_r16).DynamicDCI_0_2_r16).TwoBits_r16 {
						if err := (*(*(*ie.BetaOffsetsDCI_0_2_r16).DynamicDCI_0_2_r16).TwoBits_r16)[i].Encode(e); err != nil {
							return err
						}
					}
				}
			case UCI_OnPUSCH_DCI_0_2_r16_BetaOffsetsDCI_0_2_r16_SemiStaticDCI_0_2_r16:
				if err := (*ie.BetaOffsetsDCI_0_2_r16).SemiStaticDCI_0_2_r16.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.BetaOffsetsDCI_0_2_r16).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 3. scalingDCI-0-2-r16: enumerated
	{
		if err := e.EncodeEnumerated(ie.ScalingDCI_0_2_r16, uCIOnPUSCHDCI02R16ScalingDCI02R16Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *UCI_OnPUSCH_DCI_0_2_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uCIOnPUSCHDCI02R16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. betaOffsetsDCI-0-2-r16: choice
	{
		if seq.IsComponentPresent(0) {
			ie.BetaOffsetsDCI_0_2_r16 = &struct {
				Choice             int
				DynamicDCI_0_2_r16 *struct {
					Choice      int
					OneBit_r16  *[]BetaOffsets
					TwoBits_r16 *[]BetaOffsets
				}
				SemiStaticDCI_0_2_r16 *BetaOffsets
			}{}
			choiceDec := d.NewChoiceDecoder(uCI_OnPUSCH_DCI_0_2_r16BetaOffsetsDCI02R16Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.BetaOffsetsDCI_0_2_r16).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case UCI_OnPUSCH_DCI_0_2_r16_BetaOffsetsDCI_0_2_r16_DynamicDCI_0_2_r16:
				(*ie.BetaOffsetsDCI_0_2_r16).DynamicDCI_0_2_r16 = &struct {
					Choice      int
					OneBit_r16  *[]BetaOffsets
					TwoBits_r16 *[]BetaOffsets
				}{}
				choiceDec := d.NewChoiceDecoder(uCIOnPUSCHDCI02R16BetaOffsetsDCI02R16DynamicDCI02R16Constraints)
				index, _, _, err := choiceDec.DecodeChoice()
				if err != nil {
					return err
				}
				(*(*ie.BetaOffsetsDCI_0_2_r16).DynamicDCI_0_2_r16).Choice = int(index)
				switch index {
				case UCI_OnPUSCH_DCI_0_2_r16_BetaOffsetsDCI_0_2_r16_DynamicDCI_0_2_r16_OneBit_r16:
					(*(*ie.BetaOffsetsDCI_0_2_r16).DynamicDCI_0_2_r16).OneBit_r16 = new([]BetaOffsets)
					seqOf := d.NewSequenceOfDecoder(uCIOnPUSCHDCI02R16BetaOffsetsDCI02R16DynamicDCI02R16OneBitR16Constraints)
					n, err := seqOf.DecodeLength()
					if err != nil {
						return err
					}
					(*(*(*ie.BetaOffsetsDCI_0_2_r16).DynamicDCI_0_2_r16).OneBit_r16) = make([]BetaOffsets, n)
					for i := int64(0); i < n; i++ {
						if err := (*(*(*ie.BetaOffsetsDCI_0_2_r16).DynamicDCI_0_2_r16).OneBit_r16)[i].Decode(d); err != nil {
							return err
						}
					}
				case UCI_OnPUSCH_DCI_0_2_r16_BetaOffsetsDCI_0_2_r16_DynamicDCI_0_2_r16_TwoBits_r16:
					(*(*ie.BetaOffsetsDCI_0_2_r16).DynamicDCI_0_2_r16).TwoBits_r16 = new([]BetaOffsets)
					seqOf := d.NewSequenceOfDecoder(uCIOnPUSCHDCI02R16BetaOffsetsDCI02R16DynamicDCI02R16TwoBitsR16Constraints)
					n, err := seqOf.DecodeLength()
					if err != nil {
						return err
					}
					(*(*(*ie.BetaOffsetsDCI_0_2_r16).DynamicDCI_0_2_r16).TwoBits_r16) = make([]BetaOffsets, n)
					for i := int64(0); i < n; i++ {
						if err := (*(*(*ie.BetaOffsetsDCI_0_2_r16).DynamicDCI_0_2_r16).TwoBits_r16)[i].Decode(d); err != nil {
							return err
						}
					}
				}
			case UCI_OnPUSCH_DCI_0_2_r16_BetaOffsetsDCI_0_2_r16_SemiStaticDCI_0_2_r16:
				(*ie.BetaOffsetsDCI_0_2_r16).SemiStaticDCI_0_2_r16 = new(BetaOffsets)
				if err := (*ie.BetaOffsetsDCI_0_2_r16).SemiStaticDCI_0_2_r16.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 3. scalingDCI-0-2-r16: enumerated
	{
		v1, err := d.DecodeEnumerated(uCIOnPUSCHDCI02R16ScalingDCI02R16Constraints)
		if err != nil {
			return err
		}
		ie.ScalingDCI_0_2_r16 = v1
	}

	return nil
}
