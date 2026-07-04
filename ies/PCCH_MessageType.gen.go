// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PCCH-MessageType (line 115).

var pCCHMessageTypeConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "c1"},
		{Name: "messageClassExtension"},
	},
}

const (
	PCCH_MessageType_C1                    = 0
	PCCH_MessageType_MessageClassExtension = 1
)

var pCCHMessageTypeC1Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "paging"},
		{Name: "spare1"},
	},
}

const (
	PCCH_MessageType_C1_Paging = 0
	PCCH_MessageType_C1_Spare1 = 1
)

type PCCH_MessageType struct {
	Choice int
	C1     *struct {
		Choice int
		Paging *Paging
		Spare1 *struct{}
	}
	MessageClassExtension *struct{}
}

func (ie *PCCH_MessageType) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(pCCHMessageTypeConstraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case PCCH_MessageType_C1:
		choiceEnc := e.NewChoiceEncoder(pCCHMessageTypeC1Constraints)
		if err := choiceEnc.EncodeChoice(int64((*ie.C1).Choice), false, nil); err != nil {
			return err
		}
		switch (*ie.C1).Choice {
		case PCCH_MessageType_C1_Paging:
			if err := (*ie.C1).Paging.Encode(e); err != nil {
				return err
			}
		case PCCH_MessageType_C1_Spare1:
			if err := e.EncodeNull(); err != nil {
				return err
			}
		}
	case PCCH_MessageType_MessageClassExtension:
		// empty SEQUENCE alternative
	default:
		// Choice index out of root range.
		return &per.ConstraintViolationError{
			TypeName:   "PCCH-MessageType",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *PCCH_MessageType) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(pCCHMessageTypeConstraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "PCCH-MessageType", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case PCCH_MessageType_C1:
		ie.C1 = &struct {
			Choice int
			Paging *Paging
			Spare1 *struct{}
		}{}
		choiceDec := d.NewChoiceDecoder(pCCHMessageTypeC1Constraints)
		index, _, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		(*ie.C1).Choice = int(index)
		switch index {
		case PCCH_MessageType_C1_Paging:
			(*ie.C1).Paging = new(Paging)
			if err := (*ie.C1).Paging.Decode(d); err != nil {
				return err
			}
		case PCCH_MessageType_C1_Spare1:
			(*ie.C1).Spare1 = &struct{}{}
			if err := d.DecodeNull(); err != nil {
				return err
			}
		}
	case PCCH_MessageType_MessageClassExtension:
		ie.MessageClassExtension = &struct{}{}
		// empty SEQUENCE alternative
	default:
		return &per.DecodeError{TypeName: "PCCH-MessageType", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
