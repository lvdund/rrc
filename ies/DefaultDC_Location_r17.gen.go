// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: DefaultDC-Location-r17 (line 16408).

var defaultDCLocationR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "ul"},
		{Name: "dl"},
		{Name: "ulAndDL"},
	},
}

const (
	DefaultDC_Location_r17_Ul      = 0
	DefaultDC_Location_r17_Dl      = 1
	DefaultDC_Location_r17_UlAndDL = 2
)

type DefaultDC_Location_r17 struct {
	Choice  int
	Ul      *FrequencyComponent_r17
	Dl      *FrequencyComponent_r17
	UlAndDL *FrequencyComponent_r17
}

func (ie *DefaultDC_Location_r17) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(defaultDCLocationR17Constraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case DefaultDC_Location_r17_Ul:
		if err := ie.Ul.Encode(e); err != nil {
			return err
		}
	case DefaultDC_Location_r17_Dl:
		if err := ie.Dl.Encode(e); err != nil {
			return err
		}
	case DefaultDC_Location_r17_UlAndDL:
		if err := ie.UlAndDL.Encode(e); err != nil {
			return err
		}
	default:
		// Choice index out of root range.
		return &per.ConstraintViolationError{
			TypeName:   "DefaultDC-Location-r17",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *DefaultDC_Location_r17) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(defaultDCLocationR17Constraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "DefaultDC-Location-r17", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case DefaultDC_Location_r17_Ul:
		ie.Ul = new(FrequencyComponent_r17)
		if err := ie.Ul.Decode(d); err != nil {
			return err
		}
	case DefaultDC_Location_r17_Dl:
		ie.Dl = new(FrequencyComponent_r17)
		if err := ie.Dl.Decode(d); err != nil {
			return err
		}
	case DefaultDC_Location_r17_UlAndDL:
		ie.UlAndDL = new(FrequencyComponent_r17)
		if err := ie.UlAndDL.Decode(d); err != nil {
			return err
		}
	default:
		return &per.DecodeError{TypeName: "DefaultDC-Location-r17", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
