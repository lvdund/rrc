// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UL-AccessConfigListDCI-0-2-r17 (line 12484).
// UL-AccessConfigListDCI-0-2-r17 ::= SEQUENCE (SIZE (1..64)) OF INTEGER (0..63)

var uLAccessConfigListDCI02R17SizeConstraints = per.SizeRange(1, 64)

type UL_AccessConfigListDCI_0_2_r17 []int64

func (ie *UL_AccessConfigListDCI_0_2_r17) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(uLAccessConfigListDCI02R17SizeConstraints)
	if err := seqOf.EncodeLength(int64(len(*ie))); err != nil {
		return err
	}
	for i := range *ie {
		if err := e.EncodeInteger((*ie)[i], per.Constrained(0, 63)); err != nil {
			return err
		}
	}
	return nil
}

func (ie *UL_AccessConfigListDCI_0_2_r17) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(uLAccessConfigListDCI02R17SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(UL_AccessConfigListDCI_0_2_r17, n)
	for i := int64(0); i < n; i++ {
		v, err := d.DecodeInteger(per.Constrained(0, 63))
		if err != nil {
			return err
		}
		(*ie)[i] = v
	}
	return nil
}
