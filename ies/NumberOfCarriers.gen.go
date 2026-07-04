// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: NumberOfCarriers (line 23486).
// NumberOfCarriers ::=    INTEGER (1..16)

var numberOfCarriersConstraints = per.Constrained(1, 16)

type NumberOfCarriers struct {
	Value int64
}

func (v *NumberOfCarriers) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, numberOfCarriersConstraints)
}

func (v *NumberOfCarriers) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(numberOfCarriersConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
