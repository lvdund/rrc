// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-CG-MaxTransNumList-r16 (line 27079).
// SL-CG-MaxTransNumList-r16 ::=     SEQUENCE (SIZE (1..8)) OF SL-CG-MaxTransNum-r16

var sLCGMaxTransNumListR16SizeConstraints = per.SizeRange(1, 8)

type SL_CG_MaxTransNumList_r16 []SL_CG_MaxTransNum_r16

func (ie *SL_CG_MaxTransNumList_r16) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(sLCGMaxTransNumListR16SizeConstraints)
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

func (ie *SL_CG_MaxTransNumList_r16) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(sLCGMaxTransNumListR16SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(SL_CG_MaxTransNumList_r16, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
