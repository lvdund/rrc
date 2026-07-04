// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PerRAInfo-r16 (line 3158).

var perRAInfoR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "perRASSBInfoList-r16"},
		{Name: "perRACSI-RSInfoList-r16"},
	},
}

const (
	PerRAInfo_r16_PerRASSBInfoList_r16    = 0
	PerRAInfo_r16_PerRACSI_RSInfoList_r16 = 1
)

type PerRAInfo_r16 struct {
	Choice                  int
	PerRASSBInfoList_r16    *PerRASSBInfo_r16
	PerRACSI_RSInfoList_r16 *PerRACSI_RSInfo_r16
}

func (ie *PerRAInfo_r16) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(perRAInfoR16Constraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case PerRAInfo_r16_PerRASSBInfoList_r16:
		if err := ie.PerRASSBInfoList_r16.Encode(e); err != nil {
			return err
		}
	case PerRAInfo_r16_PerRACSI_RSInfoList_r16:
		if err := ie.PerRACSI_RSInfoList_r16.Encode(e); err != nil {
			return err
		}
	default:
		// Choice index out of root range.
		return &per.ConstraintViolationError{
			TypeName:   "PerRAInfo-r16",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *PerRAInfo_r16) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(perRAInfoR16Constraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "PerRAInfo-r16", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case PerRAInfo_r16_PerRASSBInfoList_r16:
		ie.PerRASSBInfoList_r16 = new(PerRASSBInfo_r16)
		if err := ie.PerRASSBInfoList_r16.Decode(d); err != nil {
			return err
		}
	case PerRAInfo_r16_PerRACSI_RSInfoList_r16:
		ie.PerRACSI_RSInfoList_r16 = new(PerRACSI_RSInfo_r16)
		if err := ie.PerRACSI_RSInfoList_r16.Decode(d); err != nil {
			return err
		}
	default:
		return &per.DecodeError{TypeName: "PerRAInfo-r16", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
