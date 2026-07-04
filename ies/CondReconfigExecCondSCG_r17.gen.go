// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CondReconfigExecCondSCG-r17 (line 6481).
// CondReconfigExecCondSCG-r17 ::=  SEQUENCE (SIZE (1..2)) OF MeasId

var condReconfigExecCondSCGR17SizeConstraints = per.SizeRange(1, 2)

type CondReconfigExecCondSCG_r17 []MeasId

func (ie *CondReconfigExecCondSCG_r17) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(condReconfigExecCondSCGR17SizeConstraints)
	if err := seqOf.EncodeLength(int64(len(*ie))); err != nil {
		return err
	}
	for i := range *ie {
		if err := (*ie)[i].Encode(e); err != nil {
			return err
		}
	}
	return nil
}

func (ie *CondReconfigExecCondSCG_r17) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(condReconfigExecCondSCGR17SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(CondReconfigExecCondSCG_r17, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
