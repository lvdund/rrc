// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CellsToAddModListUTRA-FDD-r16 (line 9673).
// CellsToAddModListUTRA-FDD-r16 ::=           SEQUENCE (SIZE (1..maxCellMeasUTRA-FDD-r16)) OF CellsToAddModUTRA-FDD-r16

var cellsToAddModListUTRAFDDR16SizeConstraints = per.SizeRange(1, common.MaxCellMeasUTRA_FDD_r16)

type CellsToAddModListUTRA_FDD_r16 []CellsToAddModUTRA_FDD_r16

func (ie *CellsToAddModListUTRA_FDD_r16) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(cellsToAddModListUTRAFDDR16SizeConstraints)
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

func (ie *CellsToAddModListUTRA_FDD_r16) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(cellsToAddModListUTRAFDDR16SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(CellsToAddModListUTRA_FDD_r16, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
