// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RepetitionSchemeConfig-v1630 (line 13348).

var repetitionSchemeConfigV1630Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "slotBased-v1630"},
	},
}

var repetitionSchemeConfig_v1630SlotBasedV1630Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	RepetitionSchemeConfig_v1630_SlotBased_v1630_Release = 0
	RepetitionSchemeConfig_v1630_SlotBased_v1630_Setup   = 1
)

type RepetitionSchemeConfig_v1630 struct {
	SlotBased_v1630 struct {
		Choice  int
		Release *struct{}
		Setup   *SlotBased_v1630
	}
}

func (ie *RepetitionSchemeConfig_v1630) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(repetitionSchemeConfigV1630Constraints)
	_ = seq

	// 1. slotBased-v1630: choice
	{
		choiceEnc := e.NewChoiceEncoder(repetitionSchemeConfig_v1630SlotBasedV1630Constraints)
		if err := choiceEnc.EncodeChoice(int64(ie.SlotBased_v1630.Choice), false, nil); err != nil {
			return err
		}
		switch ie.SlotBased_v1630.Choice {
		case RepetitionSchemeConfig_v1630_SlotBased_v1630_Release:
			if err := e.EncodeNull(); err != nil {
				return err
			}
		case RepetitionSchemeConfig_v1630_SlotBased_v1630_Setup:
			if err := ie.SlotBased_v1630.Setup.Encode(e); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.SlotBased_v1630.Choice), Constraint: "index out of range"}
		}
	}

	return nil
}

func (ie *RepetitionSchemeConfig_v1630) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(repetitionSchemeConfigV1630Constraints)
	_ = seq

	// 1. slotBased-v1630: choice
	{
		choiceDec := d.NewChoiceDecoder(repetitionSchemeConfig_v1630SlotBasedV1630Constraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.SlotBased_v1630.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case RepetitionSchemeConfig_v1630_SlotBased_v1630_Release:
			ie.SlotBased_v1630.Release = &struct{}{}
			if err := d.DecodeNull(); err != nil {
				return err
			}
		case RepetitionSchemeConfig_v1630_SlotBased_v1630_Setup:
			ie.SlotBased_v1630.Setup = new(SlotBased_v1630)
			if err := ie.SlotBased_v1630.Setup.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
