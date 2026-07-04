// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PhysCellIdUTRA-FDD-r16 (line 26640).
// PhysCellIdUTRA-FDD-r16 ::=        INTEGER (0..511)

var physCellIdUTRAFDDR16Constraints = per.Constrained(0, 511)

type PhysCellIdUTRA_FDD_r16 struct {
	Value int64
}

func (v *PhysCellIdUTRA_FDD_r16) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, physCellIdUTRAFDDR16Constraints)
}

func (v *PhysCellIdUTRA_FDD_r16) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(physCellIdUTRAFDDR16Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
