// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: ChoCandidateCellList-r17 (line 3463).
// ChoCandidateCellList-r17 ::=         SEQUENCE(SIZE (1..maxNrofCondCells-r16)) OF ChoCandidateCell-r17

var choCandidateCellListR17SizeConstraints = per.SizeRange(1, common.MaxNrofCondCells_r16)

type ChoCandidateCellList_r17 []ChoCandidateCell_r17

func (ie *ChoCandidateCellList_r17) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(choCandidateCellListR17SizeConstraints)
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

func (ie *ChoCandidateCellList_r17) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(choCandidateCellListR17SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(ChoCandidateCellList_r17, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
