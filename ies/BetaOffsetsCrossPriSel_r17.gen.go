// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: BetaOffsetsCrossPriSel-r17 (line 12486).

var betaOffsetsCrossPriSelR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "dynamic-r17"},
		{Name: "semiStatic-r17"},
	},
}

const (
	BetaOffsetsCrossPriSel_r17_Dynamic_r17    = 0
	BetaOffsetsCrossPriSel_r17_SemiStatic_r17 = 1
)

var betaOffsetsCrossPriSelR17DynamicR17Constraints = per.FixedSize(4)

type BetaOffsetsCrossPriSel_r17 struct {
	Choice         int
	Dynamic_r17    *[]BetaOffsetsCrossPri_r17
	SemiStatic_r17 *BetaOffsetsCrossPri_r17
}

func (ie *BetaOffsetsCrossPriSel_r17) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(betaOffsetsCrossPriSelR17Constraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case BetaOffsetsCrossPriSel_r17_Dynamic_r17:
		seqOf := e.NewSequenceOfEncoder(betaOffsetsCrossPriSelR17DynamicR17Constraints)
		if err := seqOf.EncodeLength(int64(len((*ie.Dynamic_r17)))); err != nil {
			return err
		}
		for i := range *ie.Dynamic_r17 {
			if err := (*ie.Dynamic_r17)[i].Encode(e); err != nil {
				return err
			}
		}
	case BetaOffsetsCrossPriSel_r17_SemiStatic_r17:
		if err := ie.SemiStatic_r17.Encode(e); err != nil {
			return err
		}
	default:
		// Choice index out of root range.
		return &per.ConstraintViolationError{
			TypeName:   "BetaOffsetsCrossPriSel-r17",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *BetaOffsetsCrossPriSel_r17) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(betaOffsetsCrossPriSelR17Constraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "BetaOffsetsCrossPriSel-r17", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case BetaOffsetsCrossPriSel_r17_Dynamic_r17:
		ie.Dynamic_r17 = new([]BetaOffsetsCrossPri_r17)
		seqOf := d.NewSequenceOfDecoder(betaOffsetsCrossPriSelR17DynamicR17Constraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		(*ie.Dynamic_r17) = make([]BetaOffsetsCrossPri_r17, n)
		for i := int64(0); i < n; i++ {
			if err := (*ie.Dynamic_r17)[i].Decode(d); err != nil {
				return err
			}
		}
	case BetaOffsetsCrossPriSel_r17_SemiStatic_r17:
		ie.SemiStatic_r17 = new(BetaOffsetsCrossPri_r17)
		if err := ie.SemiStatic_r17.Decode(d); err != nil {
			return err
		}
	default:
		return &per.DecodeError{TypeName: "BetaOffsetsCrossPriSel-r17", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
