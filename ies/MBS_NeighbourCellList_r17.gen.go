// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MBS-NeighbourCellList-r17 (line 28454).
// MBS-NeighbourCellList-r17 ::=     SEQUENCE (SIZE (0..maxNeighCellMBS-r17)) OF MBS-NeighbourCell-r17

var mBSNeighbourCellListR17SizeConstraints = per.SizeRange(0, common.MaxNeighCellMBS_r17)

type MBS_NeighbourCellList_r17 []MBS_NeighbourCell_r17

func (ie *MBS_NeighbourCellList_r17) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(mBSNeighbourCellListR17SizeConstraints)
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

func (ie *MBS_NeighbourCellList_r17) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(mBSNeighbourCellListR17SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(MBS_NeighbourCellList_r17, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
