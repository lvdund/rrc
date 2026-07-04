// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MTCH-SSB-MappingWindowCycleOffset-r17 (line 28614).

var mTCHSSBMappingWindowCycleOffsetR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "ms10"},
		{Name: "ms20"},
		{Name: "ms32"},
		{Name: "ms64"},
		{Name: "ms128"},
		{Name: "ms256"},
	},
}

const (
	MTCH_SSB_MappingWindowCycleOffset_r17_Ms10  = 0
	MTCH_SSB_MappingWindowCycleOffset_r17_Ms20  = 1
	MTCH_SSB_MappingWindowCycleOffset_r17_Ms32  = 2
	MTCH_SSB_MappingWindowCycleOffset_r17_Ms64  = 3
	MTCH_SSB_MappingWindowCycleOffset_r17_Ms128 = 4
	MTCH_SSB_MappingWindowCycleOffset_r17_Ms256 = 5
)

type MTCH_SSB_MappingWindowCycleOffset_r17 struct {
	Choice int
	Ms10   *int64
	Ms20   *int64
	Ms32   *int64
	Ms64   *int64
	Ms128  *int64
	Ms256  *int64
}

func (ie *MTCH_SSB_MappingWindowCycleOffset_r17) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(mTCHSSBMappingWindowCycleOffsetR17Constraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case MTCH_SSB_MappingWindowCycleOffset_r17_Ms10:
		if err := e.EncodeInteger((*ie.Ms10), per.Constrained(0, 9)); err != nil {
			return err
		}
	case MTCH_SSB_MappingWindowCycleOffset_r17_Ms20:
		if err := e.EncodeInteger((*ie.Ms20), per.Constrained(0, 19)); err != nil {
			return err
		}
	case MTCH_SSB_MappingWindowCycleOffset_r17_Ms32:
		if err := e.EncodeInteger((*ie.Ms32), per.Constrained(0, 31)); err != nil {
			return err
		}
	case MTCH_SSB_MappingWindowCycleOffset_r17_Ms64:
		if err := e.EncodeInteger((*ie.Ms64), per.Constrained(0, 63)); err != nil {
			return err
		}
	case MTCH_SSB_MappingWindowCycleOffset_r17_Ms128:
		if err := e.EncodeInteger((*ie.Ms128), per.Constrained(0, 127)); err != nil {
			return err
		}
	case MTCH_SSB_MappingWindowCycleOffset_r17_Ms256:
		if err := e.EncodeInteger((*ie.Ms256), per.Constrained(0, 255)); err != nil {
			return err
		}
	default:
		// Choice index out of root range.
		return &per.ConstraintViolationError{
			TypeName:   "MTCH-SSB-MappingWindowCycleOffset-r17",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *MTCH_SSB_MappingWindowCycleOffset_r17) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(mTCHSSBMappingWindowCycleOffsetR17Constraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "MTCH-SSB-MappingWindowCycleOffset-r17", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case MTCH_SSB_MappingWindowCycleOffset_r17_Ms10:
		ie.Ms10 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 9))
		if err != nil {
			return err
		}
		(*ie.Ms10) = v
	case MTCH_SSB_MappingWindowCycleOffset_r17_Ms20:
		ie.Ms20 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 19))
		if err != nil {
			return err
		}
		(*ie.Ms20) = v
	case MTCH_SSB_MappingWindowCycleOffset_r17_Ms32:
		ie.Ms32 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 31))
		if err != nil {
			return err
		}
		(*ie.Ms32) = v
	case MTCH_SSB_MappingWindowCycleOffset_r17_Ms64:
		ie.Ms64 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 63))
		if err != nil {
			return err
		}
		(*ie.Ms64) = v
	case MTCH_SSB_MappingWindowCycleOffset_r17_Ms128:
		ie.Ms128 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 127))
		if err != nil {
			return err
		}
		(*ie.Ms128) = v
	case MTCH_SSB_MappingWindowCycleOffset_r17_Ms256:
		ie.Ms256 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 255))
		if err != nil {
			return err
		}
		(*ie.Ms256) = v
	default:
		return &per.DecodeError{TypeName: "MTCH-SSB-MappingWindowCycleOffset-r17", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
