// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-ResourceReservePeriod-r16 (line 28077).

var sLResourceReservePeriodR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "sl-ResourceReservePeriod1-r16"},
		{Name: "sl-ResourceReservePeriod2-r16"},
	},
}

const (
	SL_ResourceReservePeriod_r16_Sl_ResourceReservePeriod1_r16 = 0
	SL_ResourceReservePeriod_r16_Sl_ResourceReservePeriod2_r16 = 1
)

const (
	SL_ResourceReservePeriod_r16_Sl_ResourceReservePeriod1_r16_Ms0    = 0
	SL_ResourceReservePeriod_r16_Sl_ResourceReservePeriod1_r16_Ms100  = 1
	SL_ResourceReservePeriod_r16_Sl_ResourceReservePeriod1_r16_Ms200  = 2
	SL_ResourceReservePeriod_r16_Sl_ResourceReservePeriod1_r16_Ms300  = 3
	SL_ResourceReservePeriod_r16_Sl_ResourceReservePeriod1_r16_Ms400  = 4
	SL_ResourceReservePeriod_r16_Sl_ResourceReservePeriod1_r16_Ms500  = 5
	SL_ResourceReservePeriod_r16_Sl_ResourceReservePeriod1_r16_Ms600  = 6
	SL_ResourceReservePeriod_r16_Sl_ResourceReservePeriod1_r16_Ms700  = 7
	SL_ResourceReservePeriod_r16_Sl_ResourceReservePeriod1_r16_Ms800  = 8
	SL_ResourceReservePeriod_r16_Sl_ResourceReservePeriod1_r16_Ms900  = 9
	SL_ResourceReservePeriod_r16_Sl_ResourceReservePeriod1_r16_Ms1000 = 10
)

var sLResourceReservePeriodR16SlResourceReservePeriod1R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_ResourceReservePeriod_r16_Sl_ResourceReservePeriod1_r16_Ms0, SL_ResourceReservePeriod_r16_Sl_ResourceReservePeriod1_r16_Ms100, SL_ResourceReservePeriod_r16_Sl_ResourceReservePeriod1_r16_Ms200, SL_ResourceReservePeriod_r16_Sl_ResourceReservePeriod1_r16_Ms300, SL_ResourceReservePeriod_r16_Sl_ResourceReservePeriod1_r16_Ms400, SL_ResourceReservePeriod_r16_Sl_ResourceReservePeriod1_r16_Ms500, SL_ResourceReservePeriod_r16_Sl_ResourceReservePeriod1_r16_Ms600, SL_ResourceReservePeriod_r16_Sl_ResourceReservePeriod1_r16_Ms700, SL_ResourceReservePeriod_r16_Sl_ResourceReservePeriod1_r16_Ms800, SL_ResourceReservePeriod_r16_Sl_ResourceReservePeriod1_r16_Ms900, SL_ResourceReservePeriod_r16_Sl_ResourceReservePeriod1_r16_Ms1000},
}

type SL_ResourceReservePeriod_r16 struct {
	Choice                        int
	Sl_ResourceReservePeriod1_r16 *int64
	Sl_ResourceReservePeriod2_r16 *int64
}

func (ie *SL_ResourceReservePeriod_r16) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(sLResourceReservePeriodR16Constraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case SL_ResourceReservePeriod_r16_Sl_ResourceReservePeriod1_r16:
		if err := e.EncodeEnumerated((*ie.Sl_ResourceReservePeriod1_r16), sLResourceReservePeriodR16SlResourceReservePeriod1R16Constraints); err != nil {
			return err
		}
	case SL_ResourceReservePeriod_r16_Sl_ResourceReservePeriod2_r16:
		if err := e.EncodeInteger((*ie.Sl_ResourceReservePeriod2_r16), per.Constrained(1, 99)); err != nil {
			return err
		}
	default:
		// Choice index out of root range.
		return &per.ConstraintViolationError{
			TypeName:   "SL-ResourceReservePeriod-r16",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *SL_ResourceReservePeriod_r16) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(sLResourceReservePeriodR16Constraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "SL-ResourceReservePeriod-r16", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case SL_ResourceReservePeriod_r16_Sl_ResourceReservePeriod1_r16:
		ie.Sl_ResourceReservePeriod1_r16 = new(int64)
		v, err := d.DecodeEnumerated(sLResourceReservePeriodR16SlResourceReservePeriod1R16Constraints)
		if err != nil {
			return err
		}
		(*ie.Sl_ResourceReservePeriod1_r16) = v
	case SL_ResourceReservePeriod_r16_Sl_ResourceReservePeriod2_r16:
		ie.Sl_ResourceReservePeriod2_r16 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(1, 99))
		if err != nil {
			return err
		}
		(*ie.Sl_ResourceReservePeriod2_r16) = v
	default:
		return &per.DecodeError{TypeName: "SL-ResourceReservePeriod-r16", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
