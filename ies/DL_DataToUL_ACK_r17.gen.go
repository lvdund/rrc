// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: DL-DataToUL-ACK-r17 (line 12165).
// DL-DataToUL-ACK-r17 ::=                    SEQUENCE (SIZE (1..8)) OF INTEGER (-1..127)

var dLDataToULACKR17SizeConstraints = per.SizeRange(1, 8)

type DL_DataToUL_ACK_r17 []int64

func (ie *DL_DataToUL_ACK_r17) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(dLDataToULACKR17SizeConstraints)
	if err := seqOf.EncodeLength(int64(len(*ie))); err != nil {
		return err
	}
	for i := range *ie {
		if err := e.EncodeInteger((*ie)[i], per.Constrained(-1, 127)); err != nil {
			return err
		}
	}
	return nil
}

func (ie *DL_DataToUL_ACK_r17) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(dLDataToULACKR17SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(DL_DataToUL_ACK_r17, n)
	for i := int64(0); i < n; i++ {
		v, err := d.DecodeInteger(per.Constrained(-1, 127))
		if err != nil {
			return err
		}
		(*ie)[i] = v
	}
	return nil
}
