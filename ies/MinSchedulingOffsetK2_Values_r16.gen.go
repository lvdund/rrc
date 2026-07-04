// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MinSchedulingOffsetK2-Values-r16 (line 12459).
// MinSchedulingOffsetK2-Values-r16 ::=    SEQUENCE (SIZE (1..maxNrOfMinSchedulingOffsetValues-r16)) OF INTEGER (0..maxK2-SchedulingOffset-r16)

var minSchedulingOffsetK2ValuesR16SizeConstraints = per.SizeRange(1, common.MaxNrOfMinSchedulingOffsetValues_r16)

type MinSchedulingOffsetK2_Values_r16 []int64

func (ie *MinSchedulingOffsetK2_Values_r16) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(minSchedulingOffsetK2ValuesR16SizeConstraints)
	if err := seqOf.EncodeLength(int64(len(*ie))); err != nil {
		return err
	}
	for i := range *ie {
		if err := e.EncodeInteger((*ie)[i], per.Constrained(0, common.MaxK2_SchedulingOffset_r16)); err != nil {
			return err
		}
	}
	return nil
}

func (ie *MinSchedulingOffsetK2_Values_r16) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(minSchedulingOffsetK2ValuesR16SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(MinSchedulingOffsetK2_Values_r16, n)
	for i := int64(0); i < n; i++ {
		v, err := d.DecodeInteger(per.Constrained(0, common.MaxK2_SchedulingOffset_r16))
		if err != nil {
			return err
		}
		(*ie)[i] = v
	}
	return nil
}
