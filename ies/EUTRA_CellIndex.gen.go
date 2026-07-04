// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: EUTRA-CellIndex (line 9408).
// EUTRA-CellIndex ::=                         INTEGER (1..maxCellMeasEUTRA)

var eUTRACellIndexConstraints = per.Constrained(1, common.MaxCellMeasEUTRA)

type EUTRA_CellIndex struct {
	Value int64
}

func (v *EUTRA_CellIndex) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, eUTRACellIndexConstraints)
}

func (v *EUTRA_CellIndex) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(eUTRACellIndexConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
