// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CO-Duration-r16 (line 15201).
// CO-Duration-r16 ::=    INTEGER (0..1120)

var cODurationR16Constraints = per.Constrained(0, 1120)

type CO_Duration_r16 struct {
	Value int64
}

func (v *CO_Duration_r16) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, cODurationR16Constraints)
}

func (v *CO_Duration_r16) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(cODurationR16Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
