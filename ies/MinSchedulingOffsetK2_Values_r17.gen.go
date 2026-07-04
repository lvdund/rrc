// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MinSchedulingOffsetK2-Values-r17 (line 12461).
// MinSchedulingOffsetK2-Values-r17 ::=    SEQUENCE (SIZE (1..maxNrOfMinSchedulingOffsetValues-r16)) OF INTEGER (0..maxK2-SchedulingOffset-r17)

var minSchedulingOffsetK2ValuesR17SizeConstraints = per.SizeRange(1, common.MaxNrOfMinSchedulingOffsetValues_r16)

type MinSchedulingOffsetK2_Values_r17 []int64

func (ie *MinSchedulingOffsetK2_Values_r17) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(minSchedulingOffsetK2ValuesR17SizeConstraints)
	if err := seqOf.EncodeLength(int64(len(*ie))); err != nil {
		return err
	}
	for i := range *ie {
		if err := e.EncodeInteger((*ie)[i], per.Constrained(0, common.MaxK2_SchedulingOffset_r17)); err != nil {
			return err
		}
	}
	return nil
}

func (ie *MinSchedulingOffsetK2_Values_r17) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(minSchedulingOffsetK2ValuesR17SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(MinSchedulingOffsetK2_Values_r17, n)
	for i := int64(0); i < n; i++ {
		v, err := d.DecodeInteger(per.Constrained(0, common.MaxK2_SchedulingOffset_r17))
		if err != nil {
			return err
		}
		(*ie)[i] = v
	}
	return nil
}
