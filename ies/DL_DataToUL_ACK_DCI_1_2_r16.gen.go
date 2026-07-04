// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: DL-DataToUL-ACK-DCI-1-2-r16 (line 12171).
// DL-DataToUL-ACK-DCI-1-2-r16 ::=            SEQUENCE (SIZE (1..8)) OF INTEGER (0..15)

var dLDataToULACKDCI12R16SizeConstraints = per.SizeRange(1, 8)

type DL_DataToUL_ACK_DCI_1_2_r16 []int64

func (ie *DL_DataToUL_ACK_DCI_1_2_r16) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(dLDataToULACKDCI12R16SizeConstraints)
	if err := seqOf.EncodeLength(int64(len(*ie))); err != nil {
		return err
	}
	for i := range *ie {
		if err := e.EncodeInteger((*ie)[i], per.Constrained(0, 15)); err != nil {
			return err
		}
	}
	return nil
}

func (ie *DL_DataToUL_ACK_DCI_1_2_r16) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(dLDataToULACKDCI12R16SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(DL_DataToUL_ACK_DCI_1_2_r16, n)
	for i := int64(0); i < n; i++ {
		v, err := d.DecodeInteger(per.Constrained(0, 15))
		if err != nil {
			return err
		}
		(*ie)[i] = v
	}
	return nil
}
