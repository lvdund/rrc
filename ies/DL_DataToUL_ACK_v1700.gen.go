// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: DL-DataToUL-ACK-v1700 (line 12167).
// DL-DataToUL-ACK-v1700 ::=                  SEQUENCE (SIZE (1..8)) OF INTEGER (16..31)

var dLDataToULACKV1700SizeConstraints = per.SizeRange(1, 8)

type DL_DataToUL_ACK_v1700 []int64

func (ie *DL_DataToUL_ACK_v1700) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(dLDataToULACKV1700SizeConstraints)
	if err := seqOf.EncodeLength(int64(len(*ie))); err != nil {
		return err
	}
	for i := range *ie {
		if err := e.EncodeInteger((*ie)[i], per.Constrained(16, 31)); err != nil {
			return err
		}
	}
	return nil
}

func (ie *DL_DataToUL_ACK_v1700) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(dLDataToULACKV1700SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(DL_DataToUL_ACK_v1700, n)
	for i := int64(0); i < n; i++ {
		v, err := d.DecodeInteger(per.Constrained(16, 31))
		if err != nil {
			return err
		}
		(*ie)[i] = v
	}
	return nil
}
