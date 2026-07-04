// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: ServCellIndex (line 14611).
// ServCellIndex ::=                   INTEGER (0..maxNrofServingCells-1)

var servCellIndexConstraints = per.Constrained(0, common.MaxNrofServingCells_1)

type ServCellIndex struct {
	Value int64
}

func (v *ServCellIndex) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, servCellIndexConstraints)
}

func (v *ServCellIndex) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(servCellIndexConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
