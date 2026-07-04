// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: BetaOffsetsCrossPri-r17 (line 5183).
// BetaOffsetsCrossPri-r17 ::= SEQUENCE (SIZE(3)) OF INTEGER(0..31)

var betaOffsetsCrossPriR17SizeConstraints = per.FixedSize(3)

type BetaOffsetsCrossPri_r17 []int64

func (ie *BetaOffsetsCrossPri_r17) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(betaOffsetsCrossPriR17SizeConstraints)
	if err := seqOf.EncodeLength(int64(len(*ie))); err != nil {
		return err
	}
	for i := range *ie {
		if err := e.EncodeInteger((*ie)[i], per.Constrained(0, 31)); err != nil {
			return err
		}
	}
	return nil
}

func (ie *BetaOffsetsCrossPri_r17) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(betaOffsetsCrossPriR17SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(BetaOffsetsCrossPri_r17, n)
	for i := int64(0); i < n; i++ {
		v, err := d.DecodeInteger(per.Constrained(0, 31))
		if err != nil {
			return err
		}
		(*ie)[i] = v
	}
	return nil
}
