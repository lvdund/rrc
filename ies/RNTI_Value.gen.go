// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RNTI-Value (line 14201).
// RNTI-Value ::=                      INTEGER (0..65535)

var rNTIValueConstraints = per.Constrained(0, 65535)

type RNTI_Value struct {
	Value int64
}

func (v *RNTI_Value) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, rNTIValueConstraints)
}

func (v *RNTI_Value) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(rNTIValueConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
