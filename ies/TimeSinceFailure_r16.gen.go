// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: TimeSinceFailure-r16 (line 3548).
// TimeSinceFailure-r16 ::= INTEGER (0..172800)

var timeSinceFailureR16Constraints = per.Constrained(0, 172800)

type TimeSinceFailure_r16 struct {
	Value int64
}

func (v *TimeSinceFailure_r16) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, timeSinceFailureR16Constraints)
}

func (v *TimeSinceFailure_r16) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(timeSinceFailureR16Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
