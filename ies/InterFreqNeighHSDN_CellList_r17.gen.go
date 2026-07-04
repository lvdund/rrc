// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: InterFreqNeighHSDN-CellList-r17 (line 4065).
// InterFreqNeighHSDN-CellList-r17 ::= SEQUENCE (SIZE (1..maxCellInter)) OF PCI-Range

var interFreqNeighHSDNCellListR17SizeConstraints = per.SizeRange(1, common.MaxCellInter)

type InterFreqNeighHSDN_CellList_r17 []PCI_Range

func (ie *InterFreqNeighHSDN_CellList_r17) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(interFreqNeighHSDNCellListR17SizeConstraints)
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

func (ie *InterFreqNeighHSDN_CellList_r17) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(interFreqNeighHSDNCellListR17SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(InterFreqNeighHSDN_CellList_r17, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
