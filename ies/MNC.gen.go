// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MNC (line 11892).
// MNC ::=                             SEQUENCE (SIZE (2..3)) OF MCC-MNC-Digit

var mNCSizeConstraints = per.SizeRange(2, 3)

type MNC []MCC_MNC_Digit

func (ie *MNC) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(mNCSizeConstraints)
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

func (ie *MNC) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(mNCSizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(MNC, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
