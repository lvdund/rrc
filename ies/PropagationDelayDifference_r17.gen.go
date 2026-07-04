// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PropagationDelayDifference-r17 (line 2713).
// PropagationDelayDifference-r17 ::=  SEQUENCE (SIZE (1..4)) OF INTEGER (-270..270)

var propagationDelayDifferenceR17SizeConstraints = per.SizeRange(1, 4)

type PropagationDelayDifference_r17 []int64

func (ie *PropagationDelayDifference_r17) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(propagationDelayDifferenceR17SizeConstraints)
	if err := seqOf.EncodeLength(int64(len(*ie))); err != nil {
		return err
	}
	for i := range *ie {
		if err := e.EncodeInteger((*ie)[i], per.Constrained(-270, 270)); err != nil {
			return err
		}
	}
	return nil
}

func (ie *PropagationDelayDifference_r17) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(propagationDelayDifferenceR17SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(PropagationDelayDifference_r17, n)
	for i := int64(0); i < n; i++ {
		v, err := d.DecodeInteger(per.Constrained(-270, 270))
		if err != nil {
			return err
		}
		(*ie)[i] = v
	}
	return nil
}
