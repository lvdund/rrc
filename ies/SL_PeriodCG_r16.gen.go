// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-PeriodCG-r16 (line 27086).

var sLPeriodCGR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "sl-PeriodCG1-r16"},
		{Name: "sl-PeriodCG2-r16"},
	},
}

const (
	SL_PeriodCG_r16_Sl_PeriodCG1_r16 = 0
	SL_PeriodCG_r16_Sl_PeriodCG2_r16 = 1
)

const (
	SL_PeriodCG_r16_Sl_PeriodCG1_r16_Ms100  = 0
	SL_PeriodCG_r16_Sl_PeriodCG1_r16_Ms200  = 1
	SL_PeriodCG_r16_Sl_PeriodCG1_r16_Ms300  = 2
	SL_PeriodCG_r16_Sl_PeriodCG1_r16_Ms400  = 3
	SL_PeriodCG_r16_Sl_PeriodCG1_r16_Ms500  = 4
	SL_PeriodCG_r16_Sl_PeriodCG1_r16_Ms600  = 5
	SL_PeriodCG_r16_Sl_PeriodCG1_r16_Ms700  = 6
	SL_PeriodCG_r16_Sl_PeriodCG1_r16_Ms800  = 7
	SL_PeriodCG_r16_Sl_PeriodCG1_r16_Ms900  = 8
	SL_PeriodCG_r16_Sl_PeriodCG1_r16_Ms1000 = 9
	SL_PeriodCG_r16_Sl_PeriodCG1_r16_Spare6 = 10
	SL_PeriodCG_r16_Sl_PeriodCG1_r16_Spare5 = 11
	SL_PeriodCG_r16_Sl_PeriodCG1_r16_Spare4 = 12
	SL_PeriodCG_r16_Sl_PeriodCG1_r16_Spare3 = 13
	SL_PeriodCG_r16_Sl_PeriodCG1_r16_Spare2 = 14
	SL_PeriodCG_r16_Sl_PeriodCG1_r16_Spare1 = 15
)

var sLPeriodCGR16SlPeriodCG1R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_PeriodCG_r16_Sl_PeriodCG1_r16_Ms100, SL_PeriodCG_r16_Sl_PeriodCG1_r16_Ms200, SL_PeriodCG_r16_Sl_PeriodCG1_r16_Ms300, SL_PeriodCG_r16_Sl_PeriodCG1_r16_Ms400, SL_PeriodCG_r16_Sl_PeriodCG1_r16_Ms500, SL_PeriodCG_r16_Sl_PeriodCG1_r16_Ms600, SL_PeriodCG_r16_Sl_PeriodCG1_r16_Ms700, SL_PeriodCG_r16_Sl_PeriodCG1_r16_Ms800, SL_PeriodCG_r16_Sl_PeriodCG1_r16_Ms900, SL_PeriodCG_r16_Sl_PeriodCG1_r16_Ms1000, SL_PeriodCG_r16_Sl_PeriodCG1_r16_Spare6, SL_PeriodCG_r16_Sl_PeriodCG1_r16_Spare5, SL_PeriodCG_r16_Sl_PeriodCG1_r16_Spare4, SL_PeriodCG_r16_Sl_PeriodCG1_r16_Spare3, SL_PeriodCG_r16_Sl_PeriodCG1_r16_Spare2, SL_PeriodCG_r16_Sl_PeriodCG1_r16_Spare1},
}

type SL_PeriodCG_r16 struct {
	Choice           int
	Sl_PeriodCG1_r16 *int64
	Sl_PeriodCG2_r16 *int64
}

func (ie *SL_PeriodCG_r16) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(sLPeriodCGR16Constraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case SL_PeriodCG_r16_Sl_PeriodCG1_r16:
		if err := e.EncodeEnumerated((*ie.Sl_PeriodCG1_r16), sLPeriodCGR16SlPeriodCG1R16Constraints); err != nil {
			return err
		}
	case SL_PeriodCG_r16_Sl_PeriodCG2_r16:
		if err := e.EncodeInteger((*ie.Sl_PeriodCG2_r16), per.Constrained(1, 99)); err != nil {
			return err
		}
	default:
		// Choice index out of root range.
		return &per.ConstraintViolationError{
			TypeName:   "SL-PeriodCG-r16",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *SL_PeriodCG_r16) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(sLPeriodCGR16Constraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "SL-PeriodCG-r16", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case SL_PeriodCG_r16_Sl_PeriodCG1_r16:
		ie.Sl_PeriodCG1_r16 = new(int64)
		v, err := d.DecodeEnumerated(sLPeriodCGR16SlPeriodCG1R16Constraints)
		if err != nil {
			return err
		}
		(*ie.Sl_PeriodCG1_r16) = v
	case SL_PeriodCG_r16_Sl_PeriodCG2_r16:
		ie.Sl_PeriodCG2_r16 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(1, 99))
		if err != nil {
			return err
		}
		(*ie.Sl_PeriodCG2_r16) = v
	default:
		return &per.DecodeError{TypeName: "SL-PeriodCG-r16", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
