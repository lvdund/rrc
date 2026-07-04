// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UL-AccessConfigListDCI-0-1-r16 (line 12480).
// UL-AccessConfigListDCI-0-1-r16 ::= SEQUENCE (SIZE (1..64)) OF INTEGER (0..63)

var uLAccessConfigListDCI01R16SizeConstraints = per.SizeRange(1, 64)

type UL_AccessConfigListDCI_0_1_r16 []int64

func (ie *UL_AccessConfigListDCI_0_1_r16) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(uLAccessConfigListDCI01R16SizeConstraints)
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

func (ie *UL_AccessConfigListDCI_0_1_r16) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(uLAccessConfigListDCI01R16SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(UL_AccessConfigListDCI_0_1_r16, n)
	for i := int64(0); i < n; i++ {
		v, err := d.DecodeInteger(per.Constrained(0, 63))
		if err != nil {
			return err
		}
		(*ie)[i] = v
	}
	return nil
}
