// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CodebookConfig-v1730 (line 6222).

var codebookConfigV1730Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "codebookType"},
	},
}

var codebookConfig_v1730CodebookTypeConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "type1"},
	},
}

const (
	CodebookConfig_v1730_CodebookType_Type1 = 0
)

var codebookConfigV1730CodebookTypeType1Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "codebookMode", Optional: true},
	},
}

type CodebookConfig_v1730 struct {
	CodebookType struct {
		Choice int
		Type1  *struct{ CodebookMode *int64 }
	}
}

func (ie *CodebookConfig_v1730) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(codebookConfigV1730Constraints)
	_ = seq

	// 1. codebookType: choice
	{
		choiceEnc := e.NewChoiceEncoder(codebookConfig_v1730CodebookTypeConstraints)
		if err := choiceEnc.EncodeChoice(int64(ie.CodebookType.Choice), false, nil); err != nil {
			return err
		}
		switch ie.CodebookType.Choice {
		case CodebookConfig_v1730_CodebookType_Type1:
			c := (*ie.CodebookType.Type1)
			codebookConfigV1730CodebookTypeType1Seq := e.NewSequenceEncoder(codebookConfigV1730CodebookTypeType1Constraints)
			if err := codebookConfigV1730CodebookTypeType1Seq.EncodePreamble([]bool{c.CodebookMode != nil}); err != nil {
				return err
			}
			if c.CodebookMode != nil {
				if err := e.EncodeInteger((*c.CodebookMode), per.Constrained(1, 2)); err != nil {
					return err
				}
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.CodebookType.Choice), Constraint: "index out of range"}
		}
	}

	return nil
}

func (ie *CodebookConfig_v1730) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(codebookConfigV1730Constraints)
	_ = seq

	// 1. codebookType: choice
	{
		choiceDec := d.NewChoiceDecoder(codebookConfig_v1730CodebookTypeConstraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.CodebookType.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case CodebookConfig_v1730_CodebookType_Type1:
			ie.CodebookType.Type1 = &struct{ CodebookMode *int64 }{}
			c := (*ie.CodebookType.Type1)
			codebookConfigV1730CodebookTypeType1Seq := d.NewSequenceDecoder(codebookConfigV1730CodebookTypeType1Constraints)
			if err := codebookConfigV1730CodebookTypeType1Seq.DecodePreamble(); err != nil {
				return err
			}
			if codebookConfigV1730CodebookTypeType1Seq.IsComponentPresent(0) {
				c.CodebookMode = new(int64)
				v, err := d.DecodeInteger(per.Constrained(1, 2))
				if err != nil {
					return err
				}
				(*c.CodebookMode) = v
			}
		}
	}

	return nil
}
