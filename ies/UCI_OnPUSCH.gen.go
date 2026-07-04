// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UCI-OnPUSCH (line 12451).

var uCIOnPUSCHConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "betaOffsets", Optional: true},
		{Name: "scaling"},
	},
}

var uCI_OnPUSCHBetaOffsetsConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "dynamic"},
		{Name: "semiStatic"},
	},
}

const (
	UCI_OnPUSCH_BetaOffsets_Dynamic    = 0
	UCI_OnPUSCH_BetaOffsets_SemiStatic = 1
)

const (
	UCI_OnPUSCH_Scaling_F0p5  = 0
	UCI_OnPUSCH_Scaling_F0p65 = 1
	UCI_OnPUSCH_Scaling_F0p8  = 2
	UCI_OnPUSCH_Scaling_F1    = 3
)

var uCIOnPUSCHScalingConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UCI_OnPUSCH_Scaling_F0p5, UCI_OnPUSCH_Scaling_F0p65, UCI_OnPUSCH_Scaling_F0p8, UCI_OnPUSCH_Scaling_F1},
}

var uCIOnPUSCHBetaOffsetsDynamicConstraints = per.FixedSize(4)

type UCI_OnPUSCH struct {
	BetaOffsets *struct {
		Choice     int
		Dynamic    *[]BetaOffsets
		SemiStatic *BetaOffsets
	}
	Scaling int64
}

func (ie *UCI_OnPUSCH) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uCIOnPUSCHConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.BetaOffsets != nil}); err != nil {
		return err
	}

	// 2. betaOffsets: choice
	{
		if ie.BetaOffsets != nil {
			choiceEnc := e.NewChoiceEncoder(uCI_OnPUSCHBetaOffsetsConstraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.BetaOffsets).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.BetaOffsets).Choice {
			case UCI_OnPUSCH_BetaOffsets_Dynamic:
				seqOf := e.NewSequenceOfEncoder(uCIOnPUSCHBetaOffsetsDynamicConstraints)
				if err := seqOf.EncodeLength(int64(len((*(*ie.BetaOffsets).Dynamic)))); err != nil {
					return err
				}
				for i := range *(*ie.BetaOffsets).Dynamic {
					if err := (*(*ie.BetaOffsets).Dynamic)[i].Encode(e); err != nil {
						return err
					}
				}
			case UCI_OnPUSCH_BetaOffsets_SemiStatic:
				if err := (*ie.BetaOffsets).SemiStatic.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.BetaOffsets).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 3. scaling: enumerated
	{
		if err := e.EncodeEnumerated(ie.Scaling, uCIOnPUSCHScalingConstraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *UCI_OnPUSCH) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uCIOnPUSCHConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. betaOffsets: choice
	{
		if seq.IsComponentPresent(0) {
			ie.BetaOffsets = &struct {
				Choice     int
				Dynamic    *[]BetaOffsets
				SemiStatic *BetaOffsets
			}{}
			choiceDec := d.NewChoiceDecoder(uCI_OnPUSCHBetaOffsetsConstraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.BetaOffsets).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case UCI_OnPUSCH_BetaOffsets_Dynamic:
				(*ie.BetaOffsets).Dynamic = new([]BetaOffsets)
				seqOf := d.NewSequenceOfDecoder(uCIOnPUSCHBetaOffsetsDynamicConstraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				(*(*ie.BetaOffsets).Dynamic) = make([]BetaOffsets, n)
				for i := int64(0); i < n; i++ {
					if err := (*(*ie.BetaOffsets).Dynamic)[i].Decode(d); err != nil {
						return err
					}
				}
			case UCI_OnPUSCH_BetaOffsets_SemiStatic:
				(*ie.BetaOffsets).SemiStatic = new(BetaOffsets)
				if err := (*ie.BetaOffsets).SemiStatic.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 3. scaling: enumerated
	{
		v1, err := d.DecodeEnumerated(uCIOnPUSCHScalingConstraints)
		if err != nil {
			return err
		}
		ie.Scaling = v1
	}

	return nil
}
