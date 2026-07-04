// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MCCH-RepetitionPeriodAndOffset-r17 (line 4576).

var mCCHRepetitionPeriodAndOffsetR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "rf1-r17"},
		{Name: "rf2-r17"},
		{Name: "rf4-r17"},
		{Name: "rf8-r17"},
		{Name: "rf16-r17"},
		{Name: "rf32-r17"},
		{Name: "rf64-r17"},
		{Name: "rf128-r17"},
		{Name: "rf256-r17"},
	},
}

const (
	MCCH_RepetitionPeriodAndOffset_r17_Rf1_r17   = 0
	MCCH_RepetitionPeriodAndOffset_r17_Rf2_r17   = 1
	MCCH_RepetitionPeriodAndOffset_r17_Rf4_r17   = 2
	MCCH_RepetitionPeriodAndOffset_r17_Rf8_r17   = 3
	MCCH_RepetitionPeriodAndOffset_r17_Rf16_r17  = 4
	MCCH_RepetitionPeriodAndOffset_r17_Rf32_r17  = 5
	MCCH_RepetitionPeriodAndOffset_r17_Rf64_r17  = 6
	MCCH_RepetitionPeriodAndOffset_r17_Rf128_r17 = 7
	MCCH_RepetitionPeriodAndOffset_r17_Rf256_r17 = 8
)

type MCCH_RepetitionPeriodAndOffset_r17 struct {
	Choice    int
	Rf1_r17   *int64
	Rf2_r17   *int64
	Rf4_r17   *int64
	Rf8_r17   *int64
	Rf16_r17  *int64
	Rf32_r17  *int64
	Rf64_r17  *int64
	Rf128_r17 *int64
	Rf256_r17 *int64
}

func (ie *MCCH_RepetitionPeriodAndOffset_r17) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(mCCHRepetitionPeriodAndOffsetR17Constraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case MCCH_RepetitionPeriodAndOffset_r17_Rf1_r17:
		if err := e.EncodeInteger((*ie.Rf1_r17), per.Constrained(0, 0)); err != nil {
			return err
		}
	case MCCH_RepetitionPeriodAndOffset_r17_Rf2_r17:
		if err := e.EncodeInteger((*ie.Rf2_r17), per.Constrained(0, 1)); err != nil {
			return err
		}
	case MCCH_RepetitionPeriodAndOffset_r17_Rf4_r17:
		if err := e.EncodeInteger((*ie.Rf4_r17), per.Constrained(0, 3)); err != nil {
			return err
		}
	case MCCH_RepetitionPeriodAndOffset_r17_Rf8_r17:
		if err := e.EncodeInteger((*ie.Rf8_r17), per.Constrained(0, 7)); err != nil {
			return err
		}
	case MCCH_RepetitionPeriodAndOffset_r17_Rf16_r17:
		if err := e.EncodeInteger((*ie.Rf16_r17), per.Constrained(0, 15)); err != nil {
			return err
		}
	case MCCH_RepetitionPeriodAndOffset_r17_Rf32_r17:
		if err := e.EncodeInteger((*ie.Rf32_r17), per.Constrained(0, 31)); err != nil {
			return err
		}
	case MCCH_RepetitionPeriodAndOffset_r17_Rf64_r17:
		if err := e.EncodeInteger((*ie.Rf64_r17), per.Constrained(0, 63)); err != nil {
			return err
		}
	case MCCH_RepetitionPeriodAndOffset_r17_Rf128_r17:
		if err := e.EncodeInteger((*ie.Rf128_r17), per.Constrained(0, 127)); err != nil {
			return err
		}
	case MCCH_RepetitionPeriodAndOffset_r17_Rf256_r17:
		if err := e.EncodeInteger((*ie.Rf256_r17), per.Constrained(0, 255)); err != nil {
			return err
		}
	default:
		// Choice index out of root range.
		return &per.ConstraintViolationError{
			TypeName:   "MCCH-RepetitionPeriodAndOffset-r17",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *MCCH_RepetitionPeriodAndOffset_r17) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(mCCHRepetitionPeriodAndOffsetR17Constraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "MCCH-RepetitionPeriodAndOffset-r17", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case MCCH_RepetitionPeriodAndOffset_r17_Rf1_r17:
		ie.Rf1_r17 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 0))
		if err != nil {
			return err
		}
		(*ie.Rf1_r17) = v
	case MCCH_RepetitionPeriodAndOffset_r17_Rf2_r17:
		ie.Rf2_r17 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 1))
		if err != nil {
			return err
		}
		(*ie.Rf2_r17) = v
	case MCCH_RepetitionPeriodAndOffset_r17_Rf4_r17:
		ie.Rf4_r17 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 3))
		if err != nil {
			return err
		}
		(*ie.Rf4_r17) = v
	case MCCH_RepetitionPeriodAndOffset_r17_Rf8_r17:
		ie.Rf8_r17 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 7))
		if err != nil {
			return err
		}
		(*ie.Rf8_r17) = v
	case MCCH_RepetitionPeriodAndOffset_r17_Rf16_r17:
		ie.Rf16_r17 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 15))
		if err != nil {
			return err
		}
		(*ie.Rf16_r17) = v
	case MCCH_RepetitionPeriodAndOffset_r17_Rf32_r17:
		ie.Rf32_r17 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 31))
		if err != nil {
			return err
		}
		(*ie.Rf32_r17) = v
	case MCCH_RepetitionPeriodAndOffset_r17_Rf64_r17:
		ie.Rf64_r17 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 63))
		if err != nil {
			return err
		}
		(*ie.Rf64_r17) = v
	case MCCH_RepetitionPeriodAndOffset_r17_Rf128_r17:
		ie.Rf128_r17 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 127))
		if err != nil {
			return err
		}
		(*ie.Rf128_r17) = v
	case MCCH_RepetitionPeriodAndOffset_r17_Rf256_r17:
		ie.Rf256_r17 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 255))
		if err != nil {
			return err
		}
		(*ie.Rf256_r17) = v
	default:
		return &per.DecodeError{TypeName: "MCCH-RepetitionPeriodAndOffset-r17", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
