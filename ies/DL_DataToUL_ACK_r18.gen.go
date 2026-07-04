// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: DL-DataToUL-ACK-r18 (line 12169).
// DL-DataToUL-ACK-r18 ::=                    SEQUENCE (SIZE (1..8)) OF INTEGER (0..31)

var dLDataToULACKR18SizeConstraints = per.SizeRange(1, 8)

type DL_DataToUL_ACK_r18 []int64

func (ie *DL_DataToUL_ACK_r18) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(dLDataToULACKR18SizeConstraints)
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

func (ie *DL_DataToUL_ACK_r18) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(dLDataToULACKR18SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(DL_DataToUL_ACK_r18, n)
	for i := int64(0); i < n; i++ {
		v, err := d.DecodeInteger(per.Constrained(0, 31))
		if err != nil {
			return err
		}
		(*ie)[i] = v
	}
	return nil
}
