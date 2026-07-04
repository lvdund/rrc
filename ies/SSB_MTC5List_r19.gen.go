// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SSB-MTC5List-r19 (line 9488).
// SSB-MTC5List-r19::=                 SEQUENCE (SIZE(1..3)) OF SSB-MTC5-r19

var sSBMTC5ListR19SizeConstraints = per.SizeRange(1, 3)

type SSB_MTC5List_r19 []SSB_MTC5_r19

func (ie *SSB_MTC5List_r19) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(sSBMTC5ListR19SizeConstraints)
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

func (ie *SSB_MTC5List_r19) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(sSBMTC5ListR19SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(SSB_MTC5List_r19, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
