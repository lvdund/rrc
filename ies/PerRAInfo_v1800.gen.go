// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PerRAInfo-v1800 (line 3165).

var perRAInfoV1800Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "perRASSBInfoList-v1800"},
		{Name: "perRACSI-RSInfoList-v1800"},
	},
}

const (
	PerRAInfo_v1800_PerRASSBInfoList_v1800    = 0
	PerRAInfo_v1800_PerRACSI_RSInfoList_v1800 = 1
)

type PerRAInfo_v1800 struct {
	Choice                    int
	PerRASSBInfoList_v1800    *PerRASSBInfo_v1800
	PerRACSI_RSInfoList_v1800 *PerRACSI_RSInfo_v1800
}

func (ie *PerRAInfo_v1800) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(perRAInfoV1800Constraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case PerRAInfo_v1800_PerRASSBInfoList_v1800:
		if err := ie.PerRASSBInfoList_v1800.Encode(e); err != nil {
			return err
		}
	case PerRAInfo_v1800_PerRACSI_RSInfoList_v1800:
		if err := ie.PerRACSI_RSInfoList_v1800.Encode(e); err != nil {
			return err
		}
	default:
		// Choice index out of root range.
		return &per.ConstraintViolationError{
			TypeName:   "PerRAInfo-v1800",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *PerRAInfo_v1800) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(perRAInfoV1800Constraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "PerRAInfo-v1800", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case PerRAInfo_v1800_PerRASSBInfoList_v1800:
		ie.PerRASSBInfoList_v1800 = new(PerRASSBInfo_v1800)
		if err := ie.PerRASSBInfoList_v1800.Decode(d); err != nil {
			return err
		}
	case PerRAInfo_v1800_PerRACSI_RSInfoList_v1800:
		ie.PerRACSI_RSInfoList_v1800 = new(PerRACSI_RSInfo_v1800)
		if err := ie.PerRACSI_RSInfoList_v1800.Decode(d); err != nil {
			return err
		}
	default:
		return &per.DecodeError{TypeName: "PerRAInfo-v1800", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
