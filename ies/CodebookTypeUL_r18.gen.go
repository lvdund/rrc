// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CodebookTypeUL-r18 (line 12523).

var codebookTypeULR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "codebook1-r18"},
		{Name: "codebook2-r18"},
		{Name: "codebook3-r18"},
		{Name: "codebook4-r18"},
	},
}

const (
	CodebookTypeUL_r18_Codebook1_r18 = 0
	CodebookTypeUL_r18_Codebook2_r18 = 1
	CodebookTypeUL_r18_Codebook3_r18 = 2
	CodebookTypeUL_r18_Codebook4_r18 = 3
)

const (
	CodebookTypeUL_r18_Codebook1_r18_Ng1n4n1 = 0
	CodebookTypeUL_r18_Codebook1_r18_Ng1n2n2 = 1
)

var codebookTypeULR18Codebook1R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookTypeUL_r18_Codebook1_r18_Ng1n4n1, CodebookTypeUL_r18_Codebook1_r18_Ng1n2n2},
}

const (
	CodebookTypeUL_r18_Codebook2_r18_Ng2 = 0
)

var codebookTypeULR18Codebook2R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookTypeUL_r18_Codebook2_r18_Ng2},
}

const (
	CodebookTypeUL_r18_Codebook3_r18_Ng4 = 0
)

var codebookTypeULR18Codebook3R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookTypeUL_r18_Codebook3_r18_Ng4},
}

const (
	CodebookTypeUL_r18_Codebook4_r18_Ng8 = 0
)

var codebookTypeULR18Codebook4R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookTypeUL_r18_Codebook4_r18_Ng8},
}

type CodebookTypeUL_r18 struct {
	Choice        int
	Codebook1_r18 *int64
	Codebook2_r18 *int64
	Codebook3_r18 *int64
	Codebook4_r18 *int64
}

func (ie *CodebookTypeUL_r18) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(codebookTypeULR18Constraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case CodebookTypeUL_r18_Codebook1_r18:
		if err := e.EncodeEnumerated((*ie.Codebook1_r18), codebookTypeULR18Codebook1R18Constraints); err != nil {
			return err
		}
	case CodebookTypeUL_r18_Codebook2_r18:
		if err := e.EncodeEnumerated((*ie.Codebook2_r18), codebookTypeULR18Codebook2R18Constraints); err != nil {
			return err
		}
	case CodebookTypeUL_r18_Codebook3_r18:
		if err := e.EncodeEnumerated((*ie.Codebook3_r18), codebookTypeULR18Codebook3R18Constraints); err != nil {
			return err
		}
	case CodebookTypeUL_r18_Codebook4_r18:
		if err := e.EncodeEnumerated((*ie.Codebook4_r18), codebookTypeULR18Codebook4R18Constraints); err != nil {
			return err
		}
	default:
		// Choice index out of root range.
		return &per.ConstraintViolationError{
			TypeName:   "CodebookTypeUL-r18",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *CodebookTypeUL_r18) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(codebookTypeULR18Constraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "CodebookTypeUL-r18", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case CodebookTypeUL_r18_Codebook1_r18:
		ie.Codebook1_r18 = new(int64)
		v, err := d.DecodeEnumerated(codebookTypeULR18Codebook1R18Constraints)
		if err != nil {
			return err
		}
		(*ie.Codebook1_r18) = v
	case CodebookTypeUL_r18_Codebook2_r18:
		ie.Codebook2_r18 = new(int64)
		v, err := d.DecodeEnumerated(codebookTypeULR18Codebook2R18Constraints)
		if err != nil {
			return err
		}
		(*ie.Codebook2_r18) = v
	case CodebookTypeUL_r18_Codebook3_r18:
		ie.Codebook3_r18 = new(int64)
		v, err := d.DecodeEnumerated(codebookTypeULR18Codebook3R18Constraints)
		if err != nil {
			return err
		}
		(*ie.Codebook3_r18) = v
	case CodebookTypeUL_r18_Codebook4_r18:
		ie.Codebook4_r18 = new(int64)
		v, err := d.DecodeEnumerated(codebookTypeULR18Codebook4R18Constraints)
		if err != nil {
			return err
		}
		(*ie.Codebook4_r18) = v
	default:
		return &per.DecodeError{TypeName: "CodebookTypeUL-r18", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
