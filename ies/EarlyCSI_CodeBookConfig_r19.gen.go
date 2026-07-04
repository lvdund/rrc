// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: EarlyCSI-CodeBookConfig-r19 (line 5718).

var earlyCSICodeBookConfigR19Constraints = per.ChoiceConstraints{
	Extensible: true,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "twoToThirtyTwoPorts-r19"},
		{Name: "moreThanThirtyTwoPorts-r19"},
	},
}

const (
	EarlyCSI_CodeBookConfig_r19_TwoToThirtyTwoPorts_r19    = 0
	EarlyCSI_CodeBookConfig_r19_MoreThanThirtyTwoPorts_r19 = 1
)

type EarlyCSI_CodeBookConfig_r19 struct {
	Choice                     int
	TwoToThirtyTwoPorts_r19    *CodebookConfig
	MoreThanThirtyTwoPorts_r19 *CodebookConfig_r19
}

func (ie *EarlyCSI_CodeBookConfig_r19) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(earlyCSICodeBookConfigR19Constraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case EarlyCSI_CodeBookConfig_r19_TwoToThirtyTwoPorts_r19:
		if err := ie.TwoToThirtyTwoPorts_r19.Encode(e); err != nil {
			return err
		}
	case EarlyCSI_CodeBookConfig_r19_MoreThanThirtyTwoPorts_r19:
		if err := ie.MoreThanThirtyTwoPorts_r19.Encode(e); err != nil {
			return err
		}
	default:
		// Extension alternative: not modeled; bail out.
		return &per.ConstraintViolationError{
			TypeName:   "EarlyCSI-CodeBookConfig-r19",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *EarlyCSI_CodeBookConfig_r19) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(earlyCSICodeBookConfigR19Constraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "EarlyCSI-CodeBookConfig-r19", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case EarlyCSI_CodeBookConfig_r19_TwoToThirtyTwoPorts_r19:
		ie.TwoToThirtyTwoPorts_r19 = new(CodebookConfig)
		if err := ie.TwoToThirtyTwoPorts_r19.Decode(d); err != nil {
			return err
		}
	case EarlyCSI_CodeBookConfig_r19_MoreThanThirtyTwoPorts_r19:
		ie.MoreThanThirtyTwoPorts_r19 = new(CodebookConfig_r19)
		if err := ie.MoreThanThirtyTwoPorts_r19.Decode(d); err != nil {
			return err
		}
	default:
		return &per.DecodeError{TypeName: "EarlyCSI-CodeBookConfig-r19", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
