// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MCC (line 11890).
// MCC ::=                             SEQUENCE (SIZE (3)) OF MCC-MNC-Digit

var mCCSizeConstraints = per.FixedSize(3)

type MCC []MCC_MNC_Digit

func (ie *MCC) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(mCCSizeConstraints)
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

func (ie *MCC) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(mCCSizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(MCC, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
