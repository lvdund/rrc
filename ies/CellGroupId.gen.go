// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CellGroupId (line 5877).
// CellGroupId ::=                             INTEGER (0.. maxSecondaryCellGroups)

var cellGroupIdConstraints = per.Constrained(0, common.MaxSecondaryCellGroups)

type CellGroupId struct {
	Value int64
}

func (v *CellGroupId) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, cellGroupIdConstraints)
}

func (v *CellGroupId) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(cellGroupIdConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
