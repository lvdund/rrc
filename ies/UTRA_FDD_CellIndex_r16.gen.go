// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: UTRA-FDD-CellIndex-r16 (line 9682).
// UTRA-FDD-CellIndex-r16 ::=                  INTEGER (1..maxCellMeasUTRA-FDD-r16)

var uTRAFDDCellIndexR16Constraints = per.Constrained(1, common.MaxCellMeasUTRA_FDD_r16)

type UTRA_FDD_CellIndex_r16 struct {
	Value int64
}

func (v *UTRA_FDD_CellIndex_r16) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, uTRAFDDCellIndexR16Constraints)
}

func (v *UTRA_FDD_CellIndex_r16) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(uTRAFDDCellIndexR16Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
