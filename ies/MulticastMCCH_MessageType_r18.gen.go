// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MulticastMCCH-MessageType-r18 (line 100).

var multicastMCCHMessageTypeR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "c1"},
		{Name: "messageClassExtension"},
	},
}

const (
	MulticastMCCH_MessageType_r18_C1                    = 0
	MulticastMCCH_MessageType_r18_MessageClassExtension = 1
)

var multicastMCCHMessageTypeR18C1Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "mbsMulticastConfiguration-r18"},
		{Name: "spare1"},
	},
}

const (
	MulticastMCCH_MessageType_r18_C1_MbsMulticastConfiguration_r18 = 0
	MulticastMCCH_MessageType_r18_C1_Spare1                        = 1
)

type MulticastMCCH_MessageType_r18 struct {
	Choice int
	C1     *struct {
		Choice                        int
		MbsMulticastConfiguration_r18 *MBSMulticastConfiguration_r18
		Spare1                        *struct{}
	}
	MessageClassExtension *struct{}
}

func (ie *MulticastMCCH_MessageType_r18) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(multicastMCCHMessageTypeR18Constraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case MulticastMCCH_MessageType_r18_C1:
		choiceEnc := e.NewChoiceEncoder(multicastMCCHMessageTypeR18C1Constraints)
		if err := choiceEnc.EncodeChoice(int64((*ie.C1).Choice), false, nil); err != nil {
			return err
		}
		switch (*ie.C1).Choice {
		case MulticastMCCH_MessageType_r18_C1_MbsMulticastConfiguration_r18:
			if err := (*ie.C1).MbsMulticastConfiguration_r18.Encode(e); err != nil {
				return err
			}
		case MulticastMCCH_MessageType_r18_C1_Spare1:
			if err := e.EncodeNull(); err != nil {
				return err
			}
		}
	case MulticastMCCH_MessageType_r18_MessageClassExtension:
		// empty SEQUENCE alternative
	default:
		// Choice index out of root range.
		return &per.ConstraintViolationError{
			TypeName:   "MulticastMCCH-MessageType-r18",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *MulticastMCCH_MessageType_r18) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(multicastMCCHMessageTypeR18Constraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "MulticastMCCH-MessageType-r18", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case MulticastMCCH_MessageType_r18_C1:
		ie.C1 = &struct {
			Choice                        int
			MbsMulticastConfiguration_r18 *MBSMulticastConfiguration_r18
			Spare1                        *struct{}
		}{}
		choiceDec := d.NewChoiceDecoder(multicastMCCHMessageTypeR18C1Constraints)
		index, _, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		(*ie.C1).Choice = int(index)
		switch index {
		case MulticastMCCH_MessageType_r18_C1_MbsMulticastConfiguration_r18:
			(*ie.C1).MbsMulticastConfiguration_r18 = new(MBSMulticastConfiguration_r18)
			if err := (*ie.C1).MbsMulticastConfiguration_r18.Decode(d); err != nil {
				return err
			}
		case MulticastMCCH_MessageType_r18_C1_Spare1:
			(*ie.C1).Spare1 = &struct{}{}
			if err := d.DecodeNull(); err != nil {
				return err
			}
		}
	case MulticastMCCH_MessageType_r18_MessageClassExtension:
		ie.MessageClassExtension = &struct{}{}
		// empty SEQUENCE alternative
	default:
		return &per.DecodeError{TypeName: "MulticastMCCH-MessageType-r18", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
