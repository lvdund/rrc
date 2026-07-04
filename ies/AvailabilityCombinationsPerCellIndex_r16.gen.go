// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: AvailabilityCombinationsPerCellIndex-r16 (line 5064).
// AvailabilityCombinationsPerCellIndex-r16 ::= INTEGER(0..maxNrofDUCells-r16)

var availabilityCombinationsPerCellIndexR16Constraints = per.Constrained(0, common.MaxNrofDUCells_r16)

type AvailabilityCombinationsPerCellIndex_r16 struct {
	Value int64
}

func (v *AvailabilityCombinationsPerCellIndex_r16) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, availabilityCombinationsPerCellIndexR16Constraints)
}

func (v *AvailabilityCombinationsPerCellIndex_r16) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(availabilityCombinationsPerCellIndexR16Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
