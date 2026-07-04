// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: DL-CCCH-MessageType (line 41).

var dLCCCHMessageTypeConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "c1"},
		{Name: "messageClassExtension"},
	},
}

const (
	DL_CCCH_MessageType_C1                    = 0
	DL_CCCH_MessageType_MessageClassExtension = 1
)

var dLCCCHMessageTypeC1Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "rrcReject"},
		{Name: "rrcSetup"},
		{Name: "spare2"},
		{Name: "spare1"},
	},
}

const (
	DL_CCCH_MessageType_C1_RrcReject = 0
	DL_CCCH_MessageType_C1_RrcSetup  = 1
	DL_CCCH_MessageType_C1_Spare2    = 2
	DL_CCCH_MessageType_C1_Spare1    = 3
)

type DL_CCCH_MessageType struct {
	Choice int
	C1     *struct {
		Choice    int
		RrcReject *RRCReject
		RrcSetup  *RRCSetup
		Spare2    *struct{}
		Spare1    *struct{}
	}
	MessageClassExtension *struct{}
}

func (ie *DL_CCCH_MessageType) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(dLCCCHMessageTypeConstraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case DL_CCCH_MessageType_C1:
		choiceEnc := e.NewChoiceEncoder(dLCCCHMessageTypeC1Constraints)
		if err := choiceEnc.EncodeChoice(int64((*ie.C1).Choice), false, nil); err != nil {
			return err
		}
		switch (*ie.C1).Choice {
		case DL_CCCH_MessageType_C1_RrcReject:
			if err := (*ie.C1).RrcReject.Encode(e); err != nil {
				return err
			}
		case DL_CCCH_MessageType_C1_RrcSetup:
			if err := (*ie.C1).RrcSetup.Encode(e); err != nil {
				return err
			}
		case DL_CCCH_MessageType_C1_Spare2:
			if err := e.EncodeNull(); err != nil {
				return err
			}
		case DL_CCCH_MessageType_C1_Spare1:
			if err := e.EncodeNull(); err != nil {
				return err
			}
		}
	case DL_CCCH_MessageType_MessageClassExtension:
		// empty SEQUENCE alternative
	default:
		// Choice index out of root range.
		return &per.ConstraintViolationError{
			TypeName:   "DL-CCCH-MessageType",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *DL_CCCH_MessageType) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(dLCCCHMessageTypeConstraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "DL-CCCH-MessageType", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case DL_CCCH_MessageType_C1:
		ie.C1 = &struct {
			Choice    int
			RrcReject *RRCReject
			RrcSetup  *RRCSetup
			Spare2    *struct{}
			Spare1    *struct{}
		}{}
		choiceDec := d.NewChoiceDecoder(dLCCCHMessageTypeC1Constraints)
		index, _, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		(*ie.C1).Choice = int(index)
		switch index {
		case DL_CCCH_MessageType_C1_RrcReject:
			(*ie.C1).RrcReject = new(RRCReject)
			if err := (*ie.C1).RrcReject.Decode(d); err != nil {
				return err
			}
		case DL_CCCH_MessageType_C1_RrcSetup:
			(*ie.C1).RrcSetup = new(RRCSetup)
			if err := (*ie.C1).RrcSetup.Decode(d); err != nil {
				return err
			}
		case DL_CCCH_MessageType_C1_Spare2:
			(*ie.C1).Spare2 = &struct{}{}
			if err := d.DecodeNull(); err != nil {
				return err
			}
		case DL_CCCH_MessageType_C1_Spare1:
			(*ie.C1).Spare1 = &struct{}{}
			if err := d.DecodeNull(); err != nil {
				return err
			}
		}
	case DL_CCCH_MessageType_MessageClassExtension:
		ie.MessageClassExtension = &struct{}{}
		// empty SEQUENCE alternative
	default:
		return &per.DecodeError{TypeName: "DL-CCCH-MessageType", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
