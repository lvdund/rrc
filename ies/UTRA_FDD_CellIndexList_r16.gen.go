// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: UTRA-FDD-CellIndexList-r16 (line 9680).
// UTRA-FDD-CellIndexList-r16 ::=              SEQUENCE (SIZE (1..maxCellMeasUTRA-FDD-r16)) OF UTRA-FDD-CellIndex-r16

var uTRAFDDCellIndexListR16SizeConstraints = per.SizeRange(1, common.MaxCellMeasUTRA_FDD_r16)

type UTRA_FDD_CellIndexList_r16 []UTRA_FDD_CellIndex_r16

func (ie *UTRA_FDD_CellIndexList_r16) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(uTRAFDDCellIndexListR16SizeConstraints)
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

func (ie *UTRA_FDD_CellIndexList_r16) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(uTRAFDDCellIndexListR16SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(UTRA_FDD_CellIndexList_r16, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
