// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MinSchedulingOffsetK0-Values-r16 (line 11443).
// MinSchedulingOffsetK0-Values-r16 ::=    SEQUENCE (SIZE (1..maxNrOfMinSchedulingOffsetValues-r16)) OF INTEGER (0..maxK0-SchedulingOffset-r16)

var minSchedulingOffsetK0ValuesR16SizeConstraints = per.SizeRange(1, common.MaxNrOfMinSchedulingOffsetValues_r16)

type MinSchedulingOffsetK0_Values_r16 []int64

func (ie *MinSchedulingOffsetK0_Values_r16) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(minSchedulingOffsetK0ValuesR16SizeConstraints)
	if err := seqOf.EncodeLength(int64(len(*ie))); err != nil {
		return err
	}
	for i := range *ie {
		if err := e.EncodeInteger((*ie)[i], per.Constrained(0, common.MaxK0_SchedulingOffset_r16)); err != nil {
			return err
		}
	}
	return nil
}

func (ie *MinSchedulingOffsetK0_Values_r16) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(minSchedulingOffsetK0ValuesR16SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(MinSchedulingOffsetK0_Values_r16, n)
	for i := int64(0); i < n; i++ {
		v, err := d.DecodeInteger(per.Constrained(0, common.MaxK0_SchedulingOffset_r16))
		if err != nil {
			return err
		}
		(*ie)[i] = v
	}
	return nil
}
