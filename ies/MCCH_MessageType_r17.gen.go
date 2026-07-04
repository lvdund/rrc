// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MCCH-MessageType-r17 (line 85).

var mCCHMessageTypeR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "c1"},
		{Name: "messageClassExtension"},
	},
}

const (
	MCCH_MessageType_r17_C1                    = 0
	MCCH_MessageType_r17_MessageClassExtension = 1
)

var mCCHMessageTypeR17C1Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "mbsBroadcastConfiguration-r17"},
		{Name: "spare1"},
	},
}

const (
	MCCH_MessageType_r17_C1_MbsBroadcastConfiguration_r17 = 0
	MCCH_MessageType_r17_C1_Spare1                        = 1
)

type MCCH_MessageType_r17 struct {
	Choice int
	C1     *struct {
		Choice                        int
		MbsBroadcastConfiguration_r17 *MBSBroadcastConfiguration_r17
		Spare1                        *struct{}
	}
	MessageClassExtension *struct{}
}

func (ie *MCCH_MessageType_r17) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(mCCHMessageTypeR17Constraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case MCCH_MessageType_r17_C1:
		choiceEnc := e.NewChoiceEncoder(mCCHMessageTypeR17C1Constraints)
		if err := choiceEnc.EncodeChoice(int64((*ie.C1).Choice), false, nil); err != nil {
			return err
		}
		switch (*ie.C1).Choice {
		case MCCH_MessageType_r17_C1_MbsBroadcastConfiguration_r17:
			if err := (*ie.C1).MbsBroadcastConfiguration_r17.Encode(e); err != nil {
				return err
			}
		case MCCH_MessageType_r17_C1_Spare1:
			if err := e.EncodeNull(); err != nil {
				return err
			}
		}
	case MCCH_MessageType_r17_MessageClassExtension:
		// empty SEQUENCE alternative
	default:
		// Choice index out of root range.
		return &per.ConstraintViolationError{
			TypeName:   "MCCH-MessageType-r17",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *MCCH_MessageType_r17) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(mCCHMessageTypeR17Constraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "MCCH-MessageType-r17", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case MCCH_MessageType_r17_C1:
		ie.C1 = &struct {
			Choice                        int
			MbsBroadcastConfiguration_r17 *MBSBroadcastConfiguration_r17
			Spare1                        *struct{}
		}{}
		choiceDec := d.NewChoiceDecoder(mCCHMessageTypeR17C1Constraints)
		index, _, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		(*ie.C1).Choice = int(index)
		switch index {
		case MCCH_MessageType_r17_C1_MbsBroadcastConfiguration_r17:
			(*ie.C1).MbsBroadcastConfiguration_r17 = new(MBSBroadcastConfiguration_r17)
			if err := (*ie.C1).MbsBroadcastConfiguration_r17.Decode(d); err != nil {
				return err
			}
		case MCCH_MessageType_r17_C1_Spare1:
			(*ie.C1).Spare1 = &struct{}{}
			if err := d.DecodeNull(); err != nil {
				return err
			}
		}
	case MCCH_MessageType_r17_MessageClassExtension:
		ie.MessageClassExtension = &struct{}{}
		// empty SEQUENCE alternative
	default:
		return &per.DecodeError{TypeName: "MCCH-MessageType-r17", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
