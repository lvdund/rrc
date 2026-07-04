// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CG-UCI-OnPUSCH (line 6670).

var cGUCIOnPUSCHConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "dynamic"},
		{Name: "semiStatic"},
	},
}

const (
	CG_UCI_OnPUSCH_Dynamic    = 0
	CG_UCI_OnPUSCH_SemiStatic = 1
)

var cGUCIOnPUSCHDynamicConstraints = per.SizeRange(1, 4)

type CG_UCI_OnPUSCH struct {
	Choice     int
	Dynamic    *[]BetaOffsets
	SemiStatic *BetaOffsets
}

func (ie *CG_UCI_OnPUSCH) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(cGUCIOnPUSCHConstraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case CG_UCI_OnPUSCH_Dynamic:
		seqOf := e.NewSequenceOfEncoder(cGUCIOnPUSCHDynamicConstraints)
		if err := seqOf.EncodeLength(int64(len((*ie.Dynamic)))); err != nil {
			return err
		}
		for i := range *ie.Dynamic {
			if err := (*ie.Dynamic)[i].Encode(e); err != nil {
				return err
			}
		}
	case CG_UCI_OnPUSCH_SemiStatic:
		if err := ie.SemiStatic.Encode(e); err != nil {
			return err
		}
	default:
		// Choice index out of root range.
		return &per.ConstraintViolationError{
			TypeName:   "CG-UCI-OnPUSCH",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *CG_UCI_OnPUSCH) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(cGUCIOnPUSCHConstraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "CG-UCI-OnPUSCH", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case CG_UCI_OnPUSCH_Dynamic:
		ie.Dynamic = new([]BetaOffsets)
		seqOf := d.NewSequenceOfDecoder(cGUCIOnPUSCHDynamicConstraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		(*ie.Dynamic) = make([]BetaOffsets, n)
		for i := int64(0); i < n; i++ {
			if err := (*ie.Dynamic)[i].Decode(d); err != nil {
				return err
			}
		}
	case CG_UCI_OnPUSCH_SemiStatic:
		ie.SemiStatic = new(BetaOffsets)
		if err := ie.SemiStatic.Decode(d); err != nil {
			return err
		}
	default:
		return &per.DecodeError{TypeName: "CG-UCI-OnPUSCH", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
