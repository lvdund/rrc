// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-SelectionWindowList-r16 (line 28082).
// SL-SelectionWindowList-r16 ::=         SEQUENCE (SIZE (8)) OF SL-SelectionWindowConfig-r16

var sLSelectionWindowListR16SizeConstraints = per.FixedSize(8)

type SL_SelectionWindowList_r16 []SL_SelectionWindowConfig_r16

func (ie *SL_SelectionWindowList_r16) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(sLSelectionWindowListR16SizeConstraints)
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

func (ie *SL_SelectionWindowList_r16) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(sLSelectionWindowListR16SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(SL_SelectionWindowList_r16, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
