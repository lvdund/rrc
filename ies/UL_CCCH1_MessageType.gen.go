// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UL-CCCH1-MessageType (line 149).

var uLCCCH1MessageTypeConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "c1"},
		{Name: "messageClassExtension"},
	},
}

const (
	UL_CCCH1_MessageType_C1                    = 0
	UL_CCCH1_MessageType_MessageClassExtension = 1
)

var uLCCCH1MessageTypeC1Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "rrcResumeRequest1"},
		{Name: "spare3"},
		{Name: "spare2"},
		{Name: "spare1"},
	},
}

const (
	UL_CCCH1_MessageType_C1_RrcResumeRequest1 = 0
	UL_CCCH1_MessageType_C1_Spare3            = 1
	UL_CCCH1_MessageType_C1_Spare2            = 2
	UL_CCCH1_MessageType_C1_Spare1            = 3
)

type UL_CCCH1_MessageType struct {
	Choice int
	C1     *struct {
		Choice            int
		RrcResumeRequest1 *RRCResumeRequest1
		Spare3            *struct{}
		Spare2            *struct{}
		Spare1            *struct{}
	}
	MessageClassExtension *struct{}
}

func (ie *UL_CCCH1_MessageType) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(uLCCCH1MessageTypeConstraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case UL_CCCH1_MessageType_C1:
		choiceEnc := e.NewChoiceEncoder(uLCCCH1MessageTypeC1Constraints)
		if err := choiceEnc.EncodeChoice(int64((*ie.C1).Choice), false, nil); err != nil {
			return err
		}
		switch (*ie.C1).Choice {
		case UL_CCCH1_MessageType_C1_RrcResumeRequest1:
			if err := (*ie.C1).RrcResumeRequest1.Encode(e); err != nil {
				return err
			}
		case UL_CCCH1_MessageType_C1_Spare3:
			if err := e.EncodeNull(); err != nil {
				return err
			}
		case UL_CCCH1_MessageType_C1_Spare2:
			if err := e.EncodeNull(); err != nil {
				return err
			}
		case UL_CCCH1_MessageType_C1_Spare1:
			if err := e.EncodeNull(); err != nil {
				return err
			}
		}
	case UL_CCCH1_MessageType_MessageClassExtension:
		// empty SEQUENCE alternative
	default:
		// Choice index out of root range.
		return &per.ConstraintViolationError{
			TypeName:   "UL-CCCH1-MessageType",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *UL_CCCH1_MessageType) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(uLCCCH1MessageTypeConstraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "UL-CCCH1-MessageType", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case UL_CCCH1_MessageType_C1:
		ie.C1 = &struct {
			Choice            int
			RrcResumeRequest1 *RRCResumeRequest1
			Spare3            *struct{}
			Spare2            *struct{}
			Spare1            *struct{}
		}{}
		choiceDec := d.NewChoiceDecoder(uLCCCH1MessageTypeC1Constraints)
		index, _, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		(*ie.C1).Choice = int(index)
		switch index {
		case UL_CCCH1_MessageType_C1_RrcResumeRequest1:
			(*ie.C1).RrcResumeRequest1 = new(RRCResumeRequest1)
			if err := (*ie.C1).RrcResumeRequest1.Decode(d); err != nil {
				return err
			}
		case UL_CCCH1_MessageType_C1_Spare3:
			(*ie.C1).Spare3 = &struct{}{}
			if err := d.DecodeNull(); err != nil {
				return err
			}
		case UL_CCCH1_MessageType_C1_Spare2:
			(*ie.C1).Spare2 = &struct{}{}
			if err := d.DecodeNull(); err != nil {
				return err
			}
		case UL_CCCH1_MessageType_C1_Spare1:
			(*ie.C1).Spare1 = &struct{}{}
			if err := d.DecodeNull(); err != nil {
				return err
			}
		}
	case UL_CCCH1_MessageType_MessageClassExtension:
		ie.MessageClassExtension = &struct{}{}
		// empty SEQUENCE alternative
	default:
		return &per.DecodeError{TypeName: "UL-CCCH1-MessageType", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
