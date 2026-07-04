// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: S-NSSAI (line 15217).

var sNSSAIConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "sst"},
		{Name: "sst-SD"},
	},
}

const (
	S_NSSAI_Sst    = 0
	S_NSSAI_Sst_SD = 1
)

type S_NSSAI struct {
	Choice int
	Sst    *per.BitString
	Sst_SD *per.BitString
}

func (ie *S_NSSAI) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(sNSSAIConstraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case S_NSSAI_Sst:
		if err := e.EncodeBitString((*ie.Sst), per.FixedSize(8)); err != nil {
			return err
		}
	case S_NSSAI_Sst_SD:
		if err := e.EncodeBitString((*ie.Sst_SD), per.FixedSize(32)); err != nil {
			return err
		}
	default:
		// Choice index out of root range.
		return &per.ConstraintViolationError{
			TypeName:   "S-NSSAI",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *S_NSSAI) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(sNSSAIConstraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "S-NSSAI", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case S_NSSAI_Sst:
		ie.Sst = new(per.BitString)
		v, err := d.DecodeBitString(per.FixedSize(8))
		if err != nil {
			return err
		}
		(*ie.Sst) = v
	case S_NSSAI_Sst_SD:
		ie.Sst_SD = new(per.BitString)
		v, err := d.DecodeBitString(per.FixedSize(32))
		if err != nil {
			return err
		}
		(*ie.Sst_SD) = v
	default:
		return &per.DecodeError{TypeName: "S-NSSAI", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
