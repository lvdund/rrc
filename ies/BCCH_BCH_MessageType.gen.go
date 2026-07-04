// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: BCCH-BCH-MessageType (line 14).

var bCCHBCHMessageTypeConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "mib"},
		{Name: "messageClassExtension"},
	},
}

const (
	BCCH_BCH_MessageType_Mib                   = 0
	BCCH_BCH_MessageType_MessageClassExtension = 1
)

type BCCH_BCH_MessageType struct {
	Choice                int
	Mib                   *MIB
	MessageClassExtension *struct{}
}

func (ie *BCCH_BCH_MessageType) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(bCCHBCHMessageTypeConstraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case BCCH_BCH_MessageType_Mib:
		if err := ie.Mib.Encode(e); err != nil {
			return err
		}
	case BCCH_BCH_MessageType_MessageClassExtension:
		// empty SEQUENCE alternative
	default:
		// Choice index out of root range.
		return &per.ConstraintViolationError{
			TypeName:   "BCCH-BCH-MessageType",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *BCCH_BCH_MessageType) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(bCCHBCHMessageTypeConstraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "BCCH-BCH-MessageType", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case BCCH_BCH_MessageType_Mib:
		ie.Mib = new(MIB)
		if err := ie.Mib.Decode(d); err != nil {
			return err
		}
	case BCCH_BCH_MessageType_MessageClassExtension:
		ie.MessageClassExtension = &struct{}{}
		// empty SEQUENCE alternative
	default:
		return &per.DecodeError{TypeName: "BCCH-BCH-MessageType", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
