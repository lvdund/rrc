// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SSB-MTC4List-r17 (line 9486).
// SSB-MTC4List-r17::=                 SEQUENCE (SIZE(1..3)) OF SSB-MTC4-r17

var sSBMTC4ListR17SizeConstraints = per.SizeRange(1, 3)

type SSB_MTC4List_r17 []SSB_MTC4_r17

func (ie *SSB_MTC4List_r17) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(sSBMTC4ListR17SizeConstraints)
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

func (ie *SSB_MTC4List_r17) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(sSBMTC4ListR17SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(SSB_MTC4List_r17, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
