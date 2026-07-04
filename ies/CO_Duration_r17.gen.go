// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CO-Duration-r17 (line 15202).
// CO-Duration-r17 ::=    INTEGER (0..4480)

var cODurationR17Constraints = per.Constrained(0, 4480)

type CO_Duration_r17 struct {
	Value int64
}

func (v *CO_Duration_r17) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, cODurationR17Constraints)
}

func (v *CO_Duration_r17) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(cODurationR17Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
