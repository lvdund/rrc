// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SRS-PeriodicityAndOffsetExt-r16 (line 15700).

var sRSPeriodicityAndOffsetExtR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "sl128"},
		{Name: "sl256"},
		{Name: "sl512"},
		{Name: "sl20480"},
	},
}

const (
	SRS_PeriodicityAndOffsetExt_r16_Sl128   = 0
	SRS_PeriodicityAndOffsetExt_r16_Sl256   = 1
	SRS_PeriodicityAndOffsetExt_r16_Sl512   = 2
	SRS_PeriodicityAndOffsetExt_r16_Sl20480 = 3
)

type SRS_PeriodicityAndOffsetExt_r16 struct {
	Choice  int
	Sl128   *int64
	Sl256   *int64
	Sl512   *int64
	Sl20480 *int64
}

func (ie *SRS_PeriodicityAndOffsetExt_r16) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(sRSPeriodicityAndOffsetExtR16Constraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case SRS_PeriodicityAndOffsetExt_r16_Sl128:
		if err := e.EncodeInteger((*ie.Sl128), per.Constrained(0, 127)); err != nil {
			return err
		}
	case SRS_PeriodicityAndOffsetExt_r16_Sl256:
		if err := e.EncodeInteger((*ie.Sl256), per.Constrained(0, 255)); err != nil {
			return err
		}
	case SRS_PeriodicityAndOffsetExt_r16_Sl512:
		if err := e.EncodeInteger((*ie.Sl512), per.Constrained(0, 511)); err != nil {
			return err
		}
	case SRS_PeriodicityAndOffsetExt_r16_Sl20480:
		if err := e.EncodeInteger((*ie.Sl20480), per.Constrained(0, 20479)); err != nil {
			return err
		}
	default:
		// Choice index out of root range.
		return &per.ConstraintViolationError{
			TypeName:   "SRS-PeriodicityAndOffsetExt-r16",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *SRS_PeriodicityAndOffsetExt_r16) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(sRSPeriodicityAndOffsetExtR16Constraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "SRS-PeriodicityAndOffsetExt-r16", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case SRS_PeriodicityAndOffsetExt_r16_Sl128:
		ie.Sl128 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 127))
		if err != nil {
			return err
		}
		(*ie.Sl128) = v
	case SRS_PeriodicityAndOffsetExt_r16_Sl256:
		ie.Sl256 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 255))
		if err != nil {
			return err
		}
		(*ie.Sl256) = v
	case SRS_PeriodicityAndOffsetExt_r16_Sl512:
		ie.Sl512 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 511))
		if err != nil {
			return err
		}
		(*ie.Sl512) = v
	case SRS_PeriodicityAndOffsetExt_r16_Sl20480:
		ie.Sl20480 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 20479))
		if err != nil {
			return err
		}
		(*ie.Sl20480) = v
	default:
		return &per.DecodeError{TypeName: "SRS-PeriodicityAndOffsetExt-r16", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
