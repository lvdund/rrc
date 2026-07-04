// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PositionStateVector-r17 (line 8285).
// PositionStateVector-r17 ::= INTEGER (-33554432..33554431)

var positionStateVectorR17Constraints = per.Constrained(-33554432, 33554431)

type PositionStateVector_r17 struct {
	Value int64
}

func (v *PositionStateVector_r17) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, positionStateVectorR17Constraints)
}

func (v *PositionStateVector_r17) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(positionStateVectorR17Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
