// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: InitialUE-Identity (line 1795).

var initialUEIdentityConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "ng-5G-S-TMSI-Part1"},
		{Name: "randomValue"},
	},
}

const (
	InitialUE_Identity_Ng_5G_S_TMSI_Part1 = 0
	InitialUE_Identity_RandomValue        = 1
)

type InitialUE_Identity struct {
	Choice             int
	Ng_5G_S_TMSI_Part1 *per.BitString
	RandomValue        *per.BitString
}

func (ie *InitialUE_Identity) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(initialUEIdentityConstraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case InitialUE_Identity_Ng_5G_S_TMSI_Part1:
		if err := e.EncodeBitString((*ie.Ng_5G_S_TMSI_Part1), per.FixedSize(39)); err != nil {
			return err
		}
	case InitialUE_Identity_RandomValue:
		if err := e.EncodeBitString((*ie.RandomValue), per.FixedSize(39)); err != nil {
			return err
		}
	default:
		// Choice index out of root range.
		return &per.ConstraintViolationError{
			TypeName:   "InitialUE-Identity",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *InitialUE_Identity) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(initialUEIdentityConstraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "InitialUE-Identity", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case InitialUE_Identity_Ng_5G_S_TMSI_Part1:
		ie.Ng_5G_S_TMSI_Part1 = new(per.BitString)
		v, err := d.DecodeBitString(per.FixedSize(39))
		if err != nil {
			return err
		}
		(*ie.Ng_5G_S_TMSI_Part1) = v
	case InitialUE_Identity_RandomValue:
		ie.RandomValue = new(per.BitString)
		v, err := d.DecodeBitString(per.FixedSize(39))
		if err != nil {
			return err
		}
		(*ie.RandomValue) = v
	default:
		return &per.DecodeError{TypeName: "InitialUE-Identity", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
