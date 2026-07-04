// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: FeaturePriority-r17 (line 2107).
// FeaturePriority-r17 ::= INTEGER (0..7)

var featurePriorityR17Constraints = per.Constrained(0, 7)

type FeaturePriority_r17 struct {
	Value int64
}

func (v *FeaturePriority_r17) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, featurePriorityR17Constraints)
}

func (v *FeaturePriority_r17) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(featurePriorityR17Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
