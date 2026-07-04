// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: DormancyGroupID-r16 (line 14782).
// DormancyGroupID-r16 ::=         INTEGER (0..4)

var dormancyGroupIDR16Constraints = per.Constrained(0, 4)

type DormancyGroupID_r16 struct {
	Value int64
}

func (v *DormancyGroupID_r16) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, dormancyGroupIDR16Constraints)
}

func (v *DormancyGroupID_r16) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(dormancyGroupIDR16Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
