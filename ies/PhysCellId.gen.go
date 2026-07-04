// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PhysCellId (line 11591).
// PhysCellId ::=                      INTEGER (0..1007)

var physCellIdConstraints = per.Constrained(0, 1007)

type PhysCellId struct {
	Value int64
}

func (v *PhysCellId) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, physCellIdConstraints)
}

func (v *PhysCellId) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(physCellIdConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
