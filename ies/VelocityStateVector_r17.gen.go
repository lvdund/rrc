// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: VelocityStateVector-r17 (line 8287).
// VelocityStateVector-r17 ::= INTEGER (-131072..131071)

var velocityStateVectorR17Constraints = per.Constrained(-131072, 131071)

type VelocityStateVector_r17 struct {
	Value int64
}

func (v *VelocityStateVector_r17) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, velocityStateVectorR17Constraints)
}

func (v *VelocityStateVector_r17) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(velocityStateVectorR17Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
