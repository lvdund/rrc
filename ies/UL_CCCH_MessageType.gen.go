// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UL-CCCH-MessageType (line 131).

var uLCCCHMessageTypeConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "c1"},
		{Name: "messageClassExtension"},
	},
}

const (
	UL_CCCH_MessageType_C1                    = 0
	UL_CCCH_MessageType_MessageClassExtension = 1
)

var uLCCCHMessageTypeC1Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "rrcSetupRequest"},
		{Name: "rrcResumeRequest"},
		{Name: "rrcReestablishmentRequest"},
		{Name: "rrcSystemInfoRequest"},
	},
}

const (
	UL_CCCH_MessageType_C1_RrcSetupRequest           = 0
	UL_CCCH_MessageType_C1_RrcResumeRequest          = 1
	UL_CCCH_MessageType_C1_RrcReestablishmentRequest = 2
	UL_CCCH_MessageType_C1_RrcSystemInfoRequest      = 3
)

type UL_CCCH_MessageType struct {
	Choice int
	C1     *struct {
		Choice                    int
		RrcSetupRequest           *RRCSetupRequest
		RrcResumeRequest          *RRCResumeRequest
		RrcReestablishmentRequest *RRCReestablishmentRequest
		RrcSystemInfoRequest      *RRCSystemInfoRequest
	}
	MessageClassExtension *struct{}
}

func (ie *UL_CCCH_MessageType) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(uLCCCHMessageTypeConstraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case UL_CCCH_MessageType_C1:
		choiceEnc := e.NewChoiceEncoder(uLCCCHMessageTypeC1Constraints)
		if err := choiceEnc.EncodeChoice(int64((*ie.C1).Choice), false, nil); err != nil {
			return err
		}
		switch (*ie.C1).Choice {
		case UL_CCCH_MessageType_C1_RrcSetupRequest:
			if err := (*ie.C1).RrcSetupRequest.Encode(e); err != nil {
				return err
			}
		case UL_CCCH_MessageType_C1_RrcResumeRequest:
			if err := (*ie.C1).RrcResumeRequest.Encode(e); err != nil {
				return err
			}
		case UL_CCCH_MessageType_C1_RrcReestablishmentRequest:
			if err := (*ie.C1).RrcReestablishmentRequest.Encode(e); err != nil {
				return err
			}
		case UL_CCCH_MessageType_C1_RrcSystemInfoRequest:
			if err := (*ie.C1).RrcSystemInfoRequest.Encode(e); err != nil {
				return err
			}
		}
	case UL_CCCH_MessageType_MessageClassExtension:
		// empty SEQUENCE alternative
	default:
		// Choice index out of root range.
		return &per.ConstraintViolationError{
			TypeName:   "UL-CCCH-MessageType",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *UL_CCCH_MessageType) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(uLCCCHMessageTypeConstraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "UL-CCCH-MessageType", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case UL_CCCH_MessageType_C1:
		ie.C1 = &struct {
			Choice                    int
			RrcSetupRequest           *RRCSetupRequest
			RrcResumeRequest          *RRCResumeRequest
			RrcReestablishmentRequest *RRCReestablishmentRequest
			RrcSystemInfoRequest      *RRCSystemInfoRequest
		}{}
		choiceDec := d.NewChoiceDecoder(uLCCCHMessageTypeC1Constraints)
		index, _, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		(*ie.C1).Choice = int(index)
		switch index {
		case UL_CCCH_MessageType_C1_RrcSetupRequest:
			(*ie.C1).RrcSetupRequest = new(RRCSetupRequest)
			if err := (*ie.C1).RrcSetupRequest.Decode(d); err != nil {
				return err
			}
		case UL_CCCH_MessageType_C1_RrcResumeRequest:
			(*ie.C1).RrcResumeRequest = new(RRCResumeRequest)
			if err := (*ie.C1).RrcResumeRequest.Decode(d); err != nil {
				return err
			}
		case UL_CCCH_MessageType_C1_RrcReestablishmentRequest:
			(*ie.C1).RrcReestablishmentRequest = new(RRCReestablishmentRequest)
			if err := (*ie.C1).RrcReestablishmentRequest.Decode(d); err != nil {
				return err
			}
		case UL_CCCH_MessageType_C1_RrcSystemInfoRequest:
			(*ie.C1).RrcSystemInfoRequest = new(RRCSystemInfoRequest)
			if err := (*ie.C1).RrcSystemInfoRequest.Decode(d); err != nil {
				return err
			}
		}
	case UL_CCCH_MessageType_MessageClassExtension:
		ie.MessageClassExtension = &struct{}{}
		// empty SEQUENCE alternative
	default:
		return &per.DecodeError{TypeName: "UL-CCCH-MessageType", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
