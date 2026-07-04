// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SSB-ToMeasure (line 15923).

var sSBToMeasureConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "shortBitmap"},
		{Name: "mediumBitmap"},
		{Name: "longBitmap"},
	},
}

const (
	SSB_ToMeasure_ShortBitmap  = 0
	SSB_ToMeasure_MediumBitmap = 1
	SSB_ToMeasure_LongBitmap   = 2
)

type SSB_ToMeasure struct {
	Choice       int
	ShortBitmap  *per.BitString
	MediumBitmap *per.BitString
	LongBitmap   *per.BitString
}

func (ie *SSB_ToMeasure) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(sSBToMeasureConstraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case SSB_ToMeasure_ShortBitmap:
		if err := e.EncodeBitString((*ie.ShortBitmap), per.FixedSize(4)); err != nil {
			return err
		}
	case SSB_ToMeasure_MediumBitmap:
		if err := e.EncodeBitString((*ie.MediumBitmap), per.FixedSize(8)); err != nil {
			return err
		}
	case SSB_ToMeasure_LongBitmap:
		if err := e.EncodeBitString((*ie.LongBitmap), per.FixedSize(64)); err != nil {
			return err
		}
	default:
		// Choice index out of root range.
		return &per.ConstraintViolationError{
			TypeName:   "SSB-ToMeasure",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *SSB_ToMeasure) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(sSBToMeasureConstraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "SSB-ToMeasure", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case SSB_ToMeasure_ShortBitmap:
		ie.ShortBitmap = new(per.BitString)
		v, err := d.DecodeBitString(per.FixedSize(4))
		if err != nil {
			return err
		}
		(*ie.ShortBitmap) = v
	case SSB_ToMeasure_MediumBitmap:
		ie.MediumBitmap = new(per.BitString)
		v, err := d.DecodeBitString(per.FixedSize(8))
		if err != nil {
			return err
		}
		(*ie.MediumBitmap) = v
	case SSB_ToMeasure_LongBitmap:
		ie.LongBitmap = new(per.BitString)
		v, err := d.DecodeBitString(per.FixedSize(64))
		if err != nil {
			return err
		}
		(*ie.LongBitmap) = v
	default:
		return &per.DecodeError{TypeName: "SSB-ToMeasure", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
