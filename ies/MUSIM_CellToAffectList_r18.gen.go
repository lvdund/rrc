// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MUSIM-CellToAffectList-r18 (line 2641).
// MUSIM-CellToAffectList-r18::=           SEQUENCE (SIZE (1..maxNrofServingCells)) OF MUSIM-CellToAffect-r18

var mUSIMCellToAffectListR18SizeConstraints = per.SizeRange(1, common.MaxNrofServingCells)

type MUSIM_CellToAffectList_r18 []MUSIM_CellToAffect_r18

func (ie *MUSIM_CellToAffectList_r18) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(mUSIMCellToAffectListR18SizeConstraints)
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

func (ie *MUSIM_CellToAffectList_r18) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(mUSIMCellToAffectListR18SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(MUSIM_CellToAffectList_r18, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
