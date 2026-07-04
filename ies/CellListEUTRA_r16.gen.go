// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CellListEUTRA-r16 (line 9313).
// CellListEUTRA-r16  ::=    SEQUENCE (SIZE (1..maxCellMeasIdle-r16)) OF EUTRA-PhysCellIdRange

var cellListEUTRAR16SizeConstraints = per.SizeRange(1, common.MaxCellMeasIdle_r16)

type CellListEUTRA_r16 []EUTRA_PhysCellIdRange

func (ie *CellListEUTRA_r16) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(cellListEUTRAR16SizeConstraints)
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

func (ie *CellListEUTRA_r16) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(cellListEUTRAR16SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(CellListEUTRA_r16, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
