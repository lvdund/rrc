// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CellReselectionPriority (line 5887).
// CellReselectionPriority ::=             INTEGER (0..7)

var cellReselectionPriorityConstraints = per.Constrained(0, 7)

type CellReselectionPriority struct {
	Value int64
}

func (v *CellReselectionPriority) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, cellReselectionPriorityConstraints)
}

func (v *CellReselectionPriority) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(cellReselectionPriorityConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
