// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: TimeBetweenEvent-r17 (line 9888).
// TimeBetweenEvent-r17 ::= INTEGER (0..1023)

var timeBetweenEventR17Constraints = per.Constrained(0, 1023)

type TimeBetweenEvent_r17 struct {
	Value int64
}

func (v *TimeBetweenEvent_r17) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, timeBetweenEventR17Constraints)
}

func (v *TimeBetweenEvent_r17) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(timeBetweenEventR17Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
