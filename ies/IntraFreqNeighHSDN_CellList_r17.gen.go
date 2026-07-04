// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: IntraFreqNeighHSDN-CellList-r17 (line 3926).
// IntraFreqNeighHSDN-CellList-r17 ::= SEQUENCE (SIZE (1..maxCellIntra)) OF PCI-Range

var intraFreqNeighHSDNCellListR17SizeConstraints = per.SizeRange(1, common.MaxCellIntra)

type IntraFreqNeighHSDN_CellList_r17 []PCI_Range

func (ie *IntraFreqNeighHSDN_CellList_r17) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(intraFreqNeighHSDNCellListR17SizeConstraints)
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

func (ie *IntraFreqNeighHSDN_CellList_r17) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(intraFreqNeighHSDNCellListR17SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(IntraFreqNeighHSDN_CellList_r17, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
