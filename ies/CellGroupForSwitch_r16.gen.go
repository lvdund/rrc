// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CellGroupForSwitch-r16 (line 11037).
// CellGroupForSwitch-r16 ::=          SEQUENCE(SIZE (1..16)) OF ServCellIndex

var cellGroupForSwitchR16SizeConstraints = per.SizeRange(1, 16)

type CellGroupForSwitch_r16 []ServCellIndex

func (ie *CellGroupForSwitch_r16) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(cellGroupForSwitchR16SizeConstraints)
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

func (ie *CellGroupForSwitch_r16) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(cellGroupForSwitchR16SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(CellGroupForSwitch_r16, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
