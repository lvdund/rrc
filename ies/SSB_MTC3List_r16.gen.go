// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SSB-MTC3List-r16 (line 9484).
// SSB-MTC3List-r16::=                 SEQUENCE (SIZE(1..4)) OF SSB-MTC3-r16

var sSBMTC3ListR16SizeConstraints = per.SizeRange(1, 4)

type SSB_MTC3List_r16 []SSB_MTC3_r16

func (ie *SSB_MTC3List_r16) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(sSBMTC3ListR16SizeConstraints)
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

func (ie *SSB_MTC3List_r16) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(sSBMTC3ListR16SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(SSB_MTC3List_r16, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
