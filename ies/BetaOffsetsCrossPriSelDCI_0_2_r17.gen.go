// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: BetaOffsetsCrossPriSelDCI-0-2-r17 (line 12491).

var betaOffsetsCrossPriSelDCI02R17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "dynamicDCI-0-2-r17"},
		{Name: "semiStaticDCI-0-2-r17"},
	},
}

const (
	BetaOffsetsCrossPriSelDCI_0_2_r17_DynamicDCI_0_2_r17    = 0
	BetaOffsetsCrossPriSelDCI_0_2_r17_SemiStaticDCI_0_2_r17 = 1
)

var betaOffsetsCrossPriSelDCI02R17DynamicDCI02R17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "oneBit-r17"},
		{Name: "twoBits-r17"},
	},
}

const (
	BetaOffsetsCrossPriSelDCI_0_2_r17_DynamicDCI_0_2_r17_OneBit_r17  = 0
	BetaOffsetsCrossPriSelDCI_0_2_r17_DynamicDCI_0_2_r17_TwoBits_r17 = 1
)

var betaOffsetsCrossPriSelDCI02R17DynamicDCI02R17OneBitR17Constraints = per.FixedSize(2)

var betaOffsetsCrossPriSelDCI02R17DynamicDCI02R17TwoBitsR17Constraints = per.FixedSize(4)

type BetaOffsetsCrossPriSelDCI_0_2_r17 struct {
	Choice             int
	DynamicDCI_0_2_r17 *struct {
		Choice      int
		OneBit_r17  *[]BetaOffsetsCrossPri_r17
		TwoBits_r17 *[]BetaOffsetsCrossPri_r17
	}
	SemiStaticDCI_0_2_r17 *BetaOffsetsCrossPri_r17
}

func (ie *BetaOffsetsCrossPriSelDCI_0_2_r17) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(betaOffsetsCrossPriSelDCI02R17Constraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case BetaOffsetsCrossPriSelDCI_0_2_r17_DynamicDCI_0_2_r17:
		choiceEnc := e.NewChoiceEncoder(betaOffsetsCrossPriSelDCI02R17DynamicDCI02R17Constraints)
		if err := choiceEnc.EncodeChoice(int64((*ie.DynamicDCI_0_2_r17).Choice), false, nil); err != nil {
			return err
		}
		switch (*ie.DynamicDCI_0_2_r17).Choice {
		case BetaOffsetsCrossPriSelDCI_0_2_r17_DynamicDCI_0_2_r17_OneBit_r17:
			seqOf := e.NewSequenceOfEncoder(betaOffsetsCrossPriSelDCI02R17DynamicDCI02R17OneBitR17Constraints)
			if err := seqOf.EncodeLength(int64(len((*(*ie.DynamicDCI_0_2_r17).OneBit_r17)))); err != nil {
				return err
			}
			for i := range *(*ie.DynamicDCI_0_2_r17).OneBit_r17 {
				if err := (*(*ie.DynamicDCI_0_2_r17).OneBit_r17)[i].Encode(e); err != nil {
					return err
				}
			}
		case BetaOffsetsCrossPriSelDCI_0_2_r17_DynamicDCI_0_2_r17_TwoBits_r17:
			seqOf := e.NewSequenceOfEncoder(betaOffsetsCrossPriSelDCI02R17DynamicDCI02R17TwoBitsR17Constraints)
			if err := seqOf.EncodeLength(int64(len((*(*ie.DynamicDCI_0_2_r17).TwoBits_r17)))); err != nil {
				return err
			}
			for i := range *(*ie.DynamicDCI_0_2_r17).TwoBits_r17 {
				if err := (*(*ie.DynamicDCI_0_2_r17).TwoBits_r17)[i].Encode(e); err != nil {
					return err
				}
			}
		}
	case BetaOffsetsCrossPriSelDCI_0_2_r17_SemiStaticDCI_0_2_r17:
		if err := ie.SemiStaticDCI_0_2_r17.Encode(e); err != nil {
			return err
		}
	default:
		// Choice index out of root range.
		return &per.ConstraintViolationError{
			TypeName:   "BetaOffsetsCrossPriSelDCI-0-2-r17",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *BetaOffsetsCrossPriSelDCI_0_2_r17) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(betaOffsetsCrossPriSelDCI02R17Constraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "BetaOffsetsCrossPriSelDCI-0-2-r17", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case BetaOffsetsCrossPriSelDCI_0_2_r17_DynamicDCI_0_2_r17:
		ie.DynamicDCI_0_2_r17 = &struct {
			Choice      int
			OneBit_r17  *[]BetaOffsetsCrossPri_r17
			TwoBits_r17 *[]BetaOffsetsCrossPri_r17
		}{}
		choiceDec := d.NewChoiceDecoder(betaOffsetsCrossPriSelDCI02R17DynamicDCI02R17Constraints)
		index, _, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		(*ie.DynamicDCI_0_2_r17).Choice = int(index)
		switch index {
		case BetaOffsetsCrossPriSelDCI_0_2_r17_DynamicDCI_0_2_r17_OneBit_r17:
			(*ie.DynamicDCI_0_2_r17).OneBit_r17 = new([]BetaOffsetsCrossPri_r17)
			seqOf := d.NewSequenceOfDecoder(betaOffsetsCrossPriSelDCI02R17DynamicDCI02R17OneBitR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			(*(*ie.DynamicDCI_0_2_r17).OneBit_r17) = make([]BetaOffsetsCrossPri_r17, n)
			for i := int64(0); i < n; i++ {
				if err := (*(*ie.DynamicDCI_0_2_r17).OneBit_r17)[i].Decode(d); err != nil {
					return err
				}
			}
		case BetaOffsetsCrossPriSelDCI_0_2_r17_DynamicDCI_0_2_r17_TwoBits_r17:
			(*ie.DynamicDCI_0_2_r17).TwoBits_r17 = new([]BetaOffsetsCrossPri_r17)
			seqOf := d.NewSequenceOfDecoder(betaOffsetsCrossPriSelDCI02R17DynamicDCI02R17TwoBitsR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			(*(*ie.DynamicDCI_0_2_r17).TwoBits_r17) = make([]BetaOffsetsCrossPri_r17, n)
			for i := int64(0); i < n; i++ {
				if err := (*(*ie.DynamicDCI_0_2_r17).TwoBits_r17)[i].Decode(d); err != nil {
					return err
				}
			}
		}
	case BetaOffsetsCrossPriSelDCI_0_2_r17_SemiStaticDCI_0_2_r17:
		ie.SemiStaticDCI_0_2_r17 = new(BetaOffsetsCrossPri_r17)
		if err := ie.SemiStaticDCI_0_2_r17.Decode(d); err != nil {
			return err
		}
	default:
		return &per.DecodeError{TypeName: "BetaOffsetsCrossPriSelDCI-0-2-r17", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
