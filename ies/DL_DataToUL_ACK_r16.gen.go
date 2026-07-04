// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: DL-DataToUL-ACK-r16 (line 12163).
// DL-DataToUL-ACK-r16 ::=                    SEQUENCE (SIZE (1..8)) OF INTEGER (-1..15)

var dLDataToULACKR16SizeConstraints = per.SizeRange(1, 8)

type DL_DataToUL_ACK_r16 []int64

func (ie *DL_DataToUL_ACK_r16) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(dLDataToULACKR16SizeConstraints)
	if err := seqOf.EncodeLength(int64(len(*ie))); err != nil {
		return err
	}
	for i := range *ie {
		if err := e.EncodeInteger((*ie)[i], per.Constrained(-1, 15)); err != nil {
			return err
		}
	}
	return nil
}

func (ie *DL_DataToUL_ACK_r16) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(dLDataToULACKR16SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(DL_DataToUL_ACK_r16, n)
	for i := int64(0); i < n; i++ {
		v, err := d.DecodeInteger(per.Constrained(-1, 15))
		if err != nil {
			return err
		}
		(*ie)[i] = v
	}
	return nil
}
