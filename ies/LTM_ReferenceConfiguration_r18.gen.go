// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: LTM-ReferenceConfiguration-r18 (line 8747).
// LTM-ReferenceConfiguration-r18 ::= CHOICE {release NULL, setup ReferenceConfiguration-r18}

var lTMReferenceConfigurationR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	LTM_ReferenceConfiguration_r18_Release = 0
	LTM_ReferenceConfiguration_r18_Setup   = 1
)

type LTM_ReferenceConfiguration_r18 struct {
	Choice  int
	Release *struct{}
	Setup   *ReferenceConfiguration_r18
}

func (ie *LTM_ReferenceConfiguration_r18) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(lTMReferenceConfigurationR18Constraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case LTM_ReferenceConfiguration_r18_Release:
		if err := e.EncodeNull(); err != nil {
			return err
		}
	case LTM_ReferenceConfiguration_r18_Setup:
		if err := ie.Setup.Encode(e); err != nil {
			return err
		}
	default:
		// Choice index out of root range.
		return &per.ConstraintViolationError{
			TypeName:   "LTM-ReferenceConfiguration-r18",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *LTM_ReferenceConfiguration_r18) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(lTMReferenceConfigurationR18Constraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "LTM-ReferenceConfiguration-r18", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case LTM_ReferenceConfiguration_r18_Release:
		ie.Release = &struct{}{}
		if err := d.DecodeNull(); err != nil {
			return err
		}
	case LTM_ReferenceConfiguration_r18_Setup:
		ie.Setup = new(ReferenceConfiguration_r18)
		if err := ie.Setup.Decode(d); err != nil {
			return err
		}
	default:
		return &per.DecodeError{TypeName: "LTM-ReferenceConfiguration-r18", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
