// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UCI-OnPUSCH-ListDCI-0-1-r16 (line 12478).
// UCI-OnPUSCH-ListDCI-0-1-r16 ::=  SEQUENCE (SIZE (1..2)) OF UCI-OnPUSCH

var uCIOnPUSCHListDCI01R16SizeConstraints = per.SizeRange(1, 2)

type UCI_OnPUSCH_ListDCI_0_1_r16 []UCI_OnPUSCH

func (ie *UCI_OnPUSCH_ListDCI_0_1_r16) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(uCIOnPUSCHListDCI01R16SizeConstraints)
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

func (ie *UCI_OnPUSCH_ListDCI_0_1_r16) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(uCIOnPUSCHListDCI01R16SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(UCI_OnPUSCH_ListDCI_0_1_r16, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
